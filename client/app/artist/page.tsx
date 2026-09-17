import ArtistProfile from "@/components/artistProfile";
import ArtworkCard from "@/components/artworkCard";
import { Artist } from "@/types/artist";
import { SearchResult } from "@/types/search";

export default async function ArtistPage({
  searchParams,
}: {
  searchParams: Promise<{ name?: string }>;
}) {
  const { name } = await searchParams;
  return (
    <div>
      <ArtistInfo name={name} />
      <ArtistWorks name={name} />;
    </div>
  );
}

async function ArtistInfo({ name }: { name?: string }) {
  if (!name) return <p className="py-32 text-center">No artist specified.</p>;
  const res = await fetch(
    `${process.env.API_URL}/api/artist?artistname=${encodeURIComponent(name)}`,
    { cache: "no-store" },
  );

  if (!res.ok) {
    return (
      <p className="px-4 py-32 text-center text-stone-500">
        No results found for &quot;{name}&quot;.
      </p>
    );
  }

  const artist: Artist = await res.json();
  return <ArtistProfile artist={artist} />;
}

async function ArtistWorks({ name }: { name?: string }) {
  const res = await fetch(`${process.env.API_URL}/api/search?artist=${name}`);
  if (!res.ok) {
    return (
      <p className="px-4 py-32 text-center text-stone-500">
        No artworks found for &quot;{name}&quot;.
      </p>
    );
  }
  const works: SearchResult = await res.json();
  const artworks = works.results
    .slice(0, 5)
    .map((artwork) => (
      <ArtworkCard key={artwork.ID} artwork={artwork} artistPage={true} />
    ));
  const artistNameFormatted = works.results[0].ArtistName.split(",");
  return (
    <section className="w-full px-4 py-8 sm:px-6 lg:px-8">
      <h2 className="mb-4 text-lg font-semibold text-stone-800 sm:text-xl">
        Works by {artistNameFormatted[0]}
      </h2>
      <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 sm:gap-6 md:grid-cols-5">
        {artworks}
      </div>
    </section>
  );
}
