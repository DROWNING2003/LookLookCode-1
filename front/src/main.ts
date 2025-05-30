import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import './style.css'
import 'vant/lib/index.css'
import App from './App.vue'
import router from './router'
import '@fortawesome/fontawesome-free/css/all.min.css'

const app = createApp(App)
const pinia = createPinia()

// 配置 Pinia 持久化插件
pinia.use(piniaPluginPersistedstate)

// Setup Pinia
app.use(pinia)

// Setup Router
app.use(router)

app.mount('#app')
