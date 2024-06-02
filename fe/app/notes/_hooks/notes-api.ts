import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import useSWR, { useSWRConfig } from "swr";
import useSWRMutation from "swr/mutation";
import { Note } from "@/app/notes/_api/models";
import useSWRImmutable from "swr/immutable";

export const useNotes = (parentId?: string) => {
  const { data, isLoading, error } = useSWRImmutable<Note[]>(
    parentId ? `${Routes.Notes}?parentId=${parentId}` : Routes.Notes,
  );
  return {
    notes: data,
    isLoading: isLoading,
    isError: error,
  };
};

export const useArchivedNotes = () => {
  const { data, isLoading, error } = useSWR<Note[]>(Routes.ArchivedNotes, {
    revalidateOnFocus: false,
    revalidateOnMount: true,
    revalidateIfStale: false,
  });
  return {
    notes: data,
    isLoading: isLoading,
    isError: error,
  };
};

export const useNote = (noteId: string) => {
  const { data, isLoading, error } = useSWR<Note>(`${Routes.Notes}/${noteId}`, {
    revalidateOnFocus: false,
    revalidateOnMount: true,
    revalidateIfStale: false,
  });
  return {
    note: data,
    isLoading: isLoading,
    isError: error,
  };
};

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
    const res = await api.post(Routes.Notes, {
      title: "Untitled",
      parentId: parentId,
    });

    await mutate(`${Routes.Notes}?parentId=${parentId}`);

    return res;
  };

  return { trigger };
};

export const useArchiveNotes = () => {
  const { mutate, cache } = useSWRConfig();

  const trigger = async (id: string, parentId?: string) => {
    await api.put(`${Routes.Notes}/${id}/archive`);

    await mutate(`${Routes.Notes}/${id}`);
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
    await mutate(`${Routes.Notes}/${id}`);
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

export type SourceType = "banner" | "trash";
export const useDeleteNote = (source?: SourceType) => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string) => {
    try {
      await api.delete(`${Routes.Notes}/${id}`);
    } catch (e) {
      throw e;
    }

    if (source === "trash") {
      await mutate(Routes.ArchivedNotes);
    }

    await mutate(`${Routes.Notes}/${id}`, undefined, {
      revalidate: false,
      populateCache: false,
      throwOnError: true,
    });
  };

  return { trigger };
};

export const useUpdateNote = () => {
  const { cache, mutate } = useSWRConfig();

  const trigger = async (id: string, title: string, parentId?: string) => {
    await api.put(`${Routes.Notes}/${id}`, {
      title: title,
    });
    mutate(`${Routes.Notes}/${id}`);
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

// async function updateRequest(
//   url: string,
//   { arg }: { arg: { id: string, title: string; icon: string } },
// ) {
//   return api.put(`${Routes.Notes}/${arg.id}`, {
//     title: arg.title,
//     icon: arg.icon
//   }).then((res) => res.data as []Note);
// }

export const useUpdateNoteIcon = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string, icon: string, parentId?: string) => {
    await api.put(`${Routes.Notes}/${id}`, {
      icon: icon,
    });
    mutate(`${Routes.Notes}/${id}`);
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

export const useRemoveNoteIcon = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string) => {
    await api.put(`${Routes.Notes}/${id}`, {
      icon: null,
    });
    mutate(`${Routes.Notes}/${id}`);
  };

  return { trigger };
};

export const useUpdateNoteContent = () => {
  const { mutate } = useSWRConfig();

  const trigger = async (id: string, content: string) => {
    await api.put(`${Routes.Notes}/${id}`, {
      content: content,
    });
    mutate(`${Routes.Notes}/${id}`);
  };

  return { trigger };
};
