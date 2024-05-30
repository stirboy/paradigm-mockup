"use client";

import React, { useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useRestoreNotes } from "@/app/notes/_hooks/createNote";
import useSWR from "swr";
import { Note } from "@/app/notes/_api/models";
import { Routes } from "@/lib/constants/routes";
import { useToast } from "@/components/ui/use-toast";
import { Spinner } from "@/components/spinner";
import { Search, Trash, Undo } from "lucide-react";
import { Input } from "@/components/ui/input";

const TrashBox = () => {
  const router = useRouter();
  const params = useParams();
  const { toast } = useToast();
  const { data: notes, isLoading } = useSWR<Note[]>(Routes.ArchivedNotes);
  const { trigger: restoreNote } = useRestoreNotes();

  const [search, setSearch] = useState("");

  const filteredNotes = notes?.filter((note) => {
    return note.title.toLowerCase().includes(search.toLowerCase());
  });

  const onClick = (noteId: string) => {
    router.push(`/notes/${noteId}`);
  };

  const onRestore = async (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    noteId: string,
    parentId?: string,
  ) => {
    event.stopPropagation();
    restoreNote(noteId, parentId).then(() => {
      toast({
        title: "Notes are restored",
      });
    });
  };

  const onRemove = async (noteId: string, parentId?: string) => {
    // delete note

    if (params.id === noteId) {
      router.push(`/notes`);
    }
  };

  if (isLoading) {
    return (
      <div className={"h-full flex items-center justify-center p-4"}>
        <Spinner size={"lg"} />
      </div>
    );
  }

  return (
    <div className={"text-sm"}>
      <div className={"flex items-center gap-x-1 p-2"}>
        <Search className={"w-4 h-4"} />
        <Input
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className={"h-7 px-2 focus-visible:ring-transparent bg-secondary"}
          placeholder={"Filter by note title..."}
        />
      </div>
      <div className={"mt-2 px-1 pb-1"}>
        <p
          className={
            "hidden last:block text-sx text-center text-muted-foreground pb-2"
          }
        >
          No notes found.
        </p>
        {filteredNotes?.map((note) => (
          <div
            key={note.id}
            role={"button"}
            onClick={() => onClick(note.id)}
            className={
              "text-sm rounded-sm w-full hover:bg-primary/5 flex items-center text-primary justify-between"
            }
          >
            <span className={"truncate pl-2"}>{note.title}</span>
            <div className={"flex items-center"}>
              <div
                role="button"
                onClick={(e) => onRestore(e, note.id, note.parentId)}
                className="rounded-sm p-2 hover:bg-neutral-200 dark:hover:bg-neutral-600"
              >
                <Undo className={"h-4 w-4 text-muted-foreground"} />
              </div>
              <div
                role={"button"}
                className={"rounded-sm p-2 hover:bg-neutral-200"}
              >
                <Trash className={"h-4 w-4 text-muted-foreground"} />
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default TrashBox;
