<script setup>
import axios from "axios"
import { reactive, ref } from "vue"
import { useRoute } from "vue-router"
import router from "../router"

const taskId = useRoute().params.id
const task = ref(null)

task.value = await (await axios.get(`/api/tasks/${taskId}`)).data.data

const rawDate = new Date(task.value.deadline)
let ye = new Intl.DateTimeFormat('en', { year: 'numeric' }).format(rawDate)
let mo = new Intl.DateTimeFormat('en', { month: '2-digit' }).format(rawDate)
let da = new Intl.DateTimeFormat('en', { day: '2-digit' }).format(rawDate)
const dateFormatted = `${ye}-${mo}-${da}`

const form = reactive({
    name: null,
    assignee_id: null,
    deadline: null,
    ...task.value,
    deadline: dateFormatted
})

const users = ref([])
const errorMessage = ref(false)
const successMessage = ref(false)

axios.get('/api/users').then(res => {
    users.value.push(...res.data.data)
}).catch(reason => {
    console.log(reason)
})

const onUpdate = () => {
    errorMessage.value = false
    axios.patch(`/api/tasks/${form.ID}`, { name: form.name, deadline: (new Date(form.deadline)).toISOString(), assignee_id: Number.parseInt(form.assignee_id) })
        .then(res => {
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
        <h2 class="mt-2">Update Task</h2>
        <form class="form-horizontal">
            <div class="form-group">
                <label class="form-label" for="input-example-1">Task</label>
                <input required v-model="form.name" class="form-input" type="text" id="input-example-1"
                    placeholder="Name">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Deadline</label>
                <input required v-model="form.deadline" class="form-input" type="date" id="input-example-1"
                    placeholder="Deadline">
            </div>

            <div class="form-group">
                <label class="form-label" for="input-example-1">Assign task to</label>
                <select required v-model="form.assignee_id" class="form-select">
                    <option v-for="user in users" :key="user.ID" :value="user.ID">{{ user.username }}</option>
                </select>
            </div>

            <div v-if="successMessage !== false" class="mb-2 toast toast-success" style="width: 100%">
                <button class="btn btn-clear float-right" @click="successMessage = false"></button>
                {{ successMessage }}
            </div>

            <div v-if="errorMessage !== false" class="mb-2 toast toast-error" style="width: 100%">
                <button class="btn btn-clear float-right" @click="errorMessage = false"></button>
                {{ errorMessage }}
            </div>
            <button class="btn btn-primary" @click.prevent="onUpdate">Update</button>
        </form>
    </main>
</template>
