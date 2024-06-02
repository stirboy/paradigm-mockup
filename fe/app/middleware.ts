import { NextRequest, NextResponse } from "next/server";
import { api } from "@/lib/restapi";
import { Routes } from "@/lib/constants/routes";
import { AxiosError } from "axios";
import { User } from "@/components/auth/user";

export default async function middleware(request: NextRequest) {
  let cookie = request.cookies.get("jwt");
  console.log("cookie", cookie);
  console.log("request.url", request.url);

  try {
    await api
      .get<User>(Routes.User, {
        withCredentials: true,
      })
      .then((res) => res.data);
  } catch (error) {
    const err = error as AxiosError;
    const status = err.response?.status;
    if (status === 401 || status === 403 || status === 404) {
      return NextResponse.redirect(new URL("/login", request.url));
    }
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    "/((?!api|login|_next/static|_next/image|favicon.ico).*)",
  ],
};
