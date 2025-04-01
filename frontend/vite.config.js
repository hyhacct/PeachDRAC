import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';



// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
    ],

    build: {
        outDir: 'dist' // 确保与 wails.json 的 outputdir 一致
    },

    // 定义基础路径，用于开发和生产环境
    base: '/',


  // 定义开发服务器的配置
  server: {
    port: 5174,
    host: 'localhost',
  },

    // 定义模块解析的配置
    resolve: {
        // 定义模块的别名
        alias: {
            '@': path.resolve(__dirname, './src'),
            '@wails': path.resolve(__dirname, './wailsjs')
        },
        // 定义模块解析时的扩展名
        extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json'],
        // 定义模块解析时的条目
        mainFields: ['browser', 'module', 'main']
    },
})
