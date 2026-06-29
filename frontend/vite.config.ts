import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [react()],
  build: {
    emptyOutDir: true,
    manifest: true,
    rollupOptions: {
      input: {
        main: 'src/main.tsx',
      },
      output: {
        dir: '../backend/internal/webassets/dist',
      },
    },
  },
})
