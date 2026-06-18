"use client";

import { useState, useEffect, useTransition } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import { Spinner } from "@/components/ui/spinner";
import { Art } from "@/types/search";
import SearchBar from "./ui/searchBar";
import SearchPagination from "./ui/searchPagination";

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
      router.push(`/?${params.toString()}`);
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
      router.push(`/?${params.toString()}`);
    });
  }

  return (
    <div className="min-h-screen px-6 py-10">
      <SearchBar
        searchText={searchText}
        searchField={searchField}
        setSearchText={setSearchText}
        setSearchField={setSearchField}
        onSubmit={initialSearch}
        isPending={isPending}
      />
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
          {searchResult?.length > 0 ? (
            <div className="grid grid-cols-2 gap-5 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
              {searchResult.map((item) => (
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
                No results found for `&quot;`{urlQuery}`&quot;`.
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
