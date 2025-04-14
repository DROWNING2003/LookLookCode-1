import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'
import Components from 'unplugin-vue-components/vite'
import MotionResolver from 'motion-v/resolver'

// https://vite.dev/config/
import path from 'path'

export default defineConfig({
  server:{
    port:8081
  },
  plugins: [vue(),tailwindcss(),Components({
    dts: true,
    resolvers: [
      MotionResolver()
    ],
  }),],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src')
    }
  }
})
