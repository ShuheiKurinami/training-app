// src/services/UserService.ts
import apiClient from './ApiService';
import { UserModel } from '../models/UserModel';

export const fetchUsers = async (): Promise<UserModel[]> => {
    const response = await apiClient.get<UserModel[]>('/users');
    return response.data;
};

export const createUser = async (user: Omit<UserModel, 'id'>): Promise<UserModel> => {
    const response = await apiClient.post<UserModel>('/users', user);
    return response.data;
};
