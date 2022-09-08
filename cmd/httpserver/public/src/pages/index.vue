<template>
  <!-- TODO: use less flex -->
  <div class="flex items-stretch justify-center flex-auto w-full h-full lg:items-center">
    <div
      class="flex h-full w-full shrink grow-0 basis-auto flex-col items-stretch pt-[15px] md:items-center lg:max-h-[28rem] lg:max-w-3xl"
    >
      <CoreNavigation />
      <CoreTerminal
        class="mb-4 flex flex-auto justify-center text-[38px] md:text-[45px]"
        path="~"
        prefix="#"
        value="Hello!"
      />
      <n-card
        content-style="padding: 0;display: flex;flex-direction: column;"
        class="flex flex-auto w-full h-full p-0"
      >
        <EventsRender class="w-full h-full overflow-x-auto shrink grow basis-0" />

        <n-layout-footer bordered class="bottom-bar">
          <span v-motion-fade class="flex flex-auto">
            <n-tooltip trigger="hover">
              <template #trigger>
                <span class="bar-item">
                  <n-icon class="align-middle">
                    <i-mdi-source-branch />
                  </n-icon>
                  {{
                    state.base.version?.commit == "unknown"
                      ? "master"
                      : state.base.version?.commit?.substring(0, 8)
                  }}
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
                  class="px-2 transition bg-blue-600 rounded-br-sm hover:bg-blue-800 hover:text-current"
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

<script setup lang="ts">
const state = useState()
</script>

<style scoped>
.bottom-bar {
  line-height: 1.5;
  @apply text-[1.2em] md:text-[1em];
}

.bar-item {
  @apply pl-1.5 pr-2 rounded-br-sm inline-flex text-zinc-400 align-middle cursor-pointer transition hover:bg-zinc-800;
}

.bar-item.misc {
  @apply hidden sm:inline-flex;
}
</style>
