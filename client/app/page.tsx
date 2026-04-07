import { Suspense } from "react";
import { fraunces } from "../lib/fonts";
import Link from "next/link";
import { SearchContent } from "@/components/searchContent";
import { Spinner } from "@/components/ui/spinner";

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
