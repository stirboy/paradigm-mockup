"use client";

import {Button} from "@/components/ui/button";
import {Card, CardContent, CardTitle} from "@/components/ui/card";
import {PlusIcon} from "lucide-react";

import {useRouter} from "next/navigation";
import {SyntheticEvent} from "react";
import useSWR from "swr";
import api from "@/utils/restapi";

interface Note {
  id: number;
  title: string;
  content: string;
}

function NotesPage() {
  //const [notes, setNotes] = useState<Note[]>([]);
  const router = useRouter();

  const fetcher = (url: string) => api.get(url).then((res) => res);
  const {data, error} = useSWR(`http://localhost:8080/api/notes`, fetcher);

  if (error) {
    throw error;
  }

  if (data?.data === undefined) {
    //alert("No notes found");
    return;
  }

  const notes: Note[] = data?.data.map((d: any) => {
    return {
      id: d.id,
      title: d.title,
      // JSON.parse(d.content).map((x: any) =>
      //   x.content.map((c: any) => c.text)
      // )[0] || "Untitled",
      content: JSON.parse(d.content)
        .map((x: any) => x.content.map((c: any) => c.text))
        .join("\n"),
    } as Note;
  });

  // useEffect(() => {
  //   api
  //     .get("http://localhost:8080/api/notes")
  //     .then((res) => {
  //       const n = res.data.map((d: any) => {
  //         return {
  //           id: d.id,
  //           title: JSON.parse(d.content).map((x: any) => x.content.map((c: any) => c.text))[0] || "Untitled",
  //           content: JSON.parse(d.content).map((x: any) => x.content.map((c: any) => c.text)).join("\n"),
  //         } as Note;
  //       });
  //       setNotes(n);
  //     })
  //     .catch((err) => {
  //       console.error(err);
  //     });
  // }, []);

  const createNote = async (e: SyntheticEvent) => {
    e.preventDefault();
    await api
      .post("http://localhost:8080/api/notes", {
        title: "Untitled",
        // content: JSON.stringify([{
        //   type: "paragraph",
        // }],),
        // content: JSON.stringify([
        //   {
        //     type: "heading",
        //     props: {
        //       textColor: "default",
        //       backgroundColor: "default",
        //       textAlignment: "left",
        //       level: 1,
        //     },
        //     content: [],
        //     children: [],
        //   },
        // ]),
      })
      .then((res) => {
        console.log(res.data.id);
        router.push(`/notes/${res.data.id}`);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <section className="py-12">
      <div className="container mt-auto px-4">
        <div className="fixed bottom-3 right-4 h-15">
          <form onSubmit={createNote}>
            <Button
              type="submit"
              size="sm"
              variant="outline"
              className="gap-1 text-sm py-2 px-4 rounded-full shadow-lg size-full h-16 w-16 md:h-full md:w-full"
            >
              <div className="flex justify-between">
                <PlusIcon className="mt-0.5 h-3.5 w-3.5"/>
                <span className="sm:hidden sr-only sm:not-sr-only">
                  Create note
                </span>
              </div>
            </Button>
          </form>
        </div>
        <div
          className="grid auto-rows-fr grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6">
          {notes.map((note) => {
            return (
              <Card
                key={note.id}
                className="group overflow-hidden rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300"
                onClick={() => {
                  router.push(`/notes/${note.id}`);
                }}
              >
                <CardContent className="p-4 flex flex-col flex-1">
                  <CardTitle className="text-lg font-semibold mb-2 text-">
                    {note.title}
                  </CardTitle>
                  <p className="text-gray-500 dark:text-gray-400 flex-1 truncate">
                    {note.content}
                  </p>
                </CardContent>
              </Card>
            );
          })}
        </div>
      </div>
    </section>
  );
}

export default NotesPage;
