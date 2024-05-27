import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import { useSWRConfig } from "swr";
import useSWRMutation from "swr/mutation";

async function sendRequest(url: string) {
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

export const useCreateNoteWithParent = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (parentId: string) => {
    await api.post(Routes.Notes, {
      title: "Untitled",
      parentId: parentId,
    });

    mutate(`${Routes.Notes}?parentId=${parentId}`);
  };

  return { trigger };
};
