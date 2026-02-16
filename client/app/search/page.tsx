import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const Search = () => {
  async function handleSearch(e: React.MouseEvent<HTMLButtonElement>) {
    e.preventDefault();

    const res = await fetch(
      `http://localhost:3001/api/search?museum=met&${field}=${searchText}&length=80`,
    );

    const data = await res.json();
    setResults(data);
  }
  return (
    <div>
      <form className="w-1/2">
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
          <Button type="submit" onClick={handleSearch}>
            Search
          </Button>
        </div>
      </form>
      <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
        {results?.Art.map((item) => (
          <div
            key={item.ID}
            className="relative aspect-square overflow-hidden rounded-lg"
          >
            <Image
              src={item.ImageSmall}
              alt={item.ArtworkTitle}
              fill
              className="object-contain"
            />
          </div>
        ))}
      </div>
    </div>
  );
};
export default Search;
