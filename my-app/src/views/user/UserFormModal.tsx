import React, { useState } from 'react';
import { createUser, updateUser, deleteUser } from '../../services/UserService';
import { useUserViewModel } from '../../viewmodels/UserViewModel';

interface UserFormModalProps {
    isOpen: boolean;
    onClose: () => void;
    modalMode: 'create' | 'edit' | 'delete';
    user?: { id?: number; username: string; email: string } | null;
}

const UserFormModal: React.FC<UserFormModalProps> = ({ isOpen, onClose, modalMode, user }) => {
    const { addUser, updateUserInState, removeUser } = useUserViewModel();
    const [name, setName] = useState(user?.username || '');
    const [email, setEmail] = useState(user?.email || '');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            if (modalMode === 'create') {
                const newUser = await createUser({ username: name, email });
                addUser(newUser);
            } else if (modalMode === 'edit' && user?.id) {
                await updateUser(user.id, { username: name, email });
                updateUserInState(user.id, { username: name, email });
            }
            onClose();
        } catch (error) {
            console.error(`Error ${modalMode} user:`, error);
        }
    };

    const handleDelete = async () => {
        if (!user?.id) return;
        try {
            await deleteUser(user.id);
            removeUser(user.id);
            onClose();
        } catch (error) {
            console.error('Error deleting user:', error);
        }
    };

    if (!isOpen) return null;

    return (
        <div className="modal">
            <div className="modal-content">
                <h2>{modalMode === 'create' ? 'Add User' : modalMode === 'edit' ? 'Edit User' : 'Delete User'}</h2>

                {modalMode !== 'delete' ? (
                    <form onSubmit={handleSubmit}>
                        <div>
                            <label>Name:</label>
                            <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
                        </div>
                        <div>
                            <label>Email:</label>
                            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
                        </div>
                        <button type="submit">{modalMode === 'create' ? 'Create' : 'Update'}</button>
                    </form>
                ) : (
                    <div>
                        <p>Are you sure you want to delete this user?</p>
                        <button onClick={handleDelete}>Delete</button>
                    </div>
                )}

                <button onClick={onClose}>Close</button>
            </div>
        </div>
    );
};

export default UserFormModal;
