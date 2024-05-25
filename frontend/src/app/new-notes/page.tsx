"use client";

import { EditorToolbar } from "@/components/editor-toolbar";

function NewNotes() {
  return (
    <div className="hidden md:flex md:flex-col w-[20%] min-h-full border-[3px] box-border">
      <EditorToolbar title="new note" />
    </div>
  );
}

export default NewNotes;
