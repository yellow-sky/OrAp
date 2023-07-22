<template>
  <q-page class="q-pa-md">
    <q-table
      flat
      title="Network Devices"
      :rows="devices.getDevices"
      :columns="columns"
      row-key="id"
      hide-pagination
      virtual-scroll
    />
  </q-page>
</template>

<script>

import {useNetworkDevicesStore} from "stores/network_devices";
import {defineComponent} from "vue";

const columns = [
  { name: 'id', label: 'ID', align: 'left', field: 'id', sortable: true },
  { name: 'interface', align: 'center', label: 'IFace', field: 'interface', sortable: true },
  { name: 'type', align: 'center', label: 'Type', field: 'type', sortable: true },
  { name: 'driver', label: 'Driver', field: 'driver', sortable: true },
  { name: 'state', label: 'State', field: 'state' },
]

const devicesStore = useNetworkDevicesStore()

export default defineComponent({
  name: 'NetworkDeviceTable',
  setup() {
    devicesStore.fetchDevices()
    return {
      columns: columns,
      devices: devicesStore,
    }
  },
  mounted() {
    this.updateDeviceInterval = setInterval(() => {
      devicesStore.fetchDevices()
    }, 1000)
  },
  unmounted() {
    clearInterval(this.updateDeviceInterval)
  }

})
</script>

<style scoped>

</style>
