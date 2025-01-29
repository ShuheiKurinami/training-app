// src/views/user/UserList.tsx
import React from 'react';
import { useUserViewModel } from '../../viewmodels/UserViewModel';

const UserList: React.FC = () => {
    const { users, loading, error } = useUserViewModel();

    if (loading) return <p>Loading users...</p>;
    if (error) return <p>{error}</p>;

    return (
        <ul>
            {users.map((user) => (
                <li key={user.id}>
                    {user.name} ({user.email})
                </li>
            ))}
        </ul>
    );
};

export default UserList;
