import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [react()],
  build: {
    manifest: true,
    rollupOptions: {
      input: {
        islands: 'src/islands.tsx', // islands bundle
        app: 'src/app.tsx', // SPA bundle
      },
    },
  },
})
