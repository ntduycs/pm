import { BuildOptions, defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import * as path from 'path';

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  const buildOptions: BuildOptions = {
    chunkSizeWarningLimit: 2000,
    reportCompressedSize: mode === 'development',
  };

  return {
    plugins: [react()],
    resolve: {
      alias: [{ find: '@pm', replacement: path.resolve(__dirname, 'src') }],
    },
    server: {
      port: 3000,
      strictPort: true,
      host: true,
    },
    build: buildOptions,
  };
});
