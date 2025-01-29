// src/services/GenreService.ts
import apiClient from './ApiService';
import { GenreModel } from '../models/GenreModel';

export const fetchGenres = async (): Promise<GenreModel[]> => {
    const response = await apiClient.get<GenreModel[]>('/genres');
    return response.data;
};

export const createGenre = async (genre: Omit<GenreModel, 'id'>): Promise<GenreModel> => {
    const response = await apiClient.post<GenreModel>('/genres', genre);
    return response.data;
};
