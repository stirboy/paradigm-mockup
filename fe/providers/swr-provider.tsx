"use client";

import { useToast } from "@/components/ui/use-toast";
import { getFetcher } from "@/lib/restapi";
import { AxiosError } from "axios";
import { useRouter } from "next/navigation";
import { SWRConfig } from "swr";

export const SWRProvider = ({ children }: { children: React.ReactNode }) => {
  const { toast } = useToast();
  const router = useRouter();

  function isSSR(): boolean {
    return typeof window === "undefined";
  }

  return (
    <SWRConfig
      value={{
        fetcher: getFetcher,
        onError: (err) => {
          const errorResponse = err?.response ?? ({} as AxiosError);
          const {
            status = 0,
            data = {
              message: errorResponse.statusText || "An error occurred",
            },
          } = errorResponse;

          if (isSSR()) {
            return Promise.reject(err);
          }

          switch (status) {
            case 401:
              toast({
                variant: "destructive",
                title: "Unauthorized",
                description: "You need to log in to access this page",
              });
              router.push("/login");
              break;
            case 403:
              toast({
                variant: "destructive",
                title: "Forbidden",
                description: "You do not have permission to access this page",
              });
              router.push("/login");
              break;
            case 404:
              toast({
                variant: "destructive",
                title: "Not found",
                description: "The requested resource was not found",
              });
              break;
            default:
              return Promise.reject({ ...data, httpStatus: status });
          }
        },
      }}
    >
      {children}
    </SWRConfig>
  );
};
