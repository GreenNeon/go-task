<script setup>
import axios from "axios"
import { reactive, ref } from "vue"
import router from "../router"
const form = reactive({
    username: '',
    password: ''
})
const errorMessage = ref(false)
const successMessage = ref(false)

const onLogin = () => {
    errorMessage.value = false
    axios.post('/api/register', form)
    .then(re => {
        form.password = ''
        form.username = ''
        successMessage.value = re.data.message + ' redirect in 5s ...'
        setTimeout(() => router.push('/login'), 5000)
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
        <h2 class="mt-2">Sign up</h2>
        <form class="form-horizontal">
            <div class="form-group">
                <label class="form-label" for="input-example-1">Username</label>
                <input v-model="form.username" class="form-input" type="text" id="input-example-1" placeholder="username">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Password</label>
                <input v-model="form.password" class="form-input" type="password" id="input-example-1" placeholder="password">
            </div>
            
            <div v-if="successMessage !== false" class="toast toast-success" style="width: 100%">
                <button class="btn btn-clear float-right"></button>
                {{successMessage}}
            </div>
            <div v-if="errorMessage !== false" class="toast toast-error" style="width: 100%">
                <button class="btn btn-clear float-right"></button>
                {{errorMessage}}
            </div>
            <button class="btn btn-primary mt-2" @click.prevent="onLogin">Register</button>
        </form>
    </main>
</template>
