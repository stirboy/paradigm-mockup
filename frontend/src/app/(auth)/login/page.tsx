"use client";

import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import axios from "axios";
import { SyntheticEvent, useState } from "react";
import api from "@/utils/restapi";
import { toast } from "@/components/ui/use-toast";

function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const needsToast = localStorage.getItem("authToast");
  if (needsToast === "true") {
    toast({
      variant: "destructive",
      title: "Unauthorized",
      description: "You need to login to access this page.",
    });
    localStorage.removeItem("authToast");
  }

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    await api
      .post("http://localhost:8080/auth/login", {
        username,
        password,
      })
      .then((res) => {
        router.push("/notes");
        toast({
          variant: "default",
          title: "Login successful",
        });
      })
      .catch((err) => {
        console.error(err);
      });
  };

  return (
    <div className="w-full lg:grid lg:min-h-[600px] lg:grid-cols-2 xl:min-h-[800px]">
      <div className="flex items-center justify-center py-12">
        <div className="mx-auto grid w-[350px] gap-6">
          <div className="grid gap-2 text-center">
            <h1 className="text-3xl font-bold">Login</h1>
            <p className="text-balance text-muted-foreground">
              Enter your email below to login to your account
            </p>
          </div>
          <div className="grid gap-4">
            <div className="grid gap-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                placeholder="m@example.com"
                required
                onChangeCapture={(e: React.FormEvent<HTMLInputElement>) => {
                  console.log("username=" + e.currentTarget.value);
                  setUsername(e.currentTarget.value);
                }}
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="password">Password</Label>
                <Link
                  href="/forgot-password"
                  className="ml-auto inline-block text-sm underline"
                >
                  Forgot your password?
                </Link>
              </div>
              <Input
                id="password"
                type="password"
                required
                onChangeCapture={(e: React.FormEvent<HTMLInputElement>) => {
                  console.log("password=" + e.currentTarget.value);
                  setPassword(e.currentTarget.value);
                }}
              />
            </div>
            <form onSubmit={submit}>
              <Button type="submit" className="w-full">
                Login
              </Button>
              <Button variant="outline" className="w-full">
                Login with Google
              </Button>
            </form>
          </div>
          <div className="mt-4 text-center text-sm">
            Don&apos;t have an account?{" "}
            <Link href="#" className="underline">
              Sign up
            </Link>
          </div>
        </div>
      </div>
      <div className="hidden bg-muted lg:block">
        <Image
          src="/placeholder.svg"
          alt="Image"
          width="1920"
          height="1080"
          className="h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
        />
      </div>
    </div>
  );
}

export default LoginPage;
