import { fraunces } from "@/lib/fonts";
import Link from "next/link";

export const NavBar = () => {
  return (
    <div className="mb-3 bg-gray-100 p-5">
      <Link href="/">
        <h1 className={`${fraunces.className} text-xl`}>Museum Passport</h1>
      </Link>
    </div>
  );
};
