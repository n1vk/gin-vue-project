import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import axios from "axios";

axios.defaults.baseURL = 'http://localhost:33445'

createApp(App).mount('#app')
