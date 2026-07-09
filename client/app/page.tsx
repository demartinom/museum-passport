import { fraunces } from "../lib/fonts";
import AOTD from "@/components/aotd";
import { HomeSearchBar } from "@/components/homeSearchBar";

export default function HomePage() {
  return (
    <div>
      <h1
        className={`my-6 text-center ${fraunces.className} text-7xl font-bold text-stone-900`}
      >
        Museum Passport
      </h1>

      <div className="px-6">
        <HomeSearchBar>
          <AOTD />
        </HomeSearchBar>
      </div>
    </div>
  );
}
