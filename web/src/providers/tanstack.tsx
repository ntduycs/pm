import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { PropsWithChildren } from 'react';
import { useQueryClient } from '@pm/hooks';

// const queryClient = new QueryClient({
//   defaultOptions: {
//     queries: {
//       refetchOnMount: true,
//       refetchOnWindowFocus: true,
//       refetchOnReconnect: true,
//       refetchInterval: false,
//       retry: false,
//       staleTime: 1000 * 60 * 5, // 5 minutes
//       suspense: false,
//     },
//   },
// });

export const TanstackProvider = ({ children }: PropsWithChildren) => {
  const queryClient = useQueryClient();
  return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};
