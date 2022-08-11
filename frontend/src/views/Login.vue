<script setup>
import axios from "axios"
import { reactive, ref } from "vue"
import router from "../router"
const form = reactive({
    username: '',
    password: ''
})
const errorMessage = ref(false)

const onLogin = () => {
    errorMessage.value = false
    axios.post('/api/login', form)
    .then(res => {
        sessionStorage.setItem('user', res.data.token)
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + sessionStorage.getItem('user');
        router.push('/')
    })
    .catch(re => {
        if (!!re.response === false) errorMessage.value = re.message
        if (re.response.data.error) {
            errorMessage.value = re.response.data.error
        } else errorMessage.value = re.response.data.message
    })
}
</script>

<template>
    <main>
        <h2 class="mt-2">Sign in</h2>
        <form class="form-horizontal">
            <div class="form-group">
                <label class="form-label" for="input-example-1">Username</label>
                <input v-model="form.username" class="form-input" type="text" id="input-example-1" placeholder="username">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Password</label>
                <input v-model="form.password" class="form-input" type="password" id="input-example-1" placeholder="password">
            </div>
            <div v-if="errorMessage !== false" class="toast toast-error" style="width: 100%">
                <button class="btn btn-clear float-right"></button>
                {{errorMessage}}
            </div>
            <button class="btn btn-primary mt-2" @click.prevent="onLogin">Login</button>
            <router-link class="btn btn-link ml-2 mt-2" to="/register">Register</router-link>
        </form>
    </main>
</template>
