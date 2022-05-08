import { RouterLink } from "vue-router"
import { NIcon } from "naive-ui"

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) })
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
  renderLink({ name: "index" }, "Home", Home),
  renderLink({ name: "about" }, "About", HelpCircleOutline),
]

export const adminSidebarOptions = [
  renderLink({ name: "admin" }, "Admin Home", Home),
  renderLink({ name: "admin-posts" }, "Blog Posts", Book),
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
    icon: renderIcon(BookOutline),
  },
  {
    label: "Pinball 1973",
    key: "pinball-1973",
    icon: renderIcon(BookOutline),
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
    icon: renderIcon(BookOutline),
  },
  {
    label: "Dance Dance Dance",
    key: "Dance Dance Dance",
    icon: renderIcon(BookOutline),
    children: [
      {
        type: "group",
        label: "People",
        key: "people",
        children: [
          {
            label: "Narrator",
            key: "narrator",
            icon: renderIcon(PersonOutline),
          },
          {
            label: "Sheep Man",
            key: "sheep-man",
            icon: renderIcon(PersonOutline),
          },
        ],
      },
      {
        label: "Beverage",
        key: "beverage",
        icon: renderIcon(WineOutline),
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
