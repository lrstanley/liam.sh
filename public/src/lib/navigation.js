import { RouterLink } from "vue-router"
import { NIcon } from "naive-ui"

function renderIcon(icon) {
  return () => h(NIcon, { style: "margin-top: -3px" }, { default: () => h(icon) })
}

function renderLink(target, title, icon) {
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
  renderLink({ name: "index" }, "Home", IconMdiHome),
  renderLink({ name: "about" }, "About", IconMdiHelpCircleOutline),
]

export const adminSidebarOptions = [
  renderLink({ name: "admin" }, "Admin Home", IconMdiHome),
  renderLink({ name: "admin-posts" }, "Blog Posts", IconMdiBookOpenPageVariantOutline),
  {
    key: "divider-1",
    type: "divider",
    props: {
      style: {
        marginLeft: "32px",
      },
    },
  },
  {
    label: "Hear the Wind Sing",
    key: "hear-the-wind-sing",
    icon: renderIcon(IconMdiBookOpenPageVariantOutline),
  },
  {
    label: "Pinball 1973",
    key: "pinball-1973",
    icon: renderIcon(IconMdiBookOpenPageVariantOutline),
    children: [
      {
        label: "Rat",
        key: "rat",
      },
    ],
  },
  {
    label: "A Wild Sheep Chase",
    key: "a-wild-sheep-chase",
    icon: renderIcon(IconMdiBookOpenPageVariantOutline),
  },
  {
    label: "Dance Dance Dance",
    key: "Dance Dance Dance",
    icon: renderIcon(IconMdiBookOpenPageVariantOutline),
    children: [
      {
        type: "group",
        label: "People",
        key: "people",
        children: [
          {
            label: "Narrator",
            key: "narrator",
            icon: renderIcon(IconMdiAccount),
          },
          {
            label: "Sheep Man",
            key: "sheep-man",
            icon: renderIcon(IconMdiAccount),
          },
        ],
      },
      {
        label: "Beverage",
        key: "beverage",
        icon: renderIcon(IconMdiAccount),
        children: [
          {
            label: "Whisky",
            key: "whisky",
          },
        ],
      },
      {
        label: "Food",
        key: "food",
        children: [
          {
            label: "Sandwich",
            key: "sandwich",
          },
        ],
      },
      {
        label: "The past increases. The future recedes.",
        key: "the-past-increases-the-future-recedes",
      },
    ],
  },
]
