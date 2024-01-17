import axios from 'axios';

const createUser = async (user) => {
    return await axios.post(`${import.meta.env.VITE_BACKEND_URL}/user`, user);
};


const getUserByID = async (id) => {
    return await axios.get(`${import.meta.env.VITE_BACKEND_URL}/user/${id}`);

};

const getUsers = async (page, limit) => {
    return await axios.get(`${import.meta.env.VITE_BACKEND_URL}/user?page=${page}&limit=${limit}`);

};

const updateUser = async (user, id) => {
    return await axios.put(`${import.meta.env.VITE_BACKEND_URL}/user/${id}`, user);

};

const deleteUser = async (id) => {
    return await axios.delete(`${import.meta.env.VITE_BACKEND_URL}/user/${id}`);

};


export { createUser, getUserByID, getUsers, deleteUser, updateUser };
