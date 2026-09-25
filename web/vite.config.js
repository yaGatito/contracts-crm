import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 5173,
    proxy: {
      '/persons': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/legal-persons': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/contracts': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  preview: {
    host: '0.0.0.0',
    port: 4173,
  },
})
