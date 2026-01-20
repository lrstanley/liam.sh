/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { RouteNamedMap } from "vue-router/auto-routes"
import type { NavigationMenuItem } from "#ui/types"

type Route = RouteNamedMap[keyof RouteNamedMap]["path"]

export type Link = NavigationMenuItem & {
  to?: Route
  functionAlias?: string
}

export const menuOptions: Link[] = [
  { to: "/", label: "Home", functionAlias: "main", icon: "lucide:home" },
  { to: "/posts", label: "Posts", functionAlias: "posts", icon: "lucide:notebook-pen" },
  { to: "/repos", label: "Repos", functionAlias: "repos", icon: "lucide:git-pull-request" },
]

export const branchMenuOptions: Link[] = [
  { to: "/", label: "master" },
  { to: "/util/gists", label: "feature/list-gists" },
]

// TODO: use badges here -- show post count, and other things like that?
export const adminSidebarOptions: Link[] = [
  {
    label: "Home",
    to: "/admin",
    icon: "lucide:lock",
  },
  {
    label: "Manage Posts",
    to: "/admin/posts",
    icon: "lucide:book-open-text",
  },
  {
    label: "Manage Labels",
    to: "/admin/labels",
    icon: "lucide:tag",
  },
  {
    label: "Banner Builder",
    to: "/admin/banner-builder",
    icon: "mdi:image-edit-outline",
  },
  {
    label: "Repository Releases",
    to: "/admin/repository-releases",
    icon: "mdi:history",
  },
]
