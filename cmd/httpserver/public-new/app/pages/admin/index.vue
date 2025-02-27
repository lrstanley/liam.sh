<script setup lang="ts">
useHead({ title: "Administration" })
definePageMeta({
  layout: "admin",
})
</script>

<template>
  <UDashboardPanel id="customers">
    <template #header>
      <UDashboardNavbar title="Home">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>

        <!-- <template #right>
          <CustomersAddModal />
        </template> -->
      </UDashboardNavbar>
    </template>

    <template #body>FOO</template>
  </UDashboardPanel>
</template>
