import axios from "axios";

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

const getFetcher = (url: string) => api.get(url).then((res) => res.data);

async function fetchByParentId(url: string, parentId: string) {
  return await api.get(`${url}?parentId=${parentId}`).then((res) => res.data);
}

export { api, fetchByParentId, getFetcher };
