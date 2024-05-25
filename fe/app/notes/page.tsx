"use client";

import React from "react";
import EditorSidebar from "./_components/sidebar";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import { PlusCircle } from "lucide-react";

const Notes = () => {
  return (
    <div className="h-full flex flex-col items-center justify-center space-y-4">
      <Image
        src="/empty.png"
        width={300}
        height={300}
        alt="Empty"
        className="dark::hidden"
      />
      <Image
        src="/empty-dark.png"
        width={300}
        height={300}
        alt="Empty"
        className="hidden dark:block"
      />
      <Button>
        <PlusCircle className="h-6 w-6 mr-2" />
        Create note
      </Button>
    </div>
  );
};

export default Notes;
