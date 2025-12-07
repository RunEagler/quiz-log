import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import relay from 'vite-plugin-relay'

export default defineConfig({
  plugins: [react(), relay],
  server: {
    proxy: {
      '/query': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
