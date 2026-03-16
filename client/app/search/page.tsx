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

export default function SearchPage() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const urlField = searchParams.get("field");
  const urlQuery = searchParams.get("q");

  const [field, setField] = useState(urlField || "general");
  const [searchText, setSearchText] = useState(urlQuery || "");
  const [results, setResults] = useState<Array<Art>>([]);
  const [searching, setSearching] = useState(false);

  useEffect(() => {
    if (!urlField || !urlQuery) return;

    async function fetchResults() {
      setSearching(true);
      setResults([]);

      try {
        const res = await fetch(
          `http://localhost:3001/api/search?museum=met&${urlField}=${encodeURIComponent(urlQuery)}&length=80`,
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
  }, [field, urlField, urlQuery]);

  function handleSearch(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault();
    if (!field || !searchText) return;

    router.push(`/search?field=${field}&q=${encodeURIComponent(searchText)}`);
  }

  return (
    <div>
      <form className="w-1/2" onSubmit={handleSearch}>
        <div className="flex">
          <Select value={field} onValueChange={setField}>
            <SelectTrigger>
              <SelectValue placeholder="Field" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="name">Artwork</SelectItem>
                <SelectItem value="artist">Artist</SelectItem>
                <SelectItem value="type">Medium</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>

          <Input
            type="text"
            id="input-search"
            value={searchText}
            onChange={(e) => setSearchText(e.target.value)}
          />
        </div>
      </form>

      {searching == true && (
        <div className="flex min-h-screen -translate-y-25 flex-col items-center justify-center">
          <h1 className="">Searching</h1>
          <Spinner className="size-40" />
        </div>
      )}

      <div className="grid grid-cols-2 gap-4 px-5 py-10 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
        {results?.Art.map((item) => (
          <a
            href={`/art/${item.ID}?back=/search?field=${urlField}&q=${urlQuery}`}
            key={item.ID}
          >
            <div className="space-y-2">
              <div className="relative h-64 w-full overflow-hidden">
                <Image
                  src={item.ImageSmall}
                  alt={item.ArtworkTitle}
                  fill
                  className="object-contain"
                />
              </div>
              <h3 className="text-sm font-semibold">
                {item.ArtworkTitle || "Untitled"}
              </h3>
              <p className="text-xs text-gray-600">
                {item.ArtistName || "Artist Unknown"}
              </p>
            </div>
          </a>
        ))}
      </div>
    </div>
  );
}
