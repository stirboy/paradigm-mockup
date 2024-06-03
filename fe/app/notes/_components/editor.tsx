"use client";

import "@blocknote/core/fonts/inter.css";
import { BlockNoteEditor, PartialBlock } from "@blocknote/core";
import { BlockNoteView } from "@blocknote/mantine";
import "@blocknote/mantine/style.css";
import { useMemo } from "react";

import "./editor.css";
import { useTheme } from "next-themes";

interface NoteEditorProps {
  onChange: (content: string) => void;
  initialContent?: string;
  editable?: boolean;
}

const NoteEditor: React.FC<NoteEditorProps> = ({
  onChange,
  initialContent,
  editable,
}) => {
  const { resolvedTheme } = useTheme();

  const editor = useMemo(() => {
    return BlockNoteEditor.create({
      initialContent: initialContent
        ? (JSON.parse(initialContent) as PartialBlock[])
        : undefined,
    });
  }, [initialContent]);

  if (editor === undefined) {
    return "Loading content...";
  }

  return (
    <div className="w-full h-full">
      <BlockNoteView
        editor={editor}
        editable={editable}
        theme={resolvedTheme === "dark" ? "dark" : "light"}
        onChange={() => {
          onChange(JSON.stringify(editor.document));
        }}
        ref={undefined}
        onSelectionChange={() => {}}
      ></BlockNoteView>
    </div>
  );
};

export default NoteEditor;
