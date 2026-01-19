<script setup lang="ts">
import type { DropdownMenuItem } from "#ui/types"

const self = useSelf()
const route = useRoute()

const title = computed(() => (route.meta.title ? (route.meta.title as string) : undefined))

const userDropdownItems = computed<DropdownMenuItem[][]>(() => [
  [
    {
      type: "label",
      label: self.value?.login,
      avatar: {
        src: self.value?.avatar_url,
        alt: self.value?.login,
      },
    },
  ],
  // [
  //   {
  //     label: "Settings",
  //     icon: "lucide:settings",
  //     to: "/settings",
  //   },
  // ],
  [
    {
      label: "Log out",
      icon: "lucide:log-out",
      onSelect() {
        window.location.href = "/-/auth/logout"
      },
    },
  ],
])
</script>

<template>
  <UDashboardGroup class="bg-default">
    <UDashboardSidebar resizable collapsible class="bg-(--ui-bg-elevated)/25 mt-2"
      :ui="{ footer: 'lg:border-t lg:border-default' }">
      <template #default="{ collapsed }">
        <div class="flex flex-row gap-2">
          <UButton color="secondary" icon="lucide:arrow-left" variant="link" :square="collapsed" to="/">
            {{ collapsed ? "" : "Back" }}
          </UButton>
          <UButton v-if="!collapsed" color="primary" variant="subtle" block :square="collapsed">
            Admin Dashboard
          </UButton>
        </div>
        <UNavigationMenu :collapsed="collapsed" :items="adminSidebarOptions" orientation="vertical" />
      </template>

      <template #footer="{ collapsed }">
        <UDropdownMenu :items="userDropdownItems" :content="{ align: 'center', collisionPadding: 12 }"
          :ui="{ content: collapsed ? 'w-48' : 'w-(--reka-dropdown-menu-trigger-width)' }">
          <UButton v-bind="{
            avatar: {
              src: self?.avatar_url,
              alt: self?.login,
            },
            label: collapsed ? undefined : self?.name,
            trailingIcon: collapsed ? undefined : 'lucide:chevrons-up-down',
          }" color="neutral" variant="ghost" block :square="collapsed" class="data-[state=open]:bg-elevated mb-2" :ui="{
            trailingIcon: 'text-dimmed',
          }" />

          <template #chip-leading="{ item }">
            <span :style="{ '--chip': `var(--color-${(item as any).chip}-400)` }"
              class="ms-0.5 size-2 rounded-full bg-(--chip)" />
          </template>
        </UDropdownMenu>
      </template>
    </UDashboardSidebar>

    <UDashboardPanel>
      <template #header>
        <UDashboardNavbar :title="title">
          <template #leading>
            <UDashboardSidebarCollapse />
          </template>
        </UDashboardNavbar>
      </template>

      <template #body>
        <slot />
      </template>
    </UDashboardPanel>
  </UDashboardGroup>
</template>
