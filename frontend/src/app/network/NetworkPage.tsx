"use client";

import { Tabs } from "@/components/ui/tabs";
import { useState } from "react";
import NetworkTable from "./NetworkTable";
import NetworkTableButtons from "./NetworkTableButtons";
import NetworkUser from "./types";
import NetworkUserDetails from "./NetworkUserDetails";

interface NetworkPageProps {
  props: { users: NetworkUser[] };
}

function NetworkPageComponent({ props }: NetworkPageProps) {
  const [selectedUser, setSelectedUser] = useState<NetworkUser | null>(null);

  const onUserClick = (user: NetworkUser) => {
    console.log(user);
    setSelectedUser(user);
  };

  return (
    <div className="flex min-h-screen w-full flex-col bg-muted/40">
      <div className="flex flex-col sm:gap-4 sm:py-4">
        <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8 lg:grid-cols-1 2xl:grid-cols-3">
          <div className="grid auto-rows-max items-start gap-4 md:gap-8 lg:col-span-2">
            <Tabs>
              <NetworkTableButtons />
              <NetworkTable
                props={{ users: props.users, onUserClick: onUserClick }}
              />
            </Tabs>
          </div>
        </main>
      </div>
    </div>
  );
}

export default NetworkPageComponent;
