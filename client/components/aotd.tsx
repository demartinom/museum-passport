import { Art } from "@/types/search";
import Image from "next/image";
import Link from "next/link";

export default async function AOTD() {
  let AOTD: Art | null = null;
  try {
    const apiURL = `${process.env.API_URL}/api/aotd`;

    const res = await fetch(apiURL, { cache: "no-store" });
    AOTD = await res.json();
  } catch (err) {
    console.error("Server-side fetch error:", err);
  }

  if (!AOTD) {
    return (
      <div className="px-6 py-12 text-center text-gray-500 md:px-12 lg:px-24">
        Couldn&apos;t load today&apos;s artwork.
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 items-center gap-0 px-4 sm:px-6 md:grid-cols-2 md:gap-10 md:px-12 lg:gap-16 lg:px-24">
      {/* The Context (Left on Desktop, Bottom on Mobile) */}
      <div className="order-last flex flex-col justify-center md:order-first">
        <span className="mb-2 text-xs font-bold tracking-widest text-gray-400 uppercase sm:mb-4">
          Artwork of the Day
        </span>

        <h2 className="mb-3 text-2xl leading-tight font-bold text-gray-900 sm:mb-4 sm:text-4xl md:text-5xl lg:text-5xl">
          <Link href={`/art/${AOTD.ID}`} className="underline">
            {AOTD.ArtworkTitle}
          </Link>
        </h2>

        <p className="mb-1.5 text-lg text-gray-700 sm:mb-2 sm:text-2xl">
          {AOTD.ArtistName || "Unknown Artist"}
        </p>

        <p className="mb-6 text-base text-gray-500 italic sm:mb-8 sm:text-lg">
          {AOTD.DateCreated}
        </p>
      </div>

      {/* The Artwork (Right on Desktop, Top on Mobile) */}
      <div className="relative order-first flex h-[38vh] w-full items-center justify-center overflow-hidden rounded-xl sm:h-[50vh] md:order-last md:h-[70vh]">
        <Image
          src={AOTD.ImageLarge || ""}
          fill
          alt={AOTD.ArtworkTitle || "Artwork of the Day"}
          className="object-contain p-3 sm:p-4"
          unoptimized
        />
      </div>
    </div>
  );
}
