import React, { useState } from 'react';
import { UserModel } from '../../models/UserModel';
import { useUserViewModel } from '../../viewmodels/UserViewModel';

interface PasswordChangeModalProps {
    isOpen: boolean;
    onClose: () => void;
    user: UserModel | any;
}

const PasswordChangeModal: React.FC<PasswordChangeModalProps> = ({ isOpen, onClose, user }) => {
    const [newPassword, setNewPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const { changePassword, error, successMessage } = useUserViewModel();

    if (!isOpen) return null;

    const handleChangePassword = async () => {
        await changePassword(user.id, newPassword, confirmPassword);
        if (!error) {
            setTimeout(onClose, 2000);
        }
    };

    return (
        <div className="modal">
            <div className="modal-content">
                <h2>Change Password for {user.username}</h2>
                {error && <p style={{ color: 'red' }}>{error}</p>}
                {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
                <input
                    type="password"
                    placeholder="New Password"
                    value={newPassword}
                    onChange={(e) => setNewPassword(e.target.value)}
                />
                <input
                    type="password"
                    placeholder="Confirm Password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                />
                <button onClick={handleChangePassword}>Change Password</button>
                <button onClick={onClose}>Cancel</button>
            </div>
        </div>
    );
};

export default PasswordChangeModal;
