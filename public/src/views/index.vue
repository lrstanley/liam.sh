<template>
  <LayoutBase>
    <div class="flex flex-row flex-auto justify-center items-stretch md:items-center">
      <div
        class="flex flex-col flex-auto max-w-3xl md:max-h-xl pt-7 md:py-7 items-stretch md:items-center"
      >
        <!-- <div class="mb-4 text-size-38px md:text-size-45px">
          <span class="inline-flex mr-10px text-green-600">
            {{ state.base.githubUser.name.split(" ")[0].toLowerCase() }}
            <span class="text-gray-500">:</span>
            ~
            <span class="text-gray-500 mr-4">$</span>
            #
          </span>
          <span v-motion-slide-right class="text-gradient from-teal-500 to-violet-600">Hello!</span>
          <span class="text-size-0.9em text-gray-600 animate-pulse">▌</span>
        </div> -->
        <CoreTerminal class="mb-4 text-size-38px md:text-size-45px" path="~" prefix="#" value="Hello!" />
        <n-card content-style="padding: 0;display: flex;flex-direction: column" class="p-0 flex-auto">
          <n-space v-motion-slide-left vertical :size="12" class="p-5 flex-auto">
            <n-alert title="Default Text" type="info"> Gee it's good to be back home </n-alert>
            <n-alert title="Info Text" type="info"> Gee it's good to be back home </n-alert>
            <n-alert title="Success Text" type="success">
              Leave it till tomorrow to unpack my case
            </n-alert>
          </n-space>

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
              <router-link :to="'admin'">
                <span class="hidden bar-item misc">sudo</span>
              </router-link>
              <n-tooltip trigger="hover">
                <template #trigger>
                  <a
                    class="px-2 rounded-br-sm bg-blue-600 hover:bg-blue-800 hover:text-current transition"
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
  </LayoutBase>
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
