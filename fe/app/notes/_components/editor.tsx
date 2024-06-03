"use client";

import "@blocknote/core/fonts/inter.css";
import {
  BasicTextStyleButton,
  BlockColorsItem,
  BlockTypeSelect,
  ColorStyleButton,
  CreateLinkButton,
  DragHandleMenu,
  FileCaptionButton,
  FileReplaceButton,
  FormattingToolbar,
  FormattingToolbarController,
  NestBlockButton,
  RemoveBlockItem,
  SideMenu,
  SideMenuController,
  TextAlignButton,
  UnnestBlockButton,
} from "@blocknote/react";
import { BlockNoteEditor, PartialBlock } from "@blocknote/core";
import { BlockNoteView } from "@blocknote/mantine";
import "@blocknote/mantine/style.css";
import { useMemo } from "react";

import "./editor.css";
import { useTheme } from "next-themes";
import { useSearch } from "@/app/notes/_hooks/use-search";

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
  const { isOpen } = useSearch();

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
        formattingToolbar={false}
        sideMenu={false}
        editor={editor}
        editable={editable}
        theme={resolvedTheme === "dark" ? "dark" : "light"}
        onChange={() => {
          onChange(JSON.stringify(editor.document));
        }}
        ref={undefined}
        onSelectionChange={() => {}}
      >
        {!isOpen && (
          <>
            <FormattingToolbarController
              formattingToolbar={() => (
                <FormattingToolbar>
                  <BlockTypeSelect key={"blockTypeSelect"} />

                  <FileCaptionButton key={"fileCaptionButton"} />
                  <FileReplaceButton key={"replaceFileButton"} />

                  <BasicTextStyleButton
                    basicTextStyle={"bold"}
                    key={"boldStyleButton"}
                  />
                  <BasicTextStyleButton
                    basicTextStyle={"italic"}
                    key={"italicStyleButton"}
                  />
                  <BasicTextStyleButton
                    basicTextStyle={"underline"}
                    key={"underlineStyleButton"}
                  />
                  <BasicTextStyleButton
                    basicTextStyle={"strike"}
                    key={"strikeStyleButton"}
                  />
                  {/* Extra button to toggle code styles */}
                  <BasicTextStyleButton
                    key={"codeStyleButton"}
                    basicTextStyle={"code"}
                  />

                  <TextAlignButton
                    textAlignment={"left"}
                    key={"textAlignLeftButton"}
                  />
                  <TextAlignButton
                    textAlignment={"center"}
                    key={"textAlignCenterButton"}
                  />
                  <TextAlignButton
                    textAlignment={"right"}
                    key={"textAlignRightButton"}
                  />

                  <ColorStyleButton key={"colorStyleButton"} />

                  <NestBlockButton key={"nestBlockButton"} />
                  <UnnestBlockButton key={"unnestBlockButton"} />

                  <CreateLinkButton key={"createLinkButton"} />
                </FormattingToolbar>
              )}
            />
            <SideMenuController
              sideMenu={(props) => (
                <SideMenu
                  {...props}
                  dragHandleMenu={(props) => (
                    <DragHandleMenu {...props}>
                      <RemoveBlockItem {...props}>Delete</RemoveBlockItem>
                      <BlockColorsItem {...props}>Colors</BlockColorsItem>
                    </DragHandleMenu>
                  )}
                />
              )}
            />
          </>
        )}
      </BlockNoteView>
    </div>
  );
};

export default NoteEditor;
