"use client";

import { useRouter } from "next/navigation";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { SyntheticEvent, useState } from "react";
import { toast } from "@/components/ui/use-toast";
import { useLogin } from "@/components/auth/auth";
import { AxiosError } from "axios";

function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();
  const { trigger } = useLogin();

  const submit = async (e: SyntheticEvent) => {
    e.preventDefault();

    await trigger({
      username,
      password,
    })
      .then((res) => {
        router.push("/notes");
      })
      .catch((error) => {
        const err = error as AxiosError;
        const status = err.response?.status;
        if (status === 404) {
          toast({
            variant: "destructive",
            title: "User not found",
            description: "The data you entered is incorrect",
          });
        }
      });
  };

  return (
    <div className="w-full lg:min-h-[600px] xl:min-h-[800px]">
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
                  setUsername(e.currentTarget.value);
                }}
              />
            </div>
            <div className="grid gap-2">
              <div className="flex items-center">
                <Label htmlFor="password">Password</Label>
              </div>
              <Input
                id="password"
                type="password"
                required
                onChangeCapture={(e: React.FormEvent<HTMLInputElement>) => {
                  setPassword(e.currentTarget.value);
                }}
              />
            </div>
            <form onSubmit={submit}>
              <Button type="submit" className="w-full">
                Login
              </Button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}

export default LoginPage;
