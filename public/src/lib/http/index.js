import axios from "axios"
import { ApiClient } from "@/lib/http/api"

const query = new ApiClient({
  BASE: "/api/query",
  CREDENTIALS: "same-origin",
  HEADERS: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
})

/**
 * @type {import("axios").AxiosInstance}
 */
const api = axios.create({
  baseURL: "/api",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  timeout: 5000,
})

export { query, api }
