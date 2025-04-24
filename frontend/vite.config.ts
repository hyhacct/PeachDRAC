import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import * as path from 'path';


// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    build: {
        rollupOptions: {
            // disable hash suffix for output files
            output: {
                entryFileNames: `assets/[name].js`,
                chunkFileNames: `assets/[name].js`,
                assetFileNames: `assets/[name].[ext]`,
            },
            external: ['#minpath'], // 告诉 Rollup 将 #minpath 视为外部模块
        },
    },


    // 定义模块解析的配置
    resolve: {
        // 定义模块的别名
        alias: {
            '@wails': path.resolve(__dirname, './wailsjs'),
            '@': path.resolve(__dirname, './src'),
        },
        // 定义模块解析时的扩展名
        extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json'],
        // 定义模块解析时的条目
        mainFields: ['browser', 'module', 'main']
    },
})
