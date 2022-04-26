import { RouterLink } from "vue-router"
import { NIcon } from "naive-ui"
import { Book20Filled } from "@vicons/fluent"

function renderIcon(icon) {
    return () => h(NIcon, null, { default: () => h(icon) })
}

export const menuOptions = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: "/",
                },
                {
                    default: () => "Home",
                }
            ),
        key: "home",
        icon: renderIcon(Book20Filled),
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: "/about",
                },
                {
                    default: () => "About",
                }
            ),
        key: "about",
        icon: renderIcon(Book20Filled),
    },
]
