<script setup lang="ts">
import type { TableColumn } from "#ui/types"

useHead({ title: "Repository Releases" })
definePageMeta({
  layout: "admin",
})

const { data, error, status } = await listOutdatedGithubReleases({
  composable: "useFetch",
})
if (error.value) throw error.value

const now = new Date()

const columns: TableColumn<OutdatedRepositoryRelease>[] = [
  {
    header: "Repository",
    accessorKey: "repository",
    cell: ({ row }) =>
      h(
        "a",
        {
          class: "flex flex-row",
          href: row.original.repository.html_url + "/releases",
          target: "_blank",
        },
        [
          h("img", { src: row.original.repository.owner.avatar_url, class: "w-6 h-6 rounded-full" }),
          h("span", { class: "ml-2" }, row.original.repository.full_name),
        ]
      ),
  },
  {
    header: "Last Release",
    accessorKey: "lastRelease",
    cell: ({ row }) =>
      h("a", { class: "flex flex-row", href: row.original.release.html_url, target: "_blank" }, [
        h("span", {}, row.original.release.name || row.original.release.tag_name),
      ]),
  },
  {
    header: "Last Release Date",
    accessorKey: "lastRelease",
    cell: ({ row }) => h("span", {}, row.original.release.created_at.toLocaleString()),
  },
  {
    header: "Last Release Relative",
    accessorKey: "lastRelease",
    cell: ({ row }) => {
      const diffMs = now.getTime() - new Date(row.original.release.created_at).getTime()
      const diffDays = Math.round(diffMs / 86400000)

      return h(
        "span",
        { class: diffDays > 90 ? "text-red-500" : "" },
        useTimeAgo(row.original.release.created_at).value
      )
    },
  },
]
</script>

<template>
  <UDashboardPanel id="customers">
    <template #header>
      <UDashboardNavbar title="Repository Releases">
        <template #leading>
          <UDashboardSidebarCollapse />
        </template>
      </UDashboardNavbar>
    </template>

    <template #body>
      <UCard
        :ui="{
          body: '!px-2 !py-0',
        }"
      >
        <UTable
          v-if="data"
          :data="data"
          :columns="columns"
          :loading="status === 'pending'"
          class="shrink-0"
        />
      </UCard>
    </template>
  </UDashboardPanel>
</template>
