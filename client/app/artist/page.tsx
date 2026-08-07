// app/artist/page.tsx
import { Artist } from "@/types/artist";
import { redirect } from "next/navigation";

export default async function SearchArtist({
  searchParams,
}: {
  searchParams: Promise<{ name?: string }>;
}) {
  const params = await searchParams;
  const name = params.name || "";
  if (!name) return <p className="py-32 text-center">No artist specified.</p>;

  let artist: Artist;

  try {
    const res = await fetch(
      `${process.env.API_BASE_URL}/api/artist?artistname=${encodeURIComponent(name)}`,
      { cache: "no-store" },
    );
    artist = await res.json();
  } catch {
    return (
      <p className="px-4 py-32 text-center text-stone-500">
        No results found for &quot;{name}&quot;.
      </p>
    );
  }

  redirect(`/artist/${artist.id}`); // server-side redirect, no client fetch needed
}
