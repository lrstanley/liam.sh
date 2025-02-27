<script setup lang="ts"></script>

<template>
  <UDashboardGroup>
    <UDashboardSidebar
      resizable
      collapsible
      class="bg-(--ui-bg-elevated)/25"
      :ui="{ footer: 'lg:border-t lg:border-(--ui-border)' }"
    >
      <template #default="{ collapsed }">
        <UNavigationMenu :collapsed="collapsed" :items="adminSidebarOptions" orientation="vertical" />

        <!-- TODO: show user on bottom, similar to dashboard template -->
      </template>
    </UDashboardSidebar>

    <slot />
  </UDashboardGroup>
</template>
