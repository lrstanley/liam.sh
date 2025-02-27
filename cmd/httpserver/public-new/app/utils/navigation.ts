/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import type { RouteNamedMap } from "vue-router/auto-routes"
import type { NavigationMenuItem } from "#ui/types"

type Route = keyof RouteNamedMap | RouteNamedMap[keyof RouteNamedMap]["path"]

export type Link = NavigationMenuItem & {
  to?: Route
  functionAlias?: string
}

export const menuOptions: Link[] = [
  { to: "/", label: "Home", functionAlias: "main", icon: "i-lucide-home" },
  // { to: "/posts", label: "Posts" },
  // { to: "/repos", label: "Repos" },
]

export const branchMenuOptions: Link[] = [
  { to: "/", label: "master" },
  { to: "/util/gists", label: "feature/list-gists" },
]

// TODO: use badges here -- show post count, and other things like that?
export const adminSidebarOptions: Link[] = [
  {
    to: "/admin",
    icon: "i-lucide-lock",
    label: "Home",
  },
  // {
  //   to: "/admin/posts",
  //   icon: "i-lucide-book-open-page-variant-outline",
  //   label: "Blog Posts",
  // },
  // {
  //   to: "/admin/banner-builder",
  //   icon: "i-lucide-image-edit-outline",
  //   label: "Banner Builder",
  // },
  // {
  //   to: "/admin/repo-needs-release",
  //   icon: "i-lucide-history",
  //   label: "Repo Releases",
  // },
  // {
  //   icon: "i-lucide-log-out",
  //   label: "Logout",
  //   onSelect() {
  //     window.location.href = "/-/auth/logout"
  //   },
  // },
  {
    label: "User Facing Pages",
    icon: "i-lucide-user",
    defaultOpen: true,
    children: menuOptions,
  },
]
