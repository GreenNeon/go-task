import axios from 'axios'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// import './assets/main.css'

const app = createApp(App)

app.use(router)

axios.defaults.baseURL = 'http://localhost:8080';

// Important: If axios is used with multiple domains, the AUTH_TOKEN will be sent to all of them.
// See below for an example using Custom instance defaults instead.
axios.defaults.headers.common['Authorization'] = 'Bearer ' +  sessionStorage.getItem('user');
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
// Add a response interceptor
axios.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
  }, function (error) {
    if (error.response && error.response.status === 401) {
        router.push('/login')
    }
    return Promise.reject(error);
  });

app.provide('user', sessionStorage.getItem('user'))

app.mount('#app')
