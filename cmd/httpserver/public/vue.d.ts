declare module "*.vue" {
  import type { DefineComponent } from "vue"
  const component: DefineComponent<{}, {}, any> // eslint-disable-line @typescript-eslint/ban-types
  export default component
}
