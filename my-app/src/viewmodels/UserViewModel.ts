import { useState, useEffect } from 'react';
import { UserModel } from '../models/UserModel';
import { fetchUsers, createUser, updateUser, deleteUser } from '../services/UserService';

export const useUserViewModel = () => {
    const [users, setUsers] = useState<UserModel[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    // ユーザー一覧を取得
    useEffect(() => {
        const getUsers = async () => {
            setLoading(true);
            try {
                const data = await fetchUsers();
                setUsers(data);
            } catch (err) {
                setError('Failed to fetch users');
            } finally {
                setLoading(false);
            }
        };

        getUsers();
    }, []);

    // ユーザーを追加
    const addUser = async (user: Omit<UserModel, 'id'>) => {
        try {
            const newUser = await createUser(user);
            setUsers([...users, newUser]);
        } catch (err) {
            setError('Failed to add user');
        }
    };

    // ユーザーを更新
    const updateUserInState = async (id: number, updatedUser: Omit<UserModel, 'id'>) => {
        try {
            const newUserData = await updateUser(id, updatedUser);
            setUsers(users.map((user) => (user.id === id ? newUserData : user)));
        } catch (err) {
            setError('Failed to update user');
        }
    };

    // ユーザーを削除
    const removeUser = async (id: number) => {
        try {
            await deleteUser(id);
            setUsers(users.filter((user) => user.id !== id));
        } catch (err) {
            setError('Failed to delete user');
        }
    };

    return { users, addUser, updateUserInState, removeUser, loading, error };
};
