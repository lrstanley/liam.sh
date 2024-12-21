import { defineConfig } from "@hey-api/openapi-ts"

export default defineConfig({
  client: "@hey-api/client-fetch",
  input: "../../../internal/database/ent/rest/openapi.json",
  output: {
    path: "utils/http",
  },
  types: {
    enums: "javascript",
    // dates: "types+transform",
  },
  plugins: ["@tanstack/vue-query"],
})
