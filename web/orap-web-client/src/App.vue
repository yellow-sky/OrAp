<template>
  <router-view />
</template>

<script>
import { defineComponent } from 'vue'
import { useAuthStore } from 'stores/auth'

export default defineComponent({
  name: 'App',
  created() {
    // Show Login page
    const auth = useAuthStore()
    auth.initInterceptor()
    auth.logout()
  }
})
</script>
