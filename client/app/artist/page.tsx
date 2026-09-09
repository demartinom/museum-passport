import ArtistProfile from "@/components/ArtistProfile";
import { Artist } from "@/types/artist";

export default async function ArtistPage({
  searchParams,
}: {
  searchParams: Promise<{ name?: string }>;
}) {
  const { name } = await searchParams;
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
