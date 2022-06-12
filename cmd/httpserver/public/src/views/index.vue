<template>
  <!-- TODO: use less flex -->
  <div class="flex flex-row flex-auto justify-center items-stretch md:items-center">
    <div
      class="flex flex-col flex-auto max-w-3xl md:max-h-xl pt-7 md:py-7 items-stretch md:items-center"
    >
      <CoreNavigation />
      <CoreTerminal
        class="md:mt-5 mb-4 text-size-38px md:text-size-45px"
        path="~"
        prefix="#"
        value="Hello!"
      />
      <n-card content-style="padding: 0;display: flex;flex-direction: column" class="p-0 flex-auto">
        <EventsRender class="h-full md:h-300px overflow-x-auto" />

        <n-layout-footer bordered class="bottom-bar">
          <span v-motion-fade class="flex flex-auto">
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item">
                  <n-icon class="align-middle">
                    <i-mdi-source-branch />
                  </n-icon>
                  {{ state.base.version?.commit?.substring(0, 8) || "master" }}
                </span>
              </template>
              {{
                state.base.version
                  ? `${state.base.version.goVersion} &middot; built: ${state.base.version.date}`
                  : "master"
              }}
            </n-tooltip>
            <span class="ml-auto" />
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item misc">spaces:4</span>
              </template>
              ... or just gofmt
            </n-tooltip>
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="hidden bar-item misc">go/vue</span>
              </template>
              built with Go and Vue.js
            </n-tooltip>
            <n-tooltip trigger="hover">
              <template #trigger>
                <a
                  class="px-2 rounded-br-sm bg-blue-600 hover:bg-blue-800 hover:text-current transition"
                  style="color: white !important"
                  :href="state.base.githubUser.htmlurl"
                >
                  <n-icon class="align-middle">
                    <i-mdi-github />
                  </n-icon>
                  @{{ state.base.githubUser.login }}
                </a>
              </template>
              <p>
                <n-icon class="align-middle"><i-mdi-github /></n-icon>
                {{ state.base.githubUser.name }} &middot; {{ state.base.githubUser.bio }}
              </p>
            </n-tooltip>
          </span>
        </n-layout-footer>
      </n-card>
    </div>
  </div>
</template>

<script setup>
const state = useState()
</script>

<style scoped>
.bottom-bar {
  line-height: 1.5;
}

.bar-item {
  @apply pl-1.5 pr-2 rounded-br-sm inline-flex text-gray-400 align-middle cursor-pointer transition hover: bg-gray-800;
}

.bar-item.misc {
  @apply <sm: hidden sm:inline-flex;
}
</style>
