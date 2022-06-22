<template>
  <span>
    <n-tag v-if="repo.owner.login != state.base.githubUser.login" type="info" size="small">
      maintainer
    </n-tag>
    <n-tag v-if="repo.fork" type="error" size="small">
      <template #icon>
        <n-icon><i-mdi-source-fork /></n-icon>
      </template>
      fork
    </n-tag>
    <n-tag v-if="repo.archived" type="warning" size="small">
      <template #icon>
        <n-icon><i-mdi-archive-outline /></n-icon>
      </template>
      archived
    </n-tag>
    <n-tag v-if="repo.license" size="small" :title="repo.license.name">
      <template #icon>
        <n-icon><i-mdi-scale-balance /></n-icon>
      </template>
      {{ repo.license.key }}
    </n-tag>
  </span>
</template>

<script setup>
const props = defineProps({
  value: {
    type: Object,
    required: true,
  },
})

const state = useState()
const repo = ref(props.value)
</script>
