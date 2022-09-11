import { NIcon } from "naive-ui"
import { RouterLink } from "vue-router"

import type { Component, VNode } from "vue"

function renderIcon(icon: Component | VNode | HTMLElement) {
  return () => h(NIcon, { style: "margin-top: -3px" }, { default: () => h(icon) })
}

function renderLink(target: Record<string, any>, title: string, icon: any) {
  return {
    label: () =>
      h(
        RouterLink,
        { to: target },
        {
          default: () => title,
        }
      ),
    key: target.name,
    icon: renderIcon(icon),
  }
}

export const menuOptions = [
  { to: { name: "index" }, name: "Home", alias: "main" },
  { to: { name: "posts" }, name: "Posts", alias: "posts" },
  { to: { name: "repos" }, name: "Repos", alias: "repos" },
]

export const adminSidebarOptions = [
  renderLink({ name: "admin" }, "Admin Home", IconMdiHome),
  renderLink({ name: "admin-posts" }, "Blog Posts", IconMdiBookOpenPageVariantOutline),
  { key: "divider-1", type: "divider" },
  ...menuOptions
    .filter((option) => option.name != "Sudo")
    .map((option) => renderLink(option.to, option.name, IconMdiLink)),
  { key: "divider-2", type: "divider" },
  {
    label: () => h("a", { href: "/-/auth/logout" }, { default: () => "Logout" }),
    key: "logout",
    icon: renderIcon(IconMdiLogout),
  },
  // { key: "divider-2", type: "divider" },
  // },
  // {
  //   label: "Pinball 1973",
  //   key: "pinball-1973",
  //   icon: renderIcon(IconMdiBookOpenPageVariantOutline),
  //   children: [
  //     {
  //       label: "Rat",
  //       key: "rat",
  //     },
  //   ],
  // },
]