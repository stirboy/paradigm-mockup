"use client";
import * as React from 'react';

type EditorToolBarProps = {
  title: string;
};
export const EditorToolbar = ({title}: EditorToolBarProps) => {
  return (
    <div>
      <div className="flex flex-col">
        <ul>
          <li>Note one</li>
          <li>Note two</li>
          <li>Note Note Three</li>
        </ul>
      </div>

    </div>
  );
};
