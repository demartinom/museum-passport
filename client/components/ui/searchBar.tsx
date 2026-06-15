import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";

interface SearchBarProps {
  searchText: string;
  searchField: string;
  setSearchText: (value: string) => void;
  setSearchField: (value: string) => void;
  onSubmit: () => void;
  isPending: boolean;
}

export default function SearchBar({
  searchText,
  searchField,
  setSearchText,
  setSearchField,
  onSubmit,
  isPending,
}: SearchBarProps) {
  // Dropdown field options
  const FIELD_OPTIONS = [
    { value: "general", label: "All" },
    { value: "name", label: "Artwork" },
    { value: "artist", label: "Artist" },
  ];

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        onSubmit();
      }}
      className="mx-auto mb-4 max-w-2xl"
    >
      <div className="flex overflow-hidden rounded-xl border border-stone-300 bg-white shadow-sm focus-within:ring-2 focus-within:ring-stone-200">
        {/* FIELD SELECT */}
        <Select
          value={searchField}
          onValueChange={(value) => setSearchField(value)}
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
  );
}
