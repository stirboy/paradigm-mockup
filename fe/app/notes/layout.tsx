import React from "react";

import Sidebar from "./_components/sidebar";

const NotesLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="h-full flex">
      <Sidebar title="" />
      <main className="flex-1 h-full overflow-y-auto">{children}</main>
    </div>
  );
};

export default NotesLayout;
