"use client";

import { useEffect, useState } from "react";
import { File } from "lucide-react";
import { useRouter } from "next/navigation";

import {
  CommandDialog,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import { useSearch } from "@/app/notes/_hooks/use-search";
import { Skeleton } from "@/components/ui/skeleton";
import { useNotes } from "@/app/notes/_hooks/notes-api";

export const SearchCommand = () => {
  const router = useRouter();
  const { notes, isLoading } = useNotes();
  const [isMounted, setIsMounted] = useState(false);

  const toggle = useSearch((store) => store.toggle);
  const isOpen = useSearch((store) => store.isOpen);
  const onClose = useSearch((store) => store.onClose);

  useEffect(() => {
    setIsMounted(true);
  }, []);

  useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "k" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        toggle();
      }
    };

    document.addEventListener("keydown", down);
    return () => document.removeEventListener("keydown", down);
  }, [toggle]);

  const onSelect = (id: string) => {
    router.push(`/notes/${id}`);
    onClose();
  };

  if (!isMounted) {
    return null;
  }

  if (isLoading) {
    return (
      <>
        <Skeleton className="h-12" />
      </>
    );
  }

  return (
    <CommandDialog open={isOpen} onOpenChange={onClose}>
      <CommandInput />
      <CommandList>
        <CommandEmpty>No results found.</CommandEmpty>
        <CommandGroup heading="Documents">
          {notes?.map((note) => (
            <CommandItem
              key={note.id}
              value={`${note.id}-${note.title}`}
              title={note.title}
              onSelect={() => onSelect(note.id)}
            >
              {note.icon ? (
                <p className="mr-2 text-[18px]">{note.icon}</p>
              ) : (
                <File className="mr-2 h-4 w-4" />
              )}
              <span>{note.title}</span>
            </CommandItem>
          ))}
        </CommandGroup>
      </CommandList>
    </CommandDialog>
  );
};
