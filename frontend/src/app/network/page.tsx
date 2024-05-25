import NetworkPageComponent from "./NetworkPage";
import NetworkTable from "./NetworkTable";
import NetworkUser from "./types";

const users: NetworkUser[] = [
  {
    id: 1,
    name: "Liam Johnson",
    email: "liam@example.com",
    position: "Sale",
    addedOn: "2023-06-23",
    addedBy: "Olivia Smith",
  },
  {
    id: 2,
    name: "Olivia Smith",
    email: "olivia@exampe.com",
    position: "Refund",
    addedOn: "2023-06-24",
    addedBy: "Noah Williams",
  },
  {
    id: 3,
    name: "Noah Williams",
    email: "noah@example.com",
    position: "Subscription",
    addedOn: "2023-06-25",
    addedBy: "Emma Brown",
  },
  {
    id: 4,
    name: "Emma Brown",
    email: "emman@example.com",
    position: "Sale",
    addedOn: "2023-06-26",
    addedBy: "Liam Johnson",
  },
  {
    id: 5,
    name: "Liam Johnson",
    email: "liam@example.com",
    position: "Sale",
    addedOn: "2023-06-23",
    addedBy: "Olivia Smith",
  },
  {
    id: 6,
    name: "Olivia Smith",
    email: "olivia@exampe.com",
    position: "Refund",
    addedOn: "2023-06-24",
    addedBy: "Noah Williams",
  },
  {
    id: 7,
    name: "Noah Williams",
    email: "noah@example.com",
    position: "Subscription",
    addedOn: "2023-06-25",
    addedBy: "Emma Brown",
  },
  {
    id: 8,
    name: "Emma Brown",
    email: "emman@example.com",
    position: "Sale",
    addedOn: "2023-06-26",
    addedBy: "Liam Johnson",
  },
];

function NetworkPage() {
  return <NetworkPageComponent props={{ users: users }} />;
}

export default NetworkPage;
