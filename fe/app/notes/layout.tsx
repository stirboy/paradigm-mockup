import React from "react";
import { SearchCommand } from "@/components/search-command";
import Sidebar from "@/app/notes/_components/sidebar";

const NotesLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="min-h-[calc(100%-64px)] flex">
      <Sidebar title="" />
      <main className="flex-1 overflow-y-auto">
        <SearchCommand />
        {children}
      </main>
    </div>
  );
};

export default NotesLayout;
