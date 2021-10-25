import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import installElementPlus from './plugins/element'
import axios from 'axios'
import './assets/css/login.css'

axios.defaults.baseURL = "http://localhost:9090"

const app = createApp(App)
    //挂载http
axios.defaults.timeout = 50000;
app.config.globalProperties.$http = axios

installElementPlus(app)
app.use(store).use(router).mount('#app')