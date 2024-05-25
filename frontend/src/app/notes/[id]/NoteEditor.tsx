"use client";

import {Block, BlockNoteEditor, PartialBlock} from "@blocknote/core";
import "@blocknote/core/fonts/inter.css";
import "@blocknote/mantine/style.css";
import {BlockNoteView} from "@blocknote/shadcn";
import {useMemo} from "react";

import "./index.css"
import {useTheme} from "next-themes";

interface NoteEditorProps {
  onChange: (document: Block[]) => void;
  initialContent?: string;
  editable?: boolean;
}

const NoteEditor: React.FC<NoteEditorProps> = ({
                                                 onChange,
                                                 initialContent,
                                                 editable,
                                               }) => {
  const theme = useTheme();
  const editor = useMemo(() => {
    if (initialContent === "loading") {
      return undefined;
    }
    return BlockNoteEditor.create({initialContent: initialContent ? JSON.parse(initialContent) as PartialBlock[] : undefined});
  }, [initialContent]);

  if (editor === undefined) {
    return "Loading content...";
  }

  return (
    <div className="w-full h-full">
      <BlockNoteView
        editor={editor}
        editable={true}
        theme={theme.theme === "dark" ? "dark" : "light"}
        onChange={() => {
          onChange(editor.document);
        }}
        ref={undefined}
        onSelectionChange={() => {
        }}
      />
    </div>
  );
};

export default NoteEditor;
