import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue({
        template: {
            compilerOptions: {
                // 将所有以md-起始标签名都视为自定义元素
                isCustomElement: (tag) => tag.startsWith('md-')
            }
        }
    })],

    server: {
        host: '0.0.0.0'
    }
})
