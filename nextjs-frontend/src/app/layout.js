// import { Inter } from "next/font/google";
import "./globals.css";



// const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: ">>>justlinks app",
  description: "created by https://github.com/x-MrPhillips-x",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  );
}
