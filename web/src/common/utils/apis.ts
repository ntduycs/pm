import { AxiosError } from 'axios';

const getErrorDescription = (error: unknown) => {
  if (error instanceof AxiosError) {
    return error.response?.data?.message || error.message;
  } else if (error instanceof Error) {
    return error.message;
  } else {
    return 'Something went wrong!!!';
  }
};

export const Apis = {
  getErrorDescription,
};
