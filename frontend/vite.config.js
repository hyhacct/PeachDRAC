import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';



// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
    ],

    // 定义基础路径，用于开发和生产环境
    base: '/',

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
