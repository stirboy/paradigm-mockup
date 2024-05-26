import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

const getFetcher = (url: string) => api.get(url).then((res) => res.data);

export { api, getFetcher };
