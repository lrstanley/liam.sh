/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { NIcon } from "naive-ui"
import { RouterLink } from "vue-router"

import type { RouteNamedMap } from "vue-router/auto-routes"
import type { RouteLocationNormalized } from "vue-router"
import type { Component, VNode } from "vue"
import { Icon } from "#components"

type Route = keyof RouteNamedMap | RouteLocationNormalized

interface Link {
  to: Route
  name: string
  alias?: string
}

function renderIcon(icon: Component | VNode | HTMLElement) {
  return () => h(NIcon, { style: "margin-top: -3px" }, { default: () => h(icon) })
}

function renderLink(target: Route, title: string, icon: any) {
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
    icon: renderIcon(icon),
  }
}

export const menuOptions: Link[] = [
  { to: "index", name: "Home", alias: "main" },
  // { to: "/posts", name: "Posts", alias: "posts" },
  // { to: "/repos", name: "Repos", alias: "repos" },
]

export const branchMenuOptions: Link[] = [
  { to: "index", name: "master" },
  // { to: "/util/gists", name: "feature/list-gists" },
]

// TODO: fix type.
export const adminSidebarOptions: any = [
  // renderLink("/admin/", "Admin Home", h(Icon, { name: "mdi:home" })),
  // renderLink("/admin/posts", "Blog Posts", h(Icon, { name: "mdi:book-open-page-variant-outline" })),
  // renderLink("/admin/banner-builder", "Banner Builder", h(Icon, { name: "mdi:image-edit-outline" })),
  // renderLink("/admin/repo-needs-release", "Repo Releases", h(Icon, { name: "mdi:history" })),
  { key: "divider-1", type: "divider" },
  // ...menuOptions.map((option) => renderLink(option.to, option.name, h(Icon, { name: "mdi:link" }))),
  { key: "divider-2", type: "divider" },
  {
    label: () => h("a", { href: "/-/auth/logout" }, { default: () => "Logout" }),
    key: "logout",
    icon: renderIcon(h(Icon, { name: "mdi:logout" })),
  },
]
