// src/views/genre/GenreList.tsx
import React from 'react';
import { useGenreViewModel } from '../../viewmodels/GenreViewModel';

const GenreList: React.FC = () => {
    const { genres, loading, error } = useGenreViewModel();

    if (loading) return <p>Loading genres...</p>;
    if (error) return <p>{error}</p>;

    return (
        <ul>
            {genres.map((genre) => (
                <li key={genre.id}>{genre.name}</li>
            ))}
        </ul>
    );
};

export default GenreList;
