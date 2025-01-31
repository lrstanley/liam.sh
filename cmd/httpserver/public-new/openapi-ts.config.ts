import { defineConfig } from "@hey-api/openapi-ts"

export default defineConfig({
  experimentalParser: true,
  input: {
    path: "../../../internal/database/ent/rest/openapi.json",
  },
  output: {
    path: "utils/http",
  },
  plugins: [
    "@hey-api/client-nuxt",
    "@hey-api/schemas",
    {
      name: "@hey-api/sdk",
      transformer: true,
    },
    {
      enums: "javascript",
      name: "@hey-api/typescript",
    },
    {
      name: "@hey-api/transformers",
      dates: true,
    },
  ],
})
