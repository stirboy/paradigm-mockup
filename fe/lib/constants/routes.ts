export const enum Routes {
  Home = "/",
  Login = "api/auth/login",
  Notes = "api/notes",
  ArchivedNotes = `${Routes.Notes}?archived=true`,
  NoteById = "/notes/[id]",
}
