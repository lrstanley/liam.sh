/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { RouterLink } from "vue-router"
import type { RouteNamedMap } from "vue-router/auto-routes"
import { UIcon } from "#components"

type Route = keyof RouteNamedMap | RouteNamedMap[keyof RouteNamedMap]["path"]

interface Link {
  to: Route
  name: string
  alias?: string
}

function renderLink(target: Route, title: string, icon: string) {
  return {
    label: () =>
      h(
        RouterLink,
        { to: target },
        {
          default: () => title,
        }
      ),
    key: title,
    icon: icon ? h(UIcon, { name: icon }) : null,
  }
}

export const menuOptions: Link[] = [
  { to: "/", name: "Home", alias: "main" },
  // { to: "/posts", name: "Posts", alias: "posts" },
  // { to: "/repos", name: "Repos", alias: "repos" },
]

export const branchMenuOptions: Link[] = [
  { to: "index", name: "master" },
  // { to: "/util/gists", name: "feature/list-gists" },
]

// TODO: fix type.
export const adminSidebarOptions: any = [
  // renderLink("/admin/", "Admin Home", h(UIcon, { name: "mdi:home" })),
  // renderLink("/admin/posts", "Blog Posts", h(UIcon, { name: "mdi:book-open-page-variant-outline" })),
  // renderLink("/admin/banner-builder", "Banner Builder", h(UIcon, { name: "mdi:image-edit-outline" })),
  // renderLink("/admin/repo-needs-release", "Repo Releases", h(UIcon, { name: "mdi:history" })),
  { key: "divider-1", type: "divider" },
  // ...menuOptions.map((option) => renderLink(option.to, option.name, h(UIcon, { name: "mdi:link" }))),
  { key: "divider-2", type: "divider" },
  {
    label: () => h("a", { href: "/-/auth/logout" }, { default: () => "Logout" }),
    key: "logout",
    icon: h(UIcon, { name: "heroicons:arrow-right-start-on-rectangle-solid" }),
  },
]
