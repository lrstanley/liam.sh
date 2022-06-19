import { createDiscreteApi } from "naive-ui"

export const { message, dialog, loadingBar } = createDiscreteApi(
  ["message", "dialog", "loadingBar"] // "notification"
)
