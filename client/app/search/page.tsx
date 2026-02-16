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
      <form className="w-1/4">
        {/* Combobox instead? */}
        <div className="flex">
          <Select>
            <SelectTrigger>
              <SelectValue placeholder="Field" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="art">Artwork</SelectItem>
                <SelectItem value="artist">Artist</SelectItem>
                <SelectItem value="medium">Medium</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>

          <Input type="text" id="input-search"></Input>
        </div>
      </form>
    </div>
  );
};
export default Search;
