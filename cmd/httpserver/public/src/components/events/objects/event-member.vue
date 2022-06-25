<template>
  <div>
    <span v-if="action == 'edited'"> updated collaborator permissions on </span>
    <span v-else>
      {{ action }}

      <EventLink :href="member.html_url" :title="member.login" />

      as collaborator on
    </span>

    <EventLink :href="repo.name" />
  </div>
</template>

<script setup>
const props = defineProps({
  event: {
    type: Object,
    required: true,
  },
})

const repo = ref(props.event.repo)
const action = ref(props.event.payload.action)
const member = ref(props.event.payload.member)
</script>
