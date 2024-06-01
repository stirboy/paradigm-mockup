"use client";

import React from "react";
import { Routes } from "@/lib/constants/routes";
import Toolbar from "@/components/toolbar";
import dynamic from "next/dynamic";
import { api } from "@/lib/restapi";
import { useNote } from "@/app/notes/_hooks/notes-api";

const Editor = dynamic(() => import("../_components/editor"), { ssr: false });

interface NoteIdPageProps {
  params: {
    noteId: string;
  };
}

const NoteIdPage = ({ params }: NoteIdPageProps) => {
  const { note, isLoading } = useNote(params.noteId);

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
