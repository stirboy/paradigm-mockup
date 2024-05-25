"use client"

import {EditorToolbar} from "@/components/editor-toolbar";

function NewNotes() {
  return (
    <div className="hidden md:block w-[20%] h-[100vh] border-[3px] box-border overflow-y-auto">
      <EditorToolbar title="new note" />
    </div>
  );
}

export default NewNotes;
