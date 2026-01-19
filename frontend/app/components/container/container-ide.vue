<script setup lang="ts">
const route = useRoute()
const { width } = useWindowSize()

const { data: version } = await useApi('/version')
const label = branchMenuOptions.find((o) => o.to === route.path)?.label
</script>

<template>
  <div class="flex flex-col flex-auto max-w-full max-h-full border rounded-sm size-full bg-zinc-900 border-zinc-700/50">
    <slot />

    <motion as="div" layout :initial="{ opacity: 0 }" :animate="{ opacity: 1 }"
      class="box-border flex border-t bottom-bar border-zinc-700/50">
      <UTooltip v-if="version" :content="{ side: 'top' }" :delay-duration="0"
        :text="'build date: ' + version.build_date">
        <span class="items-center bar-item">
          <UIcon name="logos:visual-studio-code" class="mr-1 align-middle" />
        </span>
      </UTooltip>

      <UPopover :mode="width <= 640 ? 'click' : 'hover'" :open-delay="150"
        :content="{ side: 'top', align: 'start', sideOffset: 0 }">
        <span class="items-center bar-item">
          <UIcon name="mdi:source-branch" />
          {{ label }}
        </span>

        <template #content>
          <div class="p-px">
            <ul class="flex flex-col flex-nowrap">
              <li v-for="link in branchMenuOptions" :key="link.label">
                <router-link v-if="link.to" v-slot="{ isActive, href, navigate }" :to="link.to" custom>
                  <a :href="href"
                    class="flex items-center px-1 rounded-sm hover:bg-zinc-900/50 hover:text-default text-muted"
                    :class="{ 'bg-zinc-900/50 text-default!': isActive }" @click="navigate">
                    <UIcon name="mdi:source-branch" />
                    <span>{{ link.label }}</span>
                  </a>
                </router-link>
              </li>
            </ul>
          </div>
        </template>
      </UPopover>

      <slot name="footer">
        <div class="ml-auto" />
      </slot>

      <GithubStats placement="top-end" />
    </motion>
  </div>
</template>

<style>
@reference "../../assets/main.css";

.bottom-bar {
  line-height: 1.4;
  @apply text-[1.15rem] md:text-[0.95rem];
}

.bar-item {
  @apply px-2 rounded-br-sm inline-flex text-muted align-middle cursor-pointer transition hover:bg-zinc-800 text-sm max-sm:py-1;
}

.bar-item .iconify {
  @apply align-middle mr-1;
}

.bar-item.misc {
  @apply hidden md:inline-flex;
}
</style>
