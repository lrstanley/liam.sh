<script setup lang="ts">
import { version as vueVersion } from "vue"

const { data: version } = await getServiceVersion({ composable: "useFetch" })
</script>

<template>
  <div
    class="flex flex-col flex-auto max-w-full max-h-full border rounded-sm size-full bg-zinc-900 border-zinc-700/50"
  >
    <slot />

    <div class="box-border flex border-t bottom-bar border-zinc-700/50" v-auto-animate>
      <UTooltip
        v-if="version"
        :content="{ side: 'top' }"
        :delay-duration="0"
        :text="'build date: ' + version.build_date"
      >
        <span class="items-center bar-item misc">
          <UIcon name="logos:visual-studio-code" class="mr-1 align-middle" />
        </span>
      </UTooltip>

      <UPopover mode="hover" :open-delay="150" :content="{ side: 'top', align: 'start', sideOffset: 0 }">
        <span class="items-center bar-item">
          <UIcon name="mdi:source-branch" />
          {{ branchMenuOptions.find((o) => o.to === $route.path)?.label }}
        </span>

        <template #content>
          <div class="p-[1px]">
            <ul class="flex flex-col flex-nowrap">
              <li v-for="link in branchMenuOptions" :key="link.label">
                <router-link v-if="link.to" v-slot="{ isActive, href, navigate }" :to="link.to" custom>
                  <a
                    :href="href"
                    class="flex items-center px-1 rounded-sm hover:bg-zinc-900/50 hover:text-(--ui-text) text-(--ui-text-muted)"
                    :class="{ 'bg-zinc-900/50 text-(--ui-text)!': isActive }"
                    @click="navigate"
                  >
                    <UIcon name="mdi:source-branch" />
                    <span>{{ link.label }}</span>
                  </a>
                </router-link>
              </li>
            </ul>
          </div>
        </template>
      </UPopover>

      <div v-if="version" class="items-center bar-item misc" data-allow-mismatch>
        <UIcon name="logos:gopher" />
        <span>{{ version.go_version.replace(/^go/, "") }}</span>
      </div>
      <div class="items-center bar-item misc">
        <UIcon name="logos:vue" />
        {{ vueVersion }}
      </div>

      <slot name="footer">
        <div class="ml-auto" />
      </slot>

      <UTooltip text="... or just gofmt" :content="{ side: 'top' }" :delay-duration="0">
        <span class="bar-item misc">spaces:4</span>
      </UTooltip>

      <GithubStats placement="top-end" />
    </div>
  </div>
</template>

<style>
@reference "../../assets/main.css";

.bottom-bar {
  line-height: 1.4;
  @apply text-[1.15rem] md:text-[0.95rem];
}

.bar-item {
  @apply px-2 rounded-br-sm inline-flex text-(--ui-text-muted) align-middle cursor-pointer transition hover:bg-zinc-800 text-sm;
}

.bar-item .iconify {
  @apply align-middle mr-1;
}

.bar-item.misc {
  @apply hidden sm:inline-flex;
}
</style>
