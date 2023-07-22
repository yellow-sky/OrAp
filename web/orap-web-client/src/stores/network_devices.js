import { defineStore } from 'pinia'
import { api } from 'boot/axios'


export const useNetworkDevicesStore = defineStore('network-devices', {
  state: () => ({
    networkDevices: []
  }),

  getters: {
    getDevices (state) {
      return state.networkDevices
    }
  },

  actions: {
    setDevices (newDevices) {
      this.networkDevices = newDevices
    },
    async fetchDevices() {
      const thisStore = this
      await api.get('/api/devices')
        .then(function (resp) {
          const defs = resp.data.data
          defs.sort((a, b) => a.order - b.order)
          thisStore.networkDevices = defs
        })
        .catch(function (error) {
          console.log('Error on fetch network devices:', error)
        })
    }
  }
})
