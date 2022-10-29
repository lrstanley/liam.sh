/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,html,js,ts,md}"],
  theme: {
    extend: {},
  },
  corePlugins: {
    preflight: false,
  },
  plugins: [],
  safelist: ["flex flex-row w-6 h-6 rounded-full ml-2 text-red-500"], // /admin/repo-needs-release
}
