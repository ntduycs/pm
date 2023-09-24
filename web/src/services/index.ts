import Axios from 'axios';

export const axios = Axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 3000,
});

export * from './member.service';
