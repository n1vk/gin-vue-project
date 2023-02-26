import {createApp} from 'vue'
import App from './App.vue'
import axios from "axios";
import {router} from './router'
import './style.css'

axios.defaults.baseURL = 'http://localhost:33445'

const app = createApp(App)
    .use(router)
    .mount('#app')

