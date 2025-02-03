import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueLuckyCanvas from '@lucky-canvas/vue'

const app = createApp(App)
app.use(router)
app.use(VueLuckyCanvas)
app.mount('#app')
