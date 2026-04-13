"use client";

import { useState, useEffect, useTransition } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Image from "next/image";
import Link from "next/link";

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
  { value: "general", label: "All" },
  { value: "name", label: "Artwork" },
  { value: "artist", label: "Artist" },
];

interface SearchContentProps {
  initialResults: Art[];
}

export function SearchContent({ initialResults }: SearchContentProps) {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // URL is source of truth for committed state
  const urlQuery = searchParams.get("q") || "";
  const urlField = searchParams.get("field") || "general";

  // Local state is ONLY for typing
  const [searchText, setSearchText] = useState(urlQuery);

  // Sync input when URL changes (back/forward or navigation)
  useEffect(() => {
    setSearchText(urlQuery);
  }, [urlQuery]);

  function handleSearch(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault();

    const params = new URLSearchParams(searchParams.toString());
    params.set("q", searchText.trim());
    params.set("field", urlField);

    startTransition(() => {
      router.push(`/?${params.toString()}`);
    });
  }

  return (
    <div className="min-h-screen px-6 py-10">
      {/* SEARCH BAR */}
      <form onSubmit={handleSearch} className="mx-auto mb-4 max-w-2xl">
        <div className="flex overflow-hidden rounded-xl border border-stone-300 bg-white shadow-sm focus-within:ring-2 focus-within:ring-stone-200">
          {/* FIELD SELECT */}
          <Select
            value={urlField}
            onValueChange={(value) => {
              const params = new URLSearchParams(searchParams.toString());
              params.set("field", value);

              startTransition(() => {
                router.push(`/?${params.toString()}`);
              });
            }}
          >
            <SelectTrigger className="w-32 shrink-0 border-0 border-r border-stone-200 bg-stone-50 text-sm focus:ring-0">
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

          {/* INPUT */}
          <Input
            type="text"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
            className="flex-1 border-0 focus-visible:ring-0"
            placeholder="Search artworks..."
          />

          <button
            type="submit"
            className="bg-stone-900 px-5 text-white disabled:opacity-50"
            disabled={isPending}
          >
            {isPending ? "..." : "Search"}
          </button>
        </div>
      </form>

      {/* RESULTS */}
      <div className="relative mt-10">
        {/* Loading Overlay: Only visible when isPending is true */}
        {isPending && (
          <div className="absolute inset-0 z-10 flex items-start justify-center bg-white/40 pt-20 backdrop-blur-[1px]">
            <Spinner className="size-15 text-stone-900" />
          </div>
        )}

        <div
          className={
            isPending
              ? "pointer-events-none opacity-30 transition-opacity"
              : "opacity-100"
          }
        >
          {initialResults?.length > 0 ? (
            <div className="grid grid-cols-2 gap-5 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
              {initialResults.map((item) => (
                <Link href={`/art/${item.ID}`} key={item.ID} className="group">
                  <div className="relative h-56 w-full overflow-hidden rounded-lg bg-stone-100">
                    <Image
                      src={item.ImageSmall}
                      alt={item.ArtworkTitle}
                      fill
                      unoptimized
                      className="object-contain transition-transform group-hover:scale-105"
                    />
                  </div>

                  <h3 className="mt-2 text-sm font-semibold">
                    {item.ArtworkTitle}
                  </h3>
                  <p className="text-xs text-stone-400">{item.Museum}</p>
                </Link>
              ))}
            </div>
          ) : (
            urlQuery && (
              <p className="py-32 text-center text-stone-500">
                No results found for “{urlQuery}”.
              </p>
            )
          )}
        </div>
      </div>
    </div>
  );
}
