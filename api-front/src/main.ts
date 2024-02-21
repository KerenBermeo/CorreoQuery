import './assets/main.css'
import router from './router/index' // Importa el router desde tu archivo router/index.ts

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'

const app = createApp(App)

app.use(createPinia())
app.use(router) // Usa el router en tu aplicación

app.mount('#app')
