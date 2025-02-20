"use client";

import React from "react";
import { Note } from "../_api/models";
import { useParams, useRouter } from "next/navigation";
import Item from "./Item";
import { cn } from "@/lib/utils";
import { FileIcon } from "lucide-react";
import { useNotes } from "@/app/notes/_hooks/notes-api";

interface NotesListProps {
  parentId?: string;
  level?: number;
  data?: Note[];
}

const NotesList = ({ parentId, level = 0 }: NotesListProps) => {
  const params = useParams();
  const router = useRouter();

  const { notes, isLoading } = useNotes(parentId);

  const [expanded, setExpanded] = React.useState<Record<string, boolean>>({});

  const handleExpand = (id: string) => {
    setExpanded((prev) => ({
      ...prev,
      [id]: !prev[id],
    }));
  };

  const onRedirect = (id: string) => {
    router.push(`/notes/${id}`);
  };

  if (isLoading) {
    return (
      <>
        <Item.ItemSkeleton level={level} />
        {level === 0 && (
          <>
            <Item.ItemSkeleton level={level} />
            <Item.ItemSkeleton level={level} />
          </>
        )}
      </>
    );
  }

  if (notes === undefined) {
    return (
      <>
        <Item.ItemSkeleton level={level} />
        {level === 0 && (
          <>
            <Item.ItemSkeleton level={level} />
            <Item.ItemSkeleton level={level} />
          </>
        )}
      </>
    );
  }

  return (
    <>
      <p
        style={{
          paddingLeft: level ? `${level * 12 + 25}px` : "12px",
        }}
        className={cn(
          "hidden text-sm font-medium text-muted-foreground/80",
          expanded && "last:block",
          level === 0 && "hidden",
        )}
      >
        No notes inside
      </p>
      {notes?.map((note) => (
        <div key={note.id}>
          <Item
            id={note.id}
            parentId={note.parentId}
            onClick={() => onRedirect(note.id)}
            label={note.title}
            icon={FileIcon}
            documentIcon={note.icon}
            active={params.id === note.id}
            level={level}
            onExpand={() => handleExpand(note.id)}
            expanded={expanded[note.id]}
          />
          {expanded[note.id] && (
            <NotesList parentId={note.id} level={level + 1} />
          )}
        </div>
      ))}
    </>
  );
};

export default NotesList;
