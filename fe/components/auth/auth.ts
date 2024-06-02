import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import useSWRMutation from "swr/mutation";
import useSWR, { useSWRConfig } from "swr";
import { User } from "@/components/auth/user";

async function login(
  url: string,
  { arg }: { arg: { username: string; password: string } },
) {
  return api.post(url, {
    username: arg.username,
    password: arg.password,
  });
}

export const useUser = () => {
  const { data, isLoading, error } = useSWR<User>(Routes.User);

  return {
    user: data,
    isLoading: isLoading,
    error: error,
  };
};

export const useLogin = () => {
  const { trigger, isMutating } = useSWRMutation(Routes.Login, login);

  return {
    trigger: trigger,
    isMutating: isMutating,
  };
};

export const useLogout = () => {
  const { mutate } = useSWRConfig();

  const trigger = () => mutate(() => true, undefined, { revalidate: false });

  return { trigger };
};
