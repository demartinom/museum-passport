import { fraunces } from "../../lib/fonts";
import Link from "next/link";
import { SearchContent } from "@/components/searchContent";
import { SearchResult } from "@/types/search";

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
      const apiUrl = `${process.env.API_URL}/api/search?${field}=${encodeURIComponent(
        query,
      )}&length=80&page=${page}`;
      const res = await fetch(apiUrl, {
        next: { revalidate: 60 },
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
          className={` ${fraunces.className} text-4xl font-bold text-stone-900`}
        >
          Museum Passport
        </h1>
      </Link>
      <SearchContent searchResult={initialResults.results || []} />
    </div>
  );
}
