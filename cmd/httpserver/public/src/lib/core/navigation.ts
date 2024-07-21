/**
 * Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
 * this source code is governed by the MIT license that can be found in
 * the LICENSE file.
 */

import { NIcon } from "naive-ui"
import { RouterLink } from "vue-router"

import type { RouteNamedMap } from "vue-router/auto-routes"
import type { RouteLocationNormalized } from "vue-router/auto"
import type { Component, VNode } from "vue"

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
  { to: "/", name: "Home", alias: "main" },
  { to: "/posts", name: "Posts", alias: "posts" },
  { to: "/repos", name: "Repos", alias: "repos" },
]

export const branchMenuOptions: Link[] = [
  { to: "/", name: "master" },
  { to: "/util/gists", name: "feature/list-gists" },
]

export const adminSidebarOptions = [
  renderLink("/admin/", "Admin Home", IMdiHome),
  renderLink("/admin/posts", "Blog Posts", IMdiBookOpenPageVariantOutline),
  renderLink("/admin/banner-builder", "Banner Builder", IMdiImageEditOutline),
  renderLink("/admin/repo-needs-release", "Repo Releases", IMdiHistory),
  { key: "divider-1", type: "divider" },
  ...menuOptions.map((option) => renderLink(option.to, option.name, IMdiLink)),
  { key: "divider-2", type: "divider" },
  {
    label: () => h("a", { href: "/-/auth/logout" }, { default: () => "Logout" }),
    key: "logout",
    icon: renderIcon(IMdiLogout),
  },
]
