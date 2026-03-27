"use client";
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
import Image from "next/image";
import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import { Fraunces } from "next/font/google";

const fraunces = Fraunces({
  weight: ["400", "700"],
});

const FIELD_OPTIONS = [
  {
    value: "general",
    label: "All",
    placeholder: "Search artworks, artists, mediums...",
  },
  {
    value: "name",
    label: "Artwork",
    placeholder: "e.g. Starry Night, The Thinker...",
  },
  {
    value: "artist",
    label: "Artist",
    placeholder: "e.g. Monet, Frida Kahlo...",
  },
];

// Search by Medium. Currently disabled
// const MEDIUM_OPTIONS = [
//   { label: "Painting", value: "painting" },
//   { label: "Drawing", value: "drawing" },
//   { label: "Print", value: "print" },
//   { label: "Photography", value: "photography" },
//   { label: "Sculpture", value: "sculpture" },
//   { label: "Ceramic", value: "ceramic" },
//   { label: "Textile", value: "textile" },
//   { label: "Jewelry", value: "jewelry" },
//   { label: "Furniture", value: "furniture" },
//   { label: "Glass", value: "glass" },
//   { label: "Manuscript", value: "manuscript" },
//   { label: "Mixed Media", value: "mixed media" },
// ];

export default function SearchPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const urlField = searchParams.get("field");
  const urlQuery = searchParams.get("q");

  const isMediumSearch = urlField === "type";
  const [field, setField] = useState(
    isMediumSearch ? "type" : urlField || "general",
  );
  const [searchText, setSearchText] = useState(
    isMediumSearch ? "" : urlQuery || "",
  );

  // Search by Medium. Currently disabled
  //   isMediumSearch ? urlQuery || "" : "",
  // );
  const [results, setResults] = useState<Array<Art>>([]);
  const [searching, setSearching] = useState(false);

  // Checks if there is an issue with searching
  const isMediumMode = field === "type";
  const activePlaceholder =
    FIELD_OPTIONS.find((o) => o.value === field)?.placeholder ?? "Search...";

  useEffect(() => {
    if (!urlField || !urlQuery) return;

    async function fetchResults() {
      setSearching(true);

      setResults([]);

      try {
        const res = await fetch(
          `http://localhost:3001/api/search?${urlField}=${encodeURIComponent(urlQuery!)}&length=80`,
        );
        const data = await res.json();
        setResults(data);
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

  // Search by Medium. Currently disabled
  // function handleMediumSelect(value: string) {
  //   const next = value === selectedMedium ? "" : value;
  //   setSelectedMedium(next);
  //   if (next) {
  //     router.push(`/search?field=type&q=${encodeURIComponent(next)}`);
  //   }
  // }

  function handleFieldChange(value: string) {
    setField(value);
    // setSelectedMedium("");
    setSearchText("");
  }

  const activeFieldLabel = isMediumSearch
    ? "Medium"
    : (FIELD_OPTIONS.find((o) => o.value === urlField)?.label ?? urlField);

  return (
    <div className="min-h-screen bg-stone-50 px-6 py-10">
      {/* Search bar */}
      <form onSubmit={handleSearch} className="mx-auto mb-4 max-w-2xl">
        <div
          className={`flex overflow-hidden rounded-xl border border-stone-300 bg-white shadow-sm transition-all focus-within:border-stone-500 focus-within:ring-2 focus-within:ring-stone-200 ${isMediumMode ? "opacity-100" : ""}`}
        >
          <Select value={field} onValueChange={handleFieldChange}>
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
                {/* <SelectItem value="type">Medium</SelectItem> */}
              </SelectGroup>
            </SelectContent>
          </Select>
    <div>
      <div className="flex justify-center">
        <h1 className={`my-6 ${fraunces.className} text-7xl font-bold`}>
          Museum Passport
        </h1>

          {!isMediumMode ? (
            <>
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
                className="shrink-0 bg-stone-900 px-5 text-sm font-medium text-white transition-colors hover:bg-stone-700 disabled:cursor-not-allowed disabled:opacity-40"
              >
                Search
              </button>
            </>
          ) : (
            <div className="flex flex-1 items-center px-4 text-sm text-stone-400 select-none">
              Select a medium below
            </div>
          )}
        </div>
        {/* Search by Medium. Currently disabled */}
        {/* {isMediumMode && (
          <div className="mt-3 flex flex-wrap gap-2">
            {MEDIUM_OPTIONS.map((opt) => {
              const active = selectedMedium === opt.value;
              return (
                <button
                  key={opt.value}
                  type="button"
                  onClick={() => handleMediumSelect(opt.value)}
                  className={`rounded-full border px-3 py-1 text-sm font-medium transition-colors ${
                    active
                      ? "border-stone-900 bg-stone-900 text-white"
                      : "border-stone-300 bg-white text-stone-600 hover:border-stone-500 hover:text-stone-900"
                  }`}
                >
                  {opt.label}
                </button>
              );
            })}
          </div>
        )} */}

        {urlQuery && (
          <p className="mt-3 text-xs text-stone-500">
            Showing results for{" "}
            <span className="font-semibold text-stone-700 capitalize">
              {urlQuery}
            </span>
            {urlField && urlField !== "general" && (
              <>
                {" "}
                in{" "}
                <span className="font-semibold text-stone-700">
                  {activeFieldLabel}
                </span>
              </>
            )}
          </p>
        )}
      </form>

      {searching && (
        <div className="flex flex-col items-center justify-center gap-4 py-32 text-stone-500">
          <Spinner className="size-10" />
          <p className="text-sm">Searching the collection…</p>
        </div>
      )}

      {/* No results */}
      {!searching && searchText && results.length === 0 && (
        <div className="flex flex-col items-center justify-center gap-2 py-32 text-stone-400">
          <p className="text-lg font-medium text-stone-600">No results found</p>
          <p className="text-sm">Try a different term or search field.</p>
        </div>
      )}

      {/* Initial instruction */}
      {!searching && !searchText && results.length === 0 && (
        <div className="flex flex-col items-center justify-center gap-2 py-32 text-stone-400">
          <p className="text-sm">
            Enter a search above to explore the collection.
          </p>
        </div>
      )}

      {!searching && results.length > 0 && (
        <>
          <p className="mb-4 text-xs text-stone-400">
            {results.length} results
          </p>
          <div className="grid grid-cols-2 gap-5 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
            {results.map((item) => (
              <a href={`/art/${item.ID}`} key={item.ID} className="group">
                <div className="space-y-2">
                  <div className="relative h-56 w-full overflow-hidden rounded-lg bg-stone-100">
                    <Image
                      src={item.ImageSmall}
                      alt={item.ArtworkTitle}
                      fill
                      className="object-contain transition-transform duration-300 group-hover:scale-105"
                    />
                  </div>
                  <div>
                    <h3 className="truncate text-sm font-semibold text-stone-800">
                      {item.ArtworkTitle || "Untitled"}
                    </h3>
                    <p className="truncate text-xs text-stone-500">
                      {item.Museum}
                    </p>
                  </div>
                </div>
              </a>
            ))}
          </div>
        </>
      )}
    </div>
  );
}
