"use client";

import { Note } from "@/app/notes/_api/models";
import { IconPicker } from "./icon-picker";
import { Button } from "@/components/ui/button";
import { ImageIcon, Smile, X } from "lucide-react";
import { ElementRef, useRef, useState } from "react";
import {
  useRemoveNoteIcon,
  useUpdateNote,
  useUpdateNoteIcon,
} from "@/app/notes/_hooks/notes-api";
import TextareaAutosize from "react-textarea-autosize";

interface ToolbarProps {
  initialData: Note;
  preview?: boolean;
}

const Toolbar = ({ initialData, preview }: ToolbarProps) => {
  const inputRef = useRef<ElementRef<"textarea">>(null);
  const [isEditing, setIsEditing] = useState(false);
  const [value, setValue] = useState(initialData.title);
  const { trigger: updateNoteIcon } = useUpdateNoteIcon();
  const { trigger: removeNoteIcon } = useRemoveNoteIcon();

  const { trigger: updateNote } = useUpdateNote();

  const enableInput = () => {
    if (preview) return;

    setIsEditing(true);
    setTimeout(() => {
      setValue(initialData.title);
      inputRef.current?.focus();
    });
  };

  const disableInput = () => {
    setIsEditing(false);
  };

  const onInput = (value: string) => {
    setValue(value);
    updateNote(initialData.id, value || "Untitled", initialData.parentId);
  };

  const onKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === "Enter") {
      e.preventDefault();
      disableInput();
    }
  };

  const onIconSelect = (icon: string) => {
    updateNoteIcon(initialData.id, icon, initialData.parentId);
  };

  const onRemoveIcon = () => {
    removeNoteIcon(initialData.id);
  };

  return (
    <div className={"pl-[54px] group relative"}>
      {!!initialData.icon && !preview && (
        <div className={"flex items-center gap-x-2 group/icon pt-6"}>
          <IconPicker onChange={onIconSelect}>
            <p className={"text-6xl hover:opacity-75 transition"}>
              {initialData.icon}
            </p>
          </IconPicker>
          <Button
            onClick={onRemoveIcon}
            className={
              "rounded-full opacity-0 group-hover/icon:opacity-100 transition text-muted-foreground text-xs"
            }
            variant={"outline"}
          >
            <X className={"h-4 w-4"} />
          </Button>
        </div>
      )}
      {!!initialData.icon && preview && (
        <p className={"text-6xl pt-6"}>{initialData.icon}</p>
      )}
      <div
        className={
          "opacity-0 group-hover:opacity-100 flex items-center gap-x-1 py-4"
        }
      >
        {!initialData.icon && !preview && (
          <IconPicker asChild onChange={onIconSelect}>
            <Button
              className={"text-muted-foreground text-xs"}
              variant={"outline"}
              size={"sm"}
            >
              <Smile className={"h-4 w-4 mr-2"} />
              Add icon
            </Button>
          </IconPicker>
        )}
        {!initialData.coverImage && !preview && (
          <Button
            onClick={() => {}}
            className={"text-muted-foreground text-xs"}
            variant={"outline"}
            size={"sm"}
          >
            <ImageIcon className={"h-4 w-4 mr-2"} />
            Add cover
          </Button>
        )}
      </div>
      {isEditing && !preview ? (
        <TextareaAutosize
          ref={inputRef}
          onBlur={disableInput}
          onKeyDown={onKeyDown}
          value={value}
          onChange={(e) => onInput(e.target.value)}
          className="text-5xl bg-transparent font-bold break-words outline-none text-[#3F3F3F] dark:text-[#CFCFCF] resize-none"
        />
      ) : (
        <div
          onClick={enableInput}
          className="pb-[11.5px] text-5xl font-bold break-words outline-none text-[#3F3F3F] dark:text-[#CFCFCF]"
        >
          {initialData.title}
        </div>
      )}
    </div>
  );
};

export default Toolbar;
