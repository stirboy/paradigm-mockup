"use client";

import React from "react";
import useSWR from "swr";
import { Routes } from "@/lib/constants/routes";
import { Note } from "@/app/notes/_api/models";
import Toolbar from "@/components/toolbar";
import dynamic from "next/dynamic";
import { api } from "@/lib/restapi";

const Editor = dynamic(() => import("../_components/editor"), { ssr: false });

interface NoteIdPageProps {
  params: {
    noteId: string;
  };
}

const NoteIdPage = ({ params }: NoteIdPageProps) => {
  const { data: note, isLoading } = useSWR<Note>(
    `${Routes.Notes}/${params.noteId}`,
    {
      revalidateOnFocus: false,
      revalidateOnMount: true,
      revalidateIfStale: false,
    },
  );

  if (isLoading) {
    return <div>Loading...</div>;
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
