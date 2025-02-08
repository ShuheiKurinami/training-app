import React, { useState } from 'react';
import { useUserViewModel } from '../../viewmodels/UserViewModel';
import UserFormModal from './UserFormModal';
import PasswordChangeModal from './PasswordChangeModal';

const UserList: React.FC = () => {
    const { users, loading, error } = useUserViewModel();
    const [modalOpen, setModalOpen] = useState(false);
    const [modalMode, setModalMode] = useState<'create' | 'edit' | 'delete'>('create');
    const [passwordModalOpen, setPasswordModalOpen] = useState(false); // 🔹 パスワード変更モーダルの状態
    const [selectedUser, setSelectedUser] = useState<{ id?: number; username: string; email: string } | null>(null);

    if (loading) return <p>Loading users...</p>;
    if (error) return <p>{error}</p>;

    const handleOpenModal = (mode: 'create' | 'edit' | 'delete', user?: { id: number; username: string; email: string }) => {
        setSelectedUser(user || null);
        setModalMode(mode);
        setModalOpen(true);
    };

    const handleOpenPasswordModal = (user: { id: number; username: string; email: string }) => {
        setSelectedUser(user);
        setPasswordModalOpen(true);
    };

    return (
        <div>
            <button onClick={() => handleOpenModal('create')}>Add User</button>

            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        {user.username} ({user.email})
                        <button onClick={() => handleOpenModal('edit', user)}>Edit</button>
                        <button onClick={() => handleOpenModal('delete', user)}>Delete</button>
                        <button onClick={() => handleOpenPasswordModal(user)}>Change Password</button>
                    </li>
                ))}
            </ul>

            {modalOpen && (
                <UserFormModal
                    isOpen={modalOpen}
                    onClose={() => setModalOpen(false)}
                    modalMode={modalMode}
                    user={selectedUser}
                />
            )}

            {passwordModalOpen && selectedUser && (
                <PasswordChangeModal
                    isOpen={passwordModalOpen}
                    onClose={() => setPasswordModalOpen(false)}
                    user={selectedUser}
                />
            )}
        </div>
    );
};

export default UserList;
