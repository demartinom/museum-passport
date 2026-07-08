import { fraunces } from "../lib/fonts";
import Link from "next/link";
import { SearchContent } from "@/components/searchContent";
import { SearchResult } from "@/types/search";
import AOTD from "@/components/aotd";

export const dynamic = "force-dynamic";

export default async function SearchPage({
  searchParams,
}: {
  searchParams: Promise<{ q?: string; field?: string; page?: string }>;
}) {
  const params = await searchParams;
  const query = params.q || "";
  const field = params.field || "general";
  const page = params.page || "1";

  let initialResults: SearchResult = { totalPages: 0, results: [] };

  if (query) {
    try {
      const apiUrl = `${process.env.API_URL}/api/search?${field}=${encodeURIComponent(query)}&length=80&page=${page}`;

      const res = await fetch(apiUrl, {
        next: { revalidate: 60 }, // Caches the result on the server for 1 minute
      });

      initialResults = await res.json();
    } catch (err) {
      console.error("Server-side fetch error:", err);
    }
  }
  return (
    <div>
      <Link href="/" className="flex justify-center">
        <h1
          className={`my-6 ${fraunces.className} text-7xl font-bold text-stone-900`}
        >
          Museum Passport
        </h1>
      </Link>

      <AOTD />
      <SearchContent searchResult={initialResults.results || []} />
    </div>
  );
}
