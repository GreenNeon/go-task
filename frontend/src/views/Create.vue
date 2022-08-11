<script setup>
import axios from "axios"
import { reactive, ref } from "vue"
import router from "../router"
const form = reactive({
    name: null,
    assignee_id: null,
    deadline: null
})
const users = ref([])
const errorMessage = ref(false)
const successMessage = ref(false)

axios.get('/api/users').then(res => {
    users.value.push(...res.data.data)
}).catch(reason => {
    console.log(reason)
})

const onCreate = () => {
    errorMessage.value = false
    axios.post('/api/tasks', {...form, assignee_id: Number.parseInt(form.assignee_id), deadline: (new Date(form.deadline)).toISOString()})
    .then(res => {
        form.name = ''
        form.assignee_id = 1
        form.deadline = ''
        successMessage.value = res.data.message
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
        <h2 class="mt-2">Create New Task</h2>
        <form class="form-horizontal">
            <div class="form-group">
                <label class="form-label" for="input-example-1">Task</label>
                <input required v-model="form.name" class="form-input" type="text" id="input-example-1" placeholder="Name">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Deadline</label>
                <input required v-model="form.deadline" class="form-input" type="date" id="input-example-1" placeholder="Deadline">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Assign task to</label>
                <select required v-model="form.assignee_id" class="form-select" >
                    <option v-for="user in users" :key="user.ID" :value="user.ID">{{user.username}}</option>
                </select>
            </div>
            

            <div v-if="successMessage !== false" class="mb-2 toast toast-success" style="width: 100%">
                <button class="btn btn-clear float-right" @click="successMessage = false"></button>
                {{ successMessage }}
            </div>

            <div v-if="errorMessage !== false" class="mb-2 toast toast-error" style="width: 100%">
                <button class="btn btn-clear float-right"  @click="errorMessage = false"></button>
                {{errorMessage}}
            </div>
            <button class="btn btn-primary" @click.prevent="onCreate">Create</button>
        </form>
    </main>
</template>
