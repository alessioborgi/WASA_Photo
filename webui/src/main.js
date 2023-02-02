import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';

import ErrorMessage from './components/ErrorMessage.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'
import './assets/login.css'
import './assets/myProfile.css'


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMessage", ErrorMessage);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')




