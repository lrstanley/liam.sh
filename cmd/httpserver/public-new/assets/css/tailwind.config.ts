import type { Config } from "tailwindcss"
import defaultTheme from "tailwindcss/defaultTheme"

export default <Partial<Config>>{
  content: [
    "./pages/**/*.{vue,js,ts,jsx,tsx}",
    "./layouts/**/*.{vue,js,ts,jsx,tsx}",
    "./components/**/*.{vue,js,ts,jsx,tsx}",
    "*.{vue,js,ts,jsx,tsx}",
  ],
  safelist: ["flex flex-row w-6 h-6 rounded-full ml-2 text-red-500"], // /admin/repo-needs-release
  theme: {
    extend: {
      fontFamily: {
        sans: ["DM Sans", ...defaultTheme.fontFamily.sans],
      },
    },
  },
}
