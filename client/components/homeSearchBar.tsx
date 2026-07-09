"use client";

import { useState, useTransition } from "react";
import { useRouter } from "next/navigation";
import SearchBar from "./searchBar";
import { Spinner } from "@/components/ui/spinner";

export function HomeSearchBar() {
  const router = useRouter();
  const [isPending, startTransition] = useTransition();
  const [searchText, setSearchText] = useState("");
  const [searchField, setSearchField] = useState("general");

  function initialSearch() {
    if (!searchText.trim()) return;

    const params = new URLSearchParams();
    params.set("q", searchText.trim());
    params.set("field", searchField);
    params.set("page", "1");

    startTransition(() => {
      router.push(`/search?${params.toString()}`);
    });
  }

  return (
    <div className="relative">
      {isPending && (
        <div className="absolute inset-0 z-10 flex items-start justify-center bg-white/40 pt-6 backdrop-blur-[1px]">
          <Spinner className="size-8 text-stone-900" />
        </div>
      )}
      <SearchBar
        searchText={searchText}
        searchField={searchField}
        setSearchText={setSearchText}
        setSearchField={setSearchField}
        onSubmit={initialSearch}
        isPending={isPending}
      />
    </div>
  );
}
