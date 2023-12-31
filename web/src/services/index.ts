import Axios, { AxiosError } from 'axios';
import qs from 'query-string';

export const axios = Axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 3000,
});

axios.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error instanceof AxiosError && error.code === 'ECONNABORTED') {
      return Promise.reject(new Error('Request timed out'));
    } else if (error.response?.status === 401) {
      return Promise.reject(new Error('Unauthorized'));
    } else if (error.response?.status === 404) {
      return Promise.reject(new Error('Not found'));
    }

    return Promise.reject(error);
  },
);

// interceptors log request params
axios.interceptors.request.use((config) => {
  console.log('Request', config.method?.toUpperCase(), config.url, qs.stringify(config.params));
  return config;
});

export * from './member.service';
export * from './effort-allocation.service';
export * from './papc.service';
