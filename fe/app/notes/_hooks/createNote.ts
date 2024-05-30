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

    await mutate(`${Routes.Notes}?parentId=${parentId}`);
  };

  return { trigger };
};

export const useArchiveNotes = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string, parentId?: string) => {
    await api.put(`${Routes.Notes}/${id}/archive`);

    if (parentId) {
      // refresh the parent notes list staring from the parentId node
      await mutate(`${Routes.Notes}?parentId=${parentId}`);
    } else {
      // refresh the root notes list
      await mutate(`${Routes.Notes}`);
    }
  };

  return { trigger };
};

export const useRestoreNotes = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string, parentId?: string) => {
    await api.put(`${Routes.Notes}/${id}/restore`);

    await mutate(Routes.ArchivedNotes);
    if (parentId) {
      // refresh the parent notes list staring from the parentId node
      await mutate(`${Routes.Notes}?parentId=${parentId}`);
    } else {
      // refresh the root notes list
      await mutate(`${Routes.Notes}`);
    }
  };

  return { trigger };
};
