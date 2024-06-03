"use client";

import { useToast } from "@/components/ui/use-toast";
import { getFetcher } from "@/lib/restapi";
import { AxiosError, HttpStatusCode } from "axios";
import { useRouter } from "next/navigation";
import { SWRConfig } from "swr";
import { useLogout } from "@/components/auth/auth";

export const SWRProvider = ({ children }: { children: React.ReactNode }) => {
  const { toast } = useToast();
  const router = useRouter();
  const { trigger: logout } = useLogout();

  function isSSR(): boolean {
    return typeof window === "undefined";
  }

  return (
    <SWRConfig
      value={{
        fetcher: getFetcher,
        revalidateIfStale: false,
        onError: (err) => {
          const errorResponse = err?.response ?? ({} as AxiosError);
          const {
            status = 0,
            data = {
              code: errorResponse.code || "UNKNOWN_ERROR",
              message: errorResponse.statusText || "An error occurred",
            },
          } = errorResponse;

          if (isSSR()) {
            return Promise.reject(err);
          }

          switch (status) {
            case HttpStatusCode.Unauthorized:
              toast({
                variant: "destructive",
                title: "Unauthorized",
                description: "You need to log in to access this page",
              });
              router.push("/login");
              logout();
              break;
            case HttpStatusCode.Forbidden:
              toast({
                variant: "destructive",
                title: "Forbidden",
                description: "You do not have permission to access this page",
              });
              router.push("/login");
              break;
            case HttpStatusCode.NotFound:
              toast({
                variant: "destructive",
                title: "Not found",
                description: "The requested resource was not found",
              });
              break;
            case HttpStatusCode.InternalServerError:
              toast({
                variant: "destructive",
                title: "Internal server error",
                description: "An error occurred while processing your request",
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
