import { defineConfig } from "@hey-api/openapi-ts"

export default defineConfig({
  client: "@hey-api/client-fetch",
  input: "../../../internal/database/ent/rest/openapi.json",
  output: {
    path: "src/lib/http",
  },
  types: {
    enums: "javascript",
    // dates: "types+transform",
  },
})
