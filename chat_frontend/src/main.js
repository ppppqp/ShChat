import { createApp } from 'vue'
import App from './App.vue'
import { Quasar } from 'quasar'

const app = createApp(App)
app.use(Quasar, {
    plugins: {}, // import Quasar plugins and add here
  })
app.mount('#app')
