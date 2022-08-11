<script setup>
import { RouterLink, RouterView } from 'vue-router'

import { inject } from 'vue'
const user = inject('user')
</script>

<template>
  <div class="container grid-xl p-0">
    <div class="columns col-gapless">
      <div class="column hide-md col-2"></div>
      <div class="column col-md-12 col-8">

        <header class="navbar p-2">
          <section class="navbar-section">
            <router-link to="/" class="text-bold navbar-brand mr-2">Task.Go</router-link>
            <router-link to="/" class="btn btn-link">List</router-link>
            <router-link to="/create" class="btn btn-link">New <i class="icon icon-plus"></i></router-link>
          </section>

          <section class="navbar-section">
            <router-link v-if="user === null" to="/login" class="btn btn-primary">Login</router-link>
            <router-link v-else to="/logout" class="btn btn-error ">Logout</router-link>
          </section>
        </header>

        <div class="" style="height: 100%">
          <Suspense>
            <RouterView />
          </Suspense>

        </div>
      </div>

      <div class="column hide-md col-2"></div>
    </div>
  </div>
</template>
