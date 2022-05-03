import { RouterLink } from "vue-router"
import { NIcon } from "naive-ui"
import { FingerPrint } from "@vicons/ionicons5"

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
        icon: renderIcon(FingerPrint),
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
        icon: renderIcon(FingerPrint),
    },
]
