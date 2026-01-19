<script setup lang="ts">
definePageMeta({
  title: "Home",
  layout: "terminal",
})

const githubUser = useGithubUser()
const { data: version } = await useApi('/version')
</script>

<template>
  <ContainerIde>
    <EventsRender class="relative overflow-x-hidden overflow-y-scroll size-full grow basis-0" />

    <template #footer>
      <div v-if="version" class="items-center bar-item misc" data-allow-mismatch>
        <UIcon name="logos:gopher" />
        <span>{{ version.go_version.replace(/^go/, "") }}</span>
      </div>

      <CodingStats placement="top-start" />

      <div class="ml-auto" />

      <ULink v-if="githubUser?.html_url" class="items-center bar-item misc" :href="githubUser.html_url + '.gpg'"
        target="_blank">
        <UIcon name="lucide:key-round" class="text-orange-600" />
        gpg
      </ULink>

      <div v-if="githubUser?.location" class="items-center bar-item misc">
        <UIcon name="lucide:map-pin" />
        {{ githubUser.location }}
      </div>
    </template>
  </ContainerIde>
</template>
