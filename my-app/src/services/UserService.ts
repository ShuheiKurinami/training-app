// src/services/UserService.ts
import apiClient from './ApiService';
import { UserModel } from '../models/UserModel';

export const fetchUsers = async (): Promise<UserModel[]> => {
    const response = await apiClient.get<UserModel[]>('/users');
    return response.data;
};

export const createUser = async (user: Omit<UserModel, 'id'>): Promise<UserModel> => {
    console.log("Sending user data:", user);
    const response = await apiClient.post<UserModel>('/users', user);
    console.log("API response:", response.data); // ここでレスポンス確認
    return response.data;
};

export const updateUser = async (id: number, user: Omit<UserModel, 'id'>): Promise<UserModel> => {
    const response = await apiClient.put<UserModel>(`/users/${id}`, user);
    return response.data;
};

export const deleteUser = async (id: number): Promise<void> => {
    await apiClient.delete(`/users/${id}`);
};
