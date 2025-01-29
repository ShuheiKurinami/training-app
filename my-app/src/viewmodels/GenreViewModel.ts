// src/viewmodels/GenreViewModel.ts
import { useState, useEffect } from 'react';
import { GenreModel } from '../models/GenreModel';
import { fetchGenres } from '../services/GenreService';

export const useGenreViewModel = () => {
    const [genres, setGenres] = useState<GenreModel[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const getGenres = async () => {
            setLoading(true);
            try {
                const data = await fetchGenres();
                setGenres(data);
            } catch (err) {
                setError('Failed to fetch genres');
            } finally {
                setLoading(false);
            }
        };

        getGenres();
    }, []);

    const addGenre = (genre: GenreModel) => {
        setGenres([...genres, genre]);
    };

    return { genres, addGenre, loading, error };
};
