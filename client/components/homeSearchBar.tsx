"use client";

import { useState, useTransition } from "react";
import { useRouter } from "next/navigation";
import SearchBar from "./searchBar";
import { Spinner } from "@/components/ui/spinner";

interface HomeSearchBarProps {
  children: React.ReactNode;
}

export function HomeSearchBar({ children }: HomeSearchBarProps) {
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

    if (searchField === "artist") {
      startTransition(() => {
        router.push(`/artist?name=${encodeURIComponent(searchText.trim())}`);
      });
      return;
    }

    startTransition(() => {
      router.push(`/search?${params.toString()}`);
    });
  }

  return (
    <div>
      {isPending && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-white/40 backdrop-blur-[1px]">
          <Spinner className="size-24 text-stone-900" />
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

      {!isPending && children}
    </div>
  );
}
