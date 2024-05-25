import axios, {AxiosError} from "axios";
import {toast} from "@/components/ui/use-toast";

const api = axios.create({
  baseURL: "http://localhost:8080",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

function isSSR(): boolean {
  return typeof window === "undefined";
}

api.interceptors.response.use(
  (response) => response,
  (error) => {
    const errorResponse = error?.response ?? ({} as AxiosError);
    const {
      status = 0,
      data = {
        message: errorResponse.statusText || "An error occurred",
      },
    } = errorResponse;

    if (isSSR()) {
      return Promise.reject(error)
    }

    switch (status) {
      case 401:
        window.location.assign("/login");
        localStorage.setItem("authToast", "true")
        break;
      default:
        return Promise.reject({...data, httpStatus: status})
    }
  });

export default api;
