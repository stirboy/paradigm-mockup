"use client";

import React from "react";
import EditorSidebar from "./_components/sidebar";
import Image from "next/image";

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
    </div>
  );
};

export default Notes;
