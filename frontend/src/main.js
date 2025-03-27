import { createApp } from 'vue'
import naive from 'naive-ui'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { Icon } from '@iconify/vue';

const app = createApp(App)

app.use(naive)
app.use(router)
app.use(ElementPlus)
app.component('Icon', Icon)
app.mount('#app')
