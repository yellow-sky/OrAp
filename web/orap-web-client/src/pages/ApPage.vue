<template>
  <network-device-table />
</template>

<script>
import { defineComponent } from 'vue'
import NetworkDeviceTable from "components/NetworkDevicesTable.vue";


export default defineComponent({
  name: 'ApPage',
  components: {NetworkDeviceTable},
})
</script>
