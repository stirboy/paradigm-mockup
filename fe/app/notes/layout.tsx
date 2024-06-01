import React from "react";

import Sidebar from "./_components/sidebar";
import { SearchCommand } from "@/components/search-command";

const NotesLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="min-h-[calc(100%-64px)] flex">
      <Sidebar title="" />
      <SearchCommand />
      <main className="flex-1 overflow-y-auto">{children}</main>
    </div>
  );
};

export default NotesLayout;
