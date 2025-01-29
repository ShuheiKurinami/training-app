// src/views/user/UserForm.tsx
import React, { useState } from 'react';
import { UserModel } from '../../models/UserModel';
import { useUserViewModel } from '../../viewmodels/UserViewModel';
import { createUser } from '../../services/UserService';

const UserForm: React.FC = () => {
    const { addUser } = useUserViewModel();
    const [name, setName] = useState<string>('');
    const [email, setEmail] = useState<string>('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const newUser = await createUser({ name, email });
            addUser(newUser);
            setName('');
            setEmail('');
        } catch (error) {
            console.error('Error adding user:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Name:</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Email:</label>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Add User</button>
        </form>
    );
};

export default UserForm;
