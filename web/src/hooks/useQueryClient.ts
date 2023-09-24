import { QueryClient } from '@tanstack/react-query';

export const useQueryClient = () => {
  return new QueryClient({
    defaultOptions: {
      queries: {
        refetchOnMount: true,
        refetchOnWindowFocus: true,
        refetchOnReconnect: true,
        refetchInterval: false,
        retry: false,
        staleTime: 1000 * 60 * 5, // 5 minutes
        suspense: false,
      },
    },
  });
};
