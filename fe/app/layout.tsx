import NavigationBar from "@/components/navigation-bar";
import { Toaster } from "@/components/ui/toaster";
import { SWRProvider } from "@/providers/swr-provider";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { SearchCommand } from "@/components/search-command";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <SWRProvider>
          <NavigationBar />
          <SearchCommand />
          {children}
        </SWRProvider>
        <Toaster />
      </body>
    </html>
  );
}
