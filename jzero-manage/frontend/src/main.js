import { createApp } from 'vue'
import naive from 'naive-ui'
import App from './App.vue'
import pinia from './stores'

const app = createApp(App)

app.use(naive)
app.use(pinia)

app.mount('#app')
