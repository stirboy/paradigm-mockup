import { useParams } from "next/navigation";
import React from "react";
import useSWR from "swr";
import { Note } from "@/app/notes/_api/models";
import { Routes } from "@/lib/constants/routes";
import { MenuIcon } from "lucide-react";
import Title from "@/app/notes/_components/title";
import { Banner } from "@/app/notes/_components/banner";
import { toast } from "@/components/ui/use-toast";

interface NavbarProps {
  isCollapsed: boolean;
  onResetWidth: () => void;
}

const Navbar = ({ isCollapsed, onResetWidth }: NavbarProps) => {
  const params = useParams();
  const {
    data: note,
    isLoading,
    error,
  } = useSWR<Note>(`${Routes.Notes}/${params.noteId}`);

  if (isLoading) {
    return (
      <div
        className={
          "bg-background dark:bg-[#1F1F1F] px-3 pb-2 pt-20 w-full flex items-center"
        }
      >
        <Title.Skeleton />
      </div>
    );
  }

  if (!note) {
    return null;
  }

  console.log(note);

  return (
    <>
      <nav className="bg-background dark:bg-[#1F1F1F] px-3 pb-2 py-2 w-full flex items-center gap-x-4">
        {isCollapsed && (
          <MenuIcon
            role="button"
            onClick={onResetWidth}
            className="h-6 w-6 text-muted-foreground"
          />
        )}
        <div className="flex items-center justify-between w-full">
          <Title initialData={note!} />
        </div>
      </nav>
      {note!.isArchived && (
        <Banner noteId={note!.id} parentId={note!.parentId} />
      )}
    </>
  );
};

export default Navbar;
