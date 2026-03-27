import { Suspense } from "react";
import { Fraunces } from "next/font/google";
import Link from "next/link";
import { SearchContent } from "@/components/searchContent";
import { Spinner } from "@/components/ui/spinner";

const fraunces = Fraunces({
  weight: ["400", "700"],
  subsets: ["latin"],
});

export default function SearchPage() {
  return (
    <div>
      <Link href="/" className="flex justify-center">
        <h1
          className={`my-6 ${fraunces.className} text-7xl font-bold text-stone-900`}
        >
          Museum Passport
        </h1>
      </Link>

      <Suspense
        fallback={
          <div className="flex justify-center py-32">
            <Spinner className="size-10 text-stone-300" />
          </div>
        }
      >
        <SearchContent />
      </Suspense>
    </div>
  );
}
