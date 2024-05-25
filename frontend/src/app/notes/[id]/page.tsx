import ClientNotePage from "./clientPage";

async function NotePage({ params }: { params: { id: string } }) {
  return <ClientNotePage noteId={params.id} />;
}

export default NotePage;
