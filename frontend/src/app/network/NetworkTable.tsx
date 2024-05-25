"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import NetworkUser from "./types";

interface NetworkTableProps {
  props: { users: NetworkUser[]; onUserClick: (user: NetworkUser) => void };
}

export default function NetworkTable({ props }: NetworkTableProps) {
  return (
    <Card>
      <CardHeader className="px-7">
        <CardTitle>User Network</CardTitle>
        <CardDescription>
          Where the top talent in the industry is connected.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead key={"user"}>User</TableHead>
              <TableHead key={"position"}>Position</TableHead>
              <TableHead key={"addedOn"} className="hidden sm:table-cell">
                Added On
              </TableHead>
              <TableHead key={"addedBy"} className="hidden sm:table-cell">
                Added By
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {props.users.map((user) => {
              return (
                <TableRow
                  key={user.id}
                  className="bg-accent"
                  onClick={() => {
                    console.log("Clicked");
                    props.onUserClick(user);
                  }}
                >
                  <TableCell>
                    <div className="font-medium">{user.name}</div>
                    <div className="hidden text-sm text-muted-foreground md:inline">
                      {user.email}
                    </div>
                  </TableCell>
                  <TableCell>{user.position}</TableCell>
                  <TableCell className="hidden sm:table-cell">
                    {user.addedOn}
                  </TableCell>
                  <TableCell className="hidden sm:table-cell">
                    {user.addedBy}
                  </TableCell>
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}
