export const enum Routes {
  Home = "/",
  User = "api/auth/user",
  Login = "api/auth/login",
  Notes = "api/notes",
  ArchivedNotes = `${Routes.Notes}?archived=true`,
}
