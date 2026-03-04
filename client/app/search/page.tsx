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
import { SearchResult } from "@/types/search";
import Image from "next/image";
import { useState } from "react";

const Search = () => {
  const [field, setField] = useState("");
  const [searchText, setSearchText] = useState("");
  const [results, setResults] = useState<SearchResult>();
  const [searching, setSearching] = useState(false);

  async function handleSearch(e: React.SubmitEvent<HTMLFormElement>) {
    e.preventDefault();
    setSearching(true);
    setResults(undefined);

    try {
      const res = await fetch(
        `http://localhost:3001/api/search?museum=met&${field}=${encodeURIComponent(searchText)}&length=80`,
      );
      const data = await res.json();
      setResults(data);
    } catch (err) {
      console.error(err);
    } finally {
      setSearching(false);
    }
  }

  return (
    <div>
      <form className="w-1/2" onSubmit={handleSearch}>
        {/* Combobox instead? */}
        {/* Also Field? */}
        <div className="flex">
          <Select
            value={field}
            onValueChange={(value) => {
              setField(value);
            }}
          >
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
            onChange={(e) => {
              setSearchText(e.target.value);
            }}
          ></Input>
        </div>
      </form>
      {searching == true && (
        <div className="flex flex-col min-h-screen justify-center items-center -translate-y-25">
          <h1 className="">Searching</h1>
          <Spinner className="size-40" />
        </div>
      )}
      <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4 px-5 py-10">
        {results?.Art.map((item) => (
          <div key={item.ID} className="space-y-2">
            <div className="h-64 w-full relative overflow-hidden ">
              <Image
                src={item.ImageSmall}
                alt={item.ArtworkTitle}
                fill
                className="object-contain"
              />
            </div>
            <h3 className="font-semibold text-sm">
              {item.ArtworkTitle || "Untitled"}
            </h3>
            <p className="text-xs text-gray-600">
              {item.ArtistName || "Artist Unknown"}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};
export default Search;
