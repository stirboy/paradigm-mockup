"use client";

import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { Routes } from "@/lib/constants/routes";
import { api } from "@/lib/restapi";
import { PlusCircle } from "lucide-react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { SyntheticEvent } from "react";
import { useSWRConfig } from "swr";
import { useCreateNote } from "./_hooks/createNote";

const NotesPage = () => {
  const router = useRouter();
  const { toast } = useToast();
  const { trigger, isMutating } = useCreateNote();

  const createNote = async (e: SyntheticEvent) => {
    e.preventDefault();
    trigger()
      .then((res) => {
        console.log(res.data);
        toast({
          title: "Note created",
          description: "Your new note has been created",
        });
        // router.push(`/notes/${res.data.id}`);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <div className="h-full flex flex-col items-center justify-center space-y-4">
      <Image
        src="/empty.png"
        width={300}
        height={300}
        alt="Empty"
        className="dark::hidden w-auto h-auto"
      />
      <Image
        src="/empty-dark.png"
        width={300}
        height={300}
        alt="Empty"
        className="hidden dark:block w-auto h-auto"
      />
      <Button onClick={createNote}>
        <PlusCircle className="h-6 w-6 mr-2" />
        Create note
      </Button>
    </div>
  );
};

export default NotesPage;
