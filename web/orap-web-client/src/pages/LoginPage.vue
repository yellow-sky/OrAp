<template>
  <q-page class="flex flex-center ">
<!--    <q-page class="bg-light-green window-height window-width row justify-center items-center">-->
      <div class="column">
        <div class="row">
          <h5 class="text-h5 text-white q-my-md">Company & Co</h5>
        </div>
        <div class="row">
          <q-card square bordered class="q-pa-lg shadow-1">
            <q-card-section>
              <div class="text-h6">Sign In</div>
            </q-card-section>
            <q-card-section>
              <q-form class="q-gutter-md" @keydown.enter.prevent="login">
                <q-input square filled v-model="username" type="text" label="Username" autofocus/>
                <q-input square filled v-model="password" type="password" label="Password" />
              </q-form>
            </q-card-section>
            <q-card-actions class="q-px-md">
              <q-btn size="lg" class="full-width" label="Login" @click="login" />
            </q-card-actions>
          </q-card>
        </div>
      </div>
  </q-page>
</template>

<script>
import { defineComponent } from 'vue'
import {useAuthStore} from "stores/auth";

export default defineComponent({
  name: 'LoginPage',
  data () {
    return {
      username: '',
      password: '',
    }
  },
  setup() {
    const auth = useAuthStore()
    return {
      auth: auth,
    }
  },
  methods: {
    async login() {
      const thisComp = this
      this.auth.login(this.username, this.password).then(function (data) {
        if(data) {
          thisComp.$q.notify({message: 'Logged in', position: 'top', type: 'info', progress: true, timeout: 1000})
          thisComp.$router.push({name: 'Dashboard'})
        } else {
          thisComp.$q.notify({message: 'Login failed', position: 'top', type: 'negative', progress: true, timeout: 1000})
        }
      })
    },
  }
})
</script>
