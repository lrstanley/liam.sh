<script setup lang="ts">
import type { DataTableColumns } from "naive-ui"
import { listOutdatedGithubReleases } from "@/lib/http/services.gen"
import type { OutdatedRepositoryRelease } from "@/lib/http/types.gen"

definePage({
  meta: {
    layout: "admin",
  },
})

const {
  data: results,
  isFetching,
  error,
  suspense,
} = useQuery({
  queryKey: ["releases", "outdated"],
  queryFn: () => unwrapErrors(listOutdatedGithubReleases()),
})
await suspense()

watch(error, () => {
  if (error.value) throw error.value
})

const now = new Date()

const columns: DataTableColumns<OutdatedRepositoryRelease> = [
  {
    title: "Repository",
    key: "repository",
    render: (data: OutdatedRepositoryRelease) =>
      h(
        "a",
        { class: "flex flex-row", href: data.repository.html_url + "/releases", target: "_blank" },
        [
          h("img", { src: data.repository.owner.avatar_url, class: "w-6 h-6 rounded-full" }),
          h("span", { class: "ml-2" }, data.repository.full_name),
        ]
      ),
  },
  {
    title: "Last Release",
    key: "lastRelease",
    render: (data: OutdatedRepositoryRelease) =>
      h("a", { class: "flex flex-row", href: data.release.html_url, target: "_blank" }, [
        h("span", {}, data.release.name || data.release.tag_name),
      ]),
  },
  {
    title: "Last Release Date",
    key: "lastRelease",
    render: (data: OutdatedRepositoryRelease) => h("span", {}, data.release.created_at),
  },
  {
    title: "Last Release Relative",
    key: "lastRelease",
    render: (data: OutdatedRepositoryRelease) => {
      const diffMs = now.getTime() - new Date(data.release.created_at).getTime()
      const diffDays = Math.round(diffMs / 86400000)

      return h(
        "span",
        { class: diffDays > 90 ? "text-red-500" : "" },
        useTimeAgo(data.release.created_at).value
      )
    },
  },
]
</script>

<template>
  <div class="p-4 sm:container sm:mx-auto">
    <n-page-header
      subtitle="Find repos that may need a new release published"
      class="hidden mt-4 mb-8 lg:block"
    >
      <template #avatar>
        <n-icon :size="40"><i-mdi-history /></n-icon>
      </template>
      <template #title>
        <a href="#" class="no-underline capitalize" style="color: inherit">Repository Releases</a>
      </template>
    </n-page-header>

    <n-data-table :loading="isFetching" :columns="columns" :data="results" />
  </div>
</template>
