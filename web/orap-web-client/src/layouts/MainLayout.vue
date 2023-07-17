<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-btn dense flat round icon="menu" @click="toggleLeftDrawer" />
        <q-toolbar-title>
          <q-avatar>
            <img src="~assets/logo_orap.png" />
          </q-avatar>
          OrAP
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <instance-info />
      <q-splitter horizontal />
      <q-list>
        <MenuItem
          v-for="item in menuItems"
          :key="item.title"
          v-bind="item"
        />
      </q-list>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from 'vue'
import MenuItem from "components/MenuItem.vue";
import InstanceInfo from "components/InstanceInfo.vue";

const menuItems = [
  {
    title: 'Dashboard',
    icon: 'dashboard',
    linkName: 'Dashboard'
  },
  {
    title: 'Devices',
    icon: 'lan',
    linkName: 'Devices'
  },
  {
    title: 'Logout',
    icon: 'logout',
    linkName: 'Logout'
  }
]

export default defineComponent({
  name: 'MainLayout',

  components: {
    InstanceInfo,
    MenuItem,
  },

  setup () {
    const leftDrawerOpen = ref(false)

    return {
      menuItems: menuItems,
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
})
</script>
