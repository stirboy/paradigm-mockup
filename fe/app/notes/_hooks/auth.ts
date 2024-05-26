import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import useSWRMutation from "swr/mutation";

async function login(
  url: string,
  { arg }: { arg: { username: string; password: string } },
) {
  return api.post(url, {
    username: arg.username,
    password: arg.password,
  });
}

export const useLogin = () => {
  const { trigger, isMutating } = useSWRMutation(Routes.Login, login);

  return {
    trigger: trigger,
    isMutating: isMutating,
  };
};
