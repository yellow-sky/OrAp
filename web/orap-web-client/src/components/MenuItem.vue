<template>
    <q-item
      clickable
      @click="$router.push({name: linkName})"
      :active="$route.name === linkName"
    >
      <q-item-section
        v-if="icon"
        avatar
      >
        <q-icon :name="icon" />
      </q-item-section>

      <q-item-section>
        <q-item-label>{{ title }}</q-item-label>
      </q-item-section>
    </q-item>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'MenuItem',
  props: {
    title: {
      type: String,
      required: true
    },
    linkName: {
      type: String,
      default: '#'
    },
    icon: {
      type: String,
      default: ''
    }
  },
})
</script>
