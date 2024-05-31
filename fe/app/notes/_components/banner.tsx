"use client";

import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { ConfirmModal } from "@/components/modals/confirm-modal";
import { useDeleteNote, useRestoreNotes } from "@/app/notes/_hooks/createNote";
import { toast } from "@/components/ui/use-toast";

interface BannerProps {
  noteId: string;
  parentId?: string;
}

export const Banner = ({ noteId, parentId }: BannerProps) => {
  const router = useRouter();

  const { trigger: deleteNote } = useDeleteNote();
  const { trigger: restoreNotes } = useRestoreNotes();

  const onRemove = () => {
    deleteNote(noteId).then((r) => {});

    toast({
      title: "Note deleted",
      variant: "destructive",
      description: "Your note has been deleted.",
      duration: 2000,
    });

    router.push("/notes");
  };

  const onRestore = () => {
    restoreNotes(noteId, parentId).then(() => {
      toast({
        title: "Note restored",
        description: "Your note has been restore.",
        duration: 2000,
      });
    });
  };

  return (
    <div className="w-full bg-rose-500 text-center text-sm p-2 text-white flex items-center gap-x-2 justify-center">
      <p>This page is in the Trash.</p>
      <Button
        size="sm"
        onClick={onRestore}
        variant="outline"
        className="border-white bg-transparent hover:bg-primary/5 text-white hover:text-white p-1 px-2 h-auto font-normal"
      >
        Restore page
      </Button>
      <ConfirmModal onConfirm={onRemove}>
        <Button
          size="sm"
          variant="outline"
          className="border-white bg-transparent hover:bg-primary/5 text-white hover:text-white p-1 px-2 h-auto font-normal"
        >
          Delete forever
        </Button>
      </ConfirmModal>
    </div>
  );
};
