"use client";

import dynamic from "next/dynamic";
import api from "@/utils/restapi";
import useSWR from "swr";
import {PartialBlock} from "@blocknote/core";
import Title from "@/app/notes/[id]/title";

const Editor = dynamic(() => import("./NoteEditor"), {ssr: false});

interface ClientNotePageProps {
  noteId: string;
}

const ClientNotePage: React.FC<ClientNotePageProps> = ({noteId}) => {
  const fetcher = async (url: string) =>
    await api
      .get(url)
      .then((res) => {
        console.log(res.data?.content)
        return res.data;
      })
      .catch((err) => console.error(err));
  const {data} = useSWR(`http://localhost:8080/api/notes/${noteId}`, fetcher);

  if (data === undefined) {
    return <div>Loading...</div>;
  }

  const onTitleChange = async (title: string) => {
    await api.put(`http://localhost:8080/api/notes/${noteId}`, {title: title});
  }

  const onContentChange = async (doc: PartialBlock[]) => {
    await api.put(`http://localhost:8080/api/notes/${noteId}`, {content: JSON.stringify(doc)});
  }


  return (
    <main className="min-h-screen">
      <div className="flex flex-row justify-between gap-10 px-10 md:px-15 py-10 w-full">
        <div className="w-[85%]">
          <Title initialValue = {data?.title} onChange={onTitleChange}/>
          <div className="px-0 mt-8">
            <Editor
              initialContent={data?.content}
              onChange={async (doc) => {
                await onContentChange(doc);
              }}
            />
          </div>
        </div>
      </div>
    </main>
  );

}

export default ClientNotePage;
