// src/viewmodels/UserViewModel.ts
import { useState, useEffect } from 'react';
import { UserModel } from '../models/UserModel';
import { fetchUsers } from '../services/UserService';

export const useUserViewModel = () => {
    const [users, setUsers] = useState<UserModel[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

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

    const addUser = (user: UserModel) => {
        setUsers([...users, user]);
    };

    return { users, addUser, loading, error };
};
