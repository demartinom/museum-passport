"use client";

import { useState, useTransition } from "react";
import { useRouter } from "next/navigation";
import SearchBar from "./searchBar";

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
    <SearchBar
      searchText={searchText}
      searchField={searchField}
      setSearchText={setSearchText}
      setSearchField={setSearchField}
      onSubmit={initialSearch}
      isPending={isPending}
    />
  );
}
