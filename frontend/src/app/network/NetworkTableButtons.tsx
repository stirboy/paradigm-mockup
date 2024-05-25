"use client";
import { Import, ListFilter } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

function NetworkTableButtons() {
  return (
    <div className="flex items-center pb-2">
      <div className="ml-auto flex items-center gap-2">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" size="sm" className="h-7 gap-1 text-sm">
              <ListFilter className="h-3.5 w-3.5" />
              <span className="sr-only sm:not-sr-only">Filter</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Filter by</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuCheckboxItem checked>
              Fulfilled
            </DropdownMenuCheckboxItem>
            <DropdownMenuCheckboxItem>Declined</DropdownMenuCheckboxItem>
            <DropdownMenuCheckboxItem>Refunded</DropdownMenuCheckboxItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <Button size="sm" variant="outline" className="h-7 gap-1 text-sm">
          <Import className="h-3.5 w-3.5" />
          <span className="sr-only sm:not-sr-only">Import</span>
        </Button>
      </div>
    </div>
  );
}

export default NetworkTableButtons;
