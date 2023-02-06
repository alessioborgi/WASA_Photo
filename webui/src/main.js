import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';

import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import CardProfile from './components/CardProfile.vue'

import './assets/dashboard.css'
import './assets/main.css'
import './assets/login.css'
import './assets/myProfile.css'


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("CardProfile", CardProfile);
app.use(router)
app.mount('#app')




