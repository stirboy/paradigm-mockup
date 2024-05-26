import { QueryCache, QueryClient } from "@tanstack/react-query";
import { AxiosError } from "axios";

const queryClient = new QueryClient({
  queryCache: new QueryCache({
    // @ts-ignore err: unknown -> err: AxiosError
    onError: (error: AxiosError) => {
      window.location.assign("/");
      console.error(error);
    },
  }),
});

export { queryClient };
