import { Artist } from "@/types/artist";
import Image from "next/image";
import { fraunces } from "../lib/fonts";
interface ArtistProfileProps {
  artist: Artist;
}

export default function ArtistProfile({ artist }: ArtistProfileProps) {
  return (
    <div className="min-h-screen text-[#1A1A1A]">
      <div className="mx-auto max-w-6xl px-6 py-16 md:py-24">
        <div className="grid grid-cols-1 items-center gap-10 md:grid-cols-[1fr_55%] md:gap-16">
          <div className="order-2 md:order-1 md:py-8">
            <h1
              className={`${fraunces.className} text-5xl leading-[1.05] tracking-tight md:text-6xl`}
            >
              {artist.name}
            </h1>

            <div className="mt-5 border-t border-[#DEDAD2] pt-5">
              <h2 className="font-sans text-lg text-[#3A4F3F] italic md:text-xl">
                {artist.description}
              </h2>
            </div>

            <p className="mt-6 max-w-[62ch] font-sans text-base leading-relaxed text-[#1A1A1A]/85">
              {artist.bio}
            </p>
          </div>

          {/* Artist Image */}
          <div className="relative order-1 aspect-4/5 w-full bg-[#DEDAD2] md:order-2">
            <Image
              src={artist.imageUrl}
              fill
              className="object-cover"
              alt={artist.name}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
