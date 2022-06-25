<template>
  <!-- TODO: use less flex -->
  <div class="flex flex-auto justify-center items-stretch lg:items-center h-full w-full">
    <div
      class="flex flex-auto flex-col h-full w-full lg:max-w-3xl lg:max-h-md basis-0 flex-grow-0 flex-shrink items-stretch md:items-center pt-15px"
    >
      <CoreNavigation />
      <CoreTerminal
        class="mb-4 text-size-38px md:text-size-45px flex flex-auto justify-center"
        path="~"
        prefix="#"
        value="Hello!"
      />
      <n-card
        content-style="padding: 0;display: flex;flex-direction: column;"
        class="p-0 flex flex-auto h-full w-full"
      >
        <EventsRender class="overflow-x-auto basis-0 flex-grow flex-shrink h-full w-full" />

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
  @apply <md:(text-size-1.2em leading-6);
}

.bar-item {
  @apply pl-1.5 pr-2 rounded-br-sm inline-flex text-gray-400 align-middle cursor-pointer transition hover: bg-gray-800;
}

.bar-item.misc {
  @apply <sm: hidden sm:inline-flex;
}
</style>
