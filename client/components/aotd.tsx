import { Art } from "@/types/search";
import Image from "next/image";
export default async function AOTD() {
  let AOTD: Art | null = null;
  try {
    const apiURL = `${process.env.API_URL}/api/aotd`;

    const res = await fetch(apiURL);
    AOTD = await res.json();
  } catch (err) {
    console.error("Server-side fetch error:", err);
  }
  return (
    <div className="px-10 pt-10">
      <h1 className="mb-2 text-4xl font-bold">{AOTD?.ArtworkTitle}</h1>
      <p className="text-xl text-gray-600">
        {AOTD?.ArtistName || "Unknown Artist"}
      </p>
      <p className="text-lg text-gray-500">{AOTD?.DateCreated}</p>

      <div className="mt-6 space-y-3 border-t pt-6">
        <div>
          <h3 className="text-sm font-semibold text-gray-500 uppercase">
            Medium
          </h3>
          <p className="text-lg">{AOTD?.ArtMedium}</p>
        </div>
        <div>
          <h3 className="text-sm font-semibold text-gray-500 uppercase">
            Collection
          </h3>
          <p className="text-lg">{AOTD?.Museum}</p>
        </div>
      </div>
      <div className="relative aspect-square translate-x-10 overflow-hidden rounded-lg bg-gray-100">
        <Image
          src={AOTD?.ImageLarge || ""}
          fill
          alt={AOTD?.ArtworkTitle || ""}
          className="object-contain"
          unoptimized
        />
      </div>
    </div>
  );
}
