<script setup>
import { inject, reactive, ref } from 'vue'
import axios from 'axios'

const tasks = reactive([])
const errorMessage = ref(false)
const options = reactive({ weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
const user = inject('user')
const filters = reactive({
    mark: 'all',
    user: 'all'
})

axios.get('/api/tasks?mark=' + filters.mark + '&user=' + filters.user).then(res => {
    tasks.push(...res.data.data)
}).catch(reason => {
    console.log(reason)
})

const onRefresh = () => {
    axios.get('/api/tasks?mark=' + filters.mark + '&user=' + filters.user).then(res => {
        tasks.splice(0, tasks.length)
        tasks.push(...res.data.data)
    }).catch(reason => {
        console.log(reason)
    })
}

const onMark = (id, mark) => {
    axios.patch(`/api/tasks/${id}/mark`, { is_done: mark })
        .then(res => {
            onRefresh()
        })
        .catch(re => {
            if (!!re.response === false) errorMessage.value = re.message
            if (re.response.data.error) {
                errorMessage.value = re.response.data.error
            } else errorMessage.value = re.response.data.message
        })
}

const onDelete = (id, mark) => {
    axios.delete(`/api/tasks/${id}`)
        .then(res => {
            onRefresh()
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
        <div class="hero bg-dark" style="border-radius: 5px; padding: 1em">
            <div class="hero-body">
                <h1 class="m-0">Hello User, {{ user.username }}</h1>
                <p class="m-0">This is where you can see all the task you create and shared by others.</p>
                <p class="m-0">There's <b>total {{ tasks.length }} task/s</b></p>
            </div>
        </div>

        <div v-if="errorMessage !== false" class="my-2 toast toast-error" style="width: 100%">
            {{ errorMessage }}
        </div>

        <div class="my-2">
            <label class="form-radio form-inline">
                Filter Mark <select v-model="filters.mark" class="form-select">
                    <option value="all">All</option>
                    <option value="done">Done</option>
                    <option value="not_done">Not Done</option>
                </select>
            </label>

            <label class="form-radio form-inline">
                Filter User <select v-model="filters.user" class="form-select">
                    <option value="all">All</option>
                    <option value="self">Self</option>
                    <option value="other">Other Users</option>
                </select>
            </label>
            <label class="form-radio form-inline">
                <button class="btn" @click="onRefresh">Filter</button>
            </label>
        </div>

        <table class="table table-striped table-hover">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Assignee</th>
                    <th>Deadline</th>
                    <th>Mark</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr v-if="tasks.length === 0">
                    <td colspan="5"> Empty Tasks</td>
                </tr>
                <tr v-for="task in tasks" :key="task.ID" class="active">
                    <td>{{ task.name }}</td>
                    <td>{{ task.assignee.username }}</td>
                    <td>{{ new Date(task.deadline).toLocaleDateString("id-ID", options) }}</td>
                    <td>
                        <button v-if="task.is_done" class="btn btn-dark btn-sm" @click="onMark(task.ID, false)"><i
                                class="icon icon-check"></i></button>
                        <button v-else class="btn btn-dark btn-sm" @click="onMark(task.ID, true)"><i
                                class="icon icon-cross"></i></button>
                    </td>

                    <td>
                        <router-link :to="`/edit/${task.ID}`" class="btn btn-primary btn-sm"><i
                                class="icon icon-edit"></i></router-link>
                        <button class="btn btn-error btn-sm" @click="onDelete(task.ID)"><i
                                class="icon icon-delete"></i></button>
                    </td>
                </tr>
            </tbody>
        </table>
    </main>
</template>
