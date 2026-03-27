"use client";

import { useEffect, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Image from "next/image";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Spinner } from "@/components/ui/spinner";
import { Art } from "@/types/search";

const FIELD_OPTIONS = [
  {
    value: "general",
    label: "All",
    placeholder: "Search artworks, artists...",
  },
  { value: "name", label: "Artwork", placeholder: "e.g. Starry Night..." },
  {
    value: "artist",
    label: "Artist",
    placeholder: "e.g. Monet, Frida Kahlo...",
  },
];

export function SearchContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const urlField = searchParams.get("field");
  const urlQuery = searchParams.get("q");

  const [field, setField] = useState(urlField || "general");
  const [searchText, setSearchText] = useState(urlQuery || "");
  const [results, setResults] = useState<Array<Art>>([]);
  const [searching, setSearching] = useState(false);

  const hasSearched = !!urlQuery;
  const activePlaceholder =
    FIELD_OPTIONS.find((o) => o.value === field)?.placeholder ?? "Search...";

  useEffect(() => {
    if (!urlField || !urlQuery) return;

    async function fetchResults() {
      setSearching(true);
      setResults([]);
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/search?${urlField}=${encodeURIComponent(urlQuery!)}&length=80`,
        );
        const data = await res.json();
        setResults(data || []);
      } catch (err) {
        console.error(err);
      } finally {
        setSearching(false);
      }
    }
    fetchResults();
  }, [urlField, urlQuery]);

  function handleSearch(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!field || !searchText.trim()) return;
    router.push(`/?field=${field}&q=${encodeURIComponent(searchText.trim())}`);
  }

  return (
    <div className="min-h-screen px-6 py-10">
      <form onSubmit={handleSearch} className="mx-auto mb-4 max-w-2xl">
        <div className="flex overflow-hidden rounded-xl border border-stone-300 bg-white shadow-sm transition-all focus-within:border-stone-500 focus-within:ring-2 focus-within:ring-stone-200">
          <Select
            value={field}
            onValueChange={(val) => {
              setField(val);
              setSearchText("");
            }}
          >
            <SelectTrigger className="w-32 shrink-0 rounded-none border-0 border-r border-stone-200 bg-stone-50 px-3 text-sm font-medium text-stone-600 focus:ring-0">
              <SelectValue placeholder="Field" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                {FIELD_OPTIONS.map((opt) => (
                  <SelectItem key={opt.value} value={opt.value}>
                    {opt.label}
                  </SelectItem>
                ))}
              </SelectGroup>
            </SelectContent>
          </Select>

          <Input
            type="text"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
            placeholder={activePlaceholder}
            className="flex-1 rounded-none border-0 bg-white px-4 text-sm focus-visible:ring-0 focus-visible:ring-offset-0"
          />
          <button
            type="submit"
            disabled={!searchText.trim()}
            className="shrink-0 bg-stone-900 px-5 text-sm font-medium text-white transition-colors hover:bg-stone-700 disabled:opacity-40"
          >
            Search
          </button>
        </div>
      </form>

      {/* Results Logic */}
      {searching ? (
        <div className="flex flex-col items-center justify-center py-32">
          <Spinner className="size-10" />
        </div>
      ) : results.length > 0 ? (
        <div className="grid grid-cols-2 gap-5 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
          {results.map((item) => (
            <a href={`/art/${item.ID}`} key={item.ID} className="group">
              <div className="relative h-56 w-full overflow-hidden rounded-lg bg-stone-100">
                <Image
                  src={item.ImageSmall}
                  alt={item.ArtworkTitle}
                  fill
                  className="object-contain transition-transform group-hover:scale-105"
                />
              </div>
              <h3 className="mt-2 truncate text-sm font-semibold">
                {item.ArtworkTitle || "Untitled"}
              </h3>
            </a>
          ))}
        </div>
      ) : (
        hasSearched && (
          <p className="py-32 text-center text-stone-500">No results found.</p>
        )
      )}
    </div>
  );
}
