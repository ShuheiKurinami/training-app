// src/views/genre/GenreForm.tsx
import React, { useState } from 'react';
import { GenreModel } from '../../models/GenreModel';
import { useGenreViewModel } from '../../viewmodels/GenreViewModel';
import { createGenre } from '../../services/GenreService';

const GenreForm: React.FC = () => {
    const { addGenre } = useGenreViewModel();
    const [name, setName] = useState<string>('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const newGenre = await createGenre({ name });
            addGenre(newGenre);
            setName('');
        } catch (error) {
            console.error('Error adding genre:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Genre Name:</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Add Genre</button>
        </form>
    );
};

export default GenreForm;
