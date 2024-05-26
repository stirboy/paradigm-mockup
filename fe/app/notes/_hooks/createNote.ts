import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import useSWRMutation from "swr/mutation";

async function sendRequest(url: string) {
  // const { toast } = useToast();¬
  return api.post(Routes.Notes, {
    title: "Untitled",
  });
}

export const useCreateNote = () => {
  const { trigger, isMutating } = useSWRMutation(Routes.Notes, sendRequest);

  return {
    trigger: trigger,
    isMutating: isMutating,
  };
};
