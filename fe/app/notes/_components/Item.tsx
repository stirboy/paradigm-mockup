"use client";
import { Skeleton } from "@/components/ui/skeleton";
import { cn } from "@/lib/utils";
import { ChevronDown, ChevronRight, LucideIcon, Plus } from "lucide-react";
import React from "react";
import { useCreateNote, useCreateNoteWithParent } from "../_hooks/createNote";
import { useRouter } from "next/navigation";

interface ItemProps {
  id?: string;
  documentIcon?: string;
  active?: boolean;
  expanded?: boolean;
  isSearch?: boolean;
  level?: number;
  label: string;
  icon: LucideIcon;
  onExpand?: () => void;
  onClick: () => void;
}

const Item = ({
  id,
  label,
  onClick,
  icon: Icon,
  active,
  documentIcon,
  isSearch,
  level = 0,
  onExpand,
  expanded,
}: ItemProps) => {
  const ChevronIcon = expanded ? ChevronDown : ChevronRight;
  const { trigger: createNoteWithParent } = useCreateNoteWithParent();
  const router = useRouter();

  const onCreate = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    e.stopPropagation();
    if (!id) return;

    createNoteWithParent(id).then((res) => {
      if (!expanded) onExpand?.();

      // router.push(`/notes/${res.data.id}`);
    });
  };

  const handleExpand = (e: React.MouseEvent<HTMLDivElement, MouseEvent>) => {
    e.stopPropagation();
    onExpand?.();
  };

  return (
    <div
      onClick={onClick}
      role="button"
      style={{
        paddingLeft: level ? `${level * 12 + 12}px` : "12px",
      }}
      className={cn(
        "group min-h-[27px] text-sm py-1 pr-3 w-full hover:bg-primary/5 flex items-center text-muted-foreground font-medium",
        active && "bg-primary/5 text-primary",
      )}
    >
      {!!id && (
        <div
          role="button"
          className="h-full rounded-sm hover:bg-neutral-300 dark:bg-neutral-600 mr-1"
          onClick={handleExpand}
        >
          <ChevronIcon className="h-4 w-4 shrink-0 text-muted-foreground/50" />
        </div>
      )}

      {documentIcon ? (
        <div className="shrink-0 mr-2 text-[18px">{documentIcon}</div>
      ) : (
        <Icon className="shrink-0 h-[18px] mr-2 text-muted-foreground" />
      )}

      <span className="truncate"> {label}</span>
      {isSearch && (
        <kbd className="ml-auto pointer-events-none inline-flex h-5 select-none items-center gap-1 rounded border bg-muted px-1.5 font-mono text-[10px] font-medium text-muted-foreground opacity-100">
          <span className="text-xs">&#8984;</span>K
        </kbd>
      )}
      {!!id && (
        <div className="ml-auto flex items-center gap-x-2">
          <div
            role="button"
            onClick={onCreate}
            className="opacity-0 group-hover:opacity-100 h-full ml-auto rounder-sm hover:bg-neutral-300 dark:hover:bg-neutral-600"
          >
            <Plus className="h-4 w-4 text-muted-foreground" />
          </div>
        </div>
      )}
    </div>
  );
};

Item.ItemSkeleton = function ItemSkeleton({ level }: { level?: number }) {
  return (
    <div
      style={{
        paddingLeft: level ? `${level * 12 + 25}px` : "12px",
      }}
    >
      <Skeleton className="h-4 w-4" />
      <Skeleton className="h-3 w-[30%]" />
    </div>
  );
};

export default Item;
