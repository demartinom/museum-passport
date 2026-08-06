"use client";

import { useState, useEffect, useTransition } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import { Spinner } from "@/components/ui/spinner";
import { Art } from "@/types/search";
import SearchBar from "./searchBar";
import SearchPagination from "./searchPagination";

interface SearchContentProps {
  searchResult: Art[];
}

export function SearchContent({ searchResult }: SearchContentProps) {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isPending, startTransition] = useTransition();

  // URL is source of truth for committed state
  const urlQuery = searchParams.get("q") || "";
  const urlField = searchParams.get("field") || "general";
  const urlPage = searchParams.get("page") || "1";

  // Converts pageNumber into integer
  const pageNumber = Number(urlPage);
  // Max page user is allowed to go to
  const maxPageNext = 10;

  // Local state is ONLY for typing. nothing fires until form submit
  const [searchText, setSearchText] = useState(urlQuery);
  const [searchField, setSearchField] = useState(urlField);

  // Sync inputs when URL changes (back/forward or navigation)
  useEffect(() => {
    setSearchText(urlQuery);
  }, [urlQuery]);

  useEffect(() => {
    setSearchField(urlField);
  }, [urlField]);

  // Go to page number specified in URL
  function goToPage(newPage: number) {
    const params = new URLSearchParams(searchParams.toString());
    params.set("q", searchText.trim());
    params.set("field", searchField);
    params.set("page", String(newPage));

    startTransition(() => {
      router.push(`/search?${params.toString()}`);
    });
  }

  // Function for initial search. Defaults to page 1
  function initialSearch() {
    if (!searchText.trim()) return;

    const params = new URLSearchParams(searchParams.toString());

    params.set("q", searchText.trim());
    params.set("field", searchField);
    params.set("page", "1");

    startTransition(() => {
      router.push(`/search?${params.toString()}`);
    });
  }

  return (
    <div className="px-4 pt-4 sm:px-6 sm:pt-10">
      {/* Sticky on mobile so users can re-search after scrolling; static on desktop */}
      <div className="sticky top-0 z-20 -mx-4 bg-white/90 px-4 pt-2 pb-2 backdrop-blur sm:static sm:mx-0 sm:bg-transparent sm:px-0 sm:pt-0 sm:pb-0">
        <SearchBar
          searchText={searchText}
          searchField={searchField}
          setSearchText={setSearchText}
          setSearchField={setSearchField}
          onSubmit={initialSearch}
          isPending={isPending}
        />
      </div>
      {/* RESULTS */}
      <div className="relative mt-6 sm:mt-10">
        {/* Loading Overlay: Only visible when isPending is true */}
        {isPending && (
          <div className="absolute inset-0 z-10 flex min-h-[50vh] items-center justify-center bg-white/40 backdrop-blur-[1px]">
            <Spinner className="size-10 text-stone-900 sm:size-15" />
          </div>
        )}

        <div
          className={
            isPending
              ? "pointer-events-none opacity-30 transition-opacity"
              : "opacity-100"
          }
        >
          {searchResult?.length > 0 ? (
            <div className="grid grid-cols-2 gap-2 sm:grid-cols-3 sm:gap-5 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6">
              {searchResult.map((item) => (
                <Link href={`/art/${item.ID}`} key={item.ID} className="group">
                  <div className="relative aspect-3/4 w-full overflow-hidden rounded-lg bg-stone-100">
                    <Image
                      src={item.ImageSmall}
                      alt={item.ArtworkTitle}
                      fill
                      unoptimized
                      sizes="(max-width: 640px) 50vw, (max-width: 768px) 33vw, (max-width: 1024px) 25vw, 20vw"
                      className="object-contain transition-transform group-hover:scale-105"
                    />
                  </div>

                  <h3 className="mt-1.5 line-clamp-2 text-xs font-semibold sm:mt-2 sm:text-sm">
                    {item.ArtworkTitle}
                  </h3>
                  <p className="mt-0.5 text-[11px] text-stone-400 sm:text-xs">
                    {item.Museum}
                  </p>
                </Link>
              ))}
            </div>
          ) : (
            urlQuery && (
              <p className="px-4 py-10 text-center text-sm text-stone-500 sm:py-32 sm:text-base">
                No results found for &quot;{urlQuery}&quot;.
              </p>
            )
          )}
        </div>
      </div>
      {searchResult?.length > 0 && (
        <SearchPagination
          pageNumber={pageNumber}
          goToPage={goToPage}
          maxPageNext={maxPageNext}
        />
      )}
    </div>
  );
}
