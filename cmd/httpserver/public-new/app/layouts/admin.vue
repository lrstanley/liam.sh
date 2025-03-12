<script setup lang="ts">
console.log(import.meta.env.MODE)
</script>

<template>
  <UDashboardGroup>
    <UDashboardSidebar
      resizable
      collapsible
      class="bg-(--ui-bg-elevated)/25"
      :ui="{ footer: 'lg:border-t lg:border-(--ui-border)' }"
    >
      <template #default="{ collapsed }">
        <Ubutton color="neutral" variant="ghost" block :square="collapsed">Test</Ubutton>
        <UNavigationMenu :collapsed="collapsed" :items="adminSidebarOptions" orientation="vertical" />

        <!-- TODO: show user on bottom, similar to dashboard template -->
      </template>
    </UDashboardSidebar>

    <slot />
  </UDashboardGroup>
</template>
