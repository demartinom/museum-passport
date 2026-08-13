import { Artist } from "@/types/artist";

interface ArtistProfileProps {
  artist: Artist;
}

export default function ArtistProfile({ artist }: ArtistProfileProps) {
  return (
    <div>
      <h1>{artist.name}</h1>
    </div>
  );
}
