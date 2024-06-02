"use client";

import React from "react";
import { Routes } from "@/lib/constants/routes";
import Toolbar from "@/components/toolbar";
import dynamic from "next/dynamic";
import { api } from "@/lib/restapi";
import { useNote } from "@/app/notes/_hooks/notes-api";
import { Skeleton } from "@/components/ui/skeleton";

const Editor = dynamic(() => import("../_components/editor"), { ssr: false });

interface NoteIdPageProps {
  params: {
    noteId: string;
  };
}

const NoteIdPage = ({ params }: NoteIdPageProps) => {
  const { note, isLoading } = useNote(params.noteId);

  if (isLoading || note === undefined) {
    return (
      <div>
        <Skeleton className="w-full h-[12vh]" />
        <div className="md:max-w-3xl lg:max-w-4xl mx-auto mt-10">
          <div className="space-y-4 pl-8 pt-4">
            <Skeleton className="h-14 w-[50%]" />
            <Skeleton className="h-4 w-[80%]" />
            <Skeleton className="h-4 w-[40%]" />
            <Skeleton className="h-4 w-[60%]" />
          </div>
        </div>
      </div>
    );
  }

  const onContentChange = async (doc: string) => {
    await api.put(`${Routes.Notes}/${params.noteId}`, { content: doc });
  };

  return (
    <div className={"pb-40"}>
      <div className={"h-[12vh]"} />
      <div className={"md:max-w-3xl lg:max-w-4xl mx-auto"}>
        <Toolbar initialData={note!} />
        <Editor initialContent={note!.content} onChange={onContentChange} />
      </div>
    </div>
  );
};

export default NoteIdPage;
