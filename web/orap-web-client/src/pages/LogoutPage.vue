<template>
  <q-page class="flex flex-center ">
    <div class="q-pa-md">
      <div class="q-gutter-md row">
        <q-spinner
          color="primary"
          size="3em"
        />
      </div>
    </div>
  </q-page>
</template>

<script>
import { defineComponent } from 'vue'
import {useAuthStore} from "stores/auth";

export default defineComponent({
  name: 'LogoutPage',
  setup() {
    const auth = useAuthStore()
    return {
      auth: auth,
    }
  },
  mounted() {
    this.$q.notify({message: 'Logged out', position: 'top', type: 'info', progress: true, timeout: 1000})
    this.auth.logout()
  },
})
</script>
