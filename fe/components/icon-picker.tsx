"use client";

import EmojiPicker, { Theme } from "emoji-picker-react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { useState } from "react";

interface IconPickerProps {
  onChange: (icon: string) => void;
  children: React.ReactNode;
  asChild?: boolean;
}

export const IconPicker = ({
  onChange,
  children,
  asChild,
}: IconPickerProps) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <Popover open={isOpen} onOpenChange={setIsOpen}>
      <PopoverTrigger asChild={asChild}>{children}</PopoverTrigger>
      <PopoverContent
        align={"start"}
        className={"p-0 w-full border-none shadow-none"}
      >
        <EmojiPicker
          height={350}
          theme={Theme.LIGHT}
          onEmojiClick={(data) => {
            onChange(data.emoji);
            setIsOpen(false);
          }}
        />
      </PopoverContent>
    </Popover>
  );
};
