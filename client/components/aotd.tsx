import { Art } from "@/types/search";
import Image from "next/image";
import Link from "next/link";

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
    <div className="grid grid-cols-1 items-center gap-6 px-6 md:grid-cols-2 md:px-12 lg:gap-16 lg:px-24">
      {/* The Context (Left on Desktop, Bottom on Mobile) */}
      <div className="order-last flex flex-col justify-center md:order-first">
        <span className="mb-4 text-xs font-bold tracking-widest text-gray-400 uppercase">
          Artwork of the Day
        </span>

        <Link
          href={`/art/${AOTD?.ID}`}
          className="mb-4 text-4xl leading-tight font-bold text-gray-900 underline md:text-5xl lg:text-5xl"
        >
          {AOTD?.ArtworkTitle}
        </Link>

        <p className="mb-2 text-2xl text-gray-700">
          {AOTD?.ArtistName || "Unknown Artist"}
        </p>

        <p className="mb-8 text-lg text-gray-500 italic">{AOTD?.DateCreated}</p>

        <div className="mt-4 flex flex-col gap-4 border-t pt-8 md:flex-row md:gap-12">
          <div>
            <h3 className="text-xs font-semibold tracking-wider text-gray-400 uppercase">
              Medium
            </h3>
            <p className="mt-1 text-lg text-gray-800">{AOTD?.ArtMedium}</p>
          </div>
          <div>
            <h3 className="text-xs font-semibold tracking-wider text-gray-400 uppercase">
              Collection
            </h3>
            <Link
              href={AOTD?.URL || ""}
              className="text-blue-500 hover:underline"
            >
              <p className="mt-1 text-lg">{AOTD?.Museum}</p>
            </Link>
          </div>
        </div>
      </div>

      {/* The Artwork (Right on Desktop, Top on Mobile) */}
      <div className="relative order-first flex h-[50vh] w-full items-center justify-center overflow-hidden rounded-xl md:order-last md:h-[70vh]">
        <Image
          src={AOTD?.ImageLarge || ""}
          fill
          alt={AOTD?.ArtworkTitle || "Artwork of the Day"}
          className="object-contain p-4"
          unoptimized
        />
      </div>
    </div>
  );
}
