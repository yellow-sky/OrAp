import { defineStore } from 'pinia'
import {api} from "boot/axios";

export const useAuthStore = defineStore('auth', {
  state: () => ({
    authenticated: false,
    jwt: '',
  }),

  getters: {
    isAuthenticated (state) {
      return state.authenticated
    }
  },

  actions: {
    async login (username, password) {
      const thisStore = this
      return await api.post('/api/auth/token', {}, {auth: {username: username, password: password}})
        .then(function (resp) {
          const jwt = resp.data.data
          thisStore.jwt = jwt
          api.defaults.headers.common = {'Authorization': `Bearer ${jwt}`}
          thisStore.authenticated = true
          return true
        })
        .catch(function (error) {
          thisStore.jwt = ''
          api.defaults.headers.common = {'Authorization': ''}
          thisStore.authenticated = false
          console.log('[Error on login] ', error)
          //if (error.response.status === 401) {
          //}
          return false
        })
    },
    logout() {
      this.authenticated = false
      this.jwt = ''
      api.defaults.headers.common = {'Authorization': ''}
      this.router.push({name: 'Login'})
    },
    initInterceptor() {
      const thisStore = this
      api.interceptors.response.use(function (response) {
        return response
      }, function (error) {
        console.log("[Error interceptor] " + error.response.data)
        if (error.response.status === 401) {
          thisStore.logout()
        }
        return Promise.reject(error)
      })
    }
  }
})
