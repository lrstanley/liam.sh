import { ApiClient } from "@/lib/http/api"
import axios from "axios"

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
})

export { query, api }
