<script setup lang="ts">
import type { GithubRelease, GithubRepository, GithubUser } from "@/lib/api"
import { useLatestRepoReleasesQuery } from "@/lib/api"
import { useTimeAgo } from "@vueuse/core"
import type { DataTableColumns } from "naive-ui"

definePage({
  meta: {
    layout: "admin",
  },
})

const { data, error } = await useLatestRepoReleasesQuery()

watch(error, () => {
  if (error.value) throw error.value
})

interface RowData {
  repo: Pick<GithubRepository, "fullName" | "htmlURL">
  owner: Pick<GithubUser, "avatarURL">
  release: Pick<GithubRelease, "createdAt">
}

const now = new Date()

const columns: DataTableColumns<RowData> = [
  {
    title: "Repository",
    key: "repository",
    render: ({ repo, owner }) =>
      h("a", { class: "flex flex-row", href: repo.htmlURL, target: "_blank" }, [
        h("img", { src: owner.avatarURL, class: "w-6 h-6 rounded-full" }),
        h("span", { class: "ml-2" }, repo.fullName),
      ]),
  },
  {
    title: "Last Release Date",
    key: "lastRelease",
    render: ({ release }) => h("span", {}, release.createdAt),
  },
  {
    title: "Last Release Relative",
    key: "lastRelease",
    render: ({ release }) => {
      const diffMs = now.getTime() - new Date(release.createdAt).getTime()
      const diffDays = Math.round(diffMs / 86400000)

      return h(
        "span",
        { class: diffDays > 90 ? "text-red-500" : "" },
        useTimeAgo(release.createdAt).value
      )
    },
  },
]

const results = computed<RowData[]>(() => {
  return (
    data.value?.githubRepositories.edges
      ?.map(({ node }) => {
        const release = node.releases.edges?.[0]?.node
        return {
          repo: node,
          owner: node.owner,
          release: release,
        }
      })
      .sort((a, b) => {
        return new Date(a.release.createdAt).getTime() - new Date(b.release.createdAt).getTime()
      }) ?? []
  )
})
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

    <n-data-table :columns="columns" :data="results" />
  </div>
</template>
