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
  // {
  //   to: "/admin/posts",
  //   icon: "lucide:book-open-page-variant-outline",
  //   label: "Blog Posts",
  // },
  // {
  //   to: "/admin/banner-builder",
  //   icon: "lucide:image-edit-outline",
  //   label: "Banner Builder",
  // },
  {
    label: "Repository Releases",
    to: "/admin/repository-releases",
    icon: "mdi:history",
  },
  // {
  //   icon: "lucide:log-out",
  //   label: "Logout",
  //   onSelect() {
  //     window.location.href = "/-/auth/logout"
  //   },
  // },
]
