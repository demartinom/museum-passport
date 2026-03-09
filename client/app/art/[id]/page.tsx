import { Art } from "@/types/search";
import Image from "next/image";

const SingleArtwork = async ({
  params,
}: {
  params: Promise<{ id: string }>;
}) => {
  const { id } = await params;

  let data: Art;
  try {
    const res = await fetch(`http://localhost:3001/api/artwork/${id}`);

    data = await res.json();
    console.log(data);
  } catch (err) {
    console.error(err);
    return <div>Something went wrong</div>;
  }

  return (
    <div className="container mx-auto max-w-7xl px-4 py-12">
      <div className="grid grid-cols-1 items-start gap-12 lg:grid-cols-3">
        <div className="pt-24">
          <h1 className="mb-2 text-4xl font-bold">{data.ArtworkTitle}</h1>
          <p className="text-xl text-gray-600">
            {data.ArtistName || "Unknown Artist"}
          </p>
          <p className="text-lg text-gray-500">{data.DateCreated}</p>

          <div className="mt-6 space-y-3 border-t pt-6">
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">
                Medium
              </h3>
              <p className="text-lg">{data.ArtMedium}</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">
                Collection
              </h3>
              <p className="text-lg">{data.Museum}</p>
            </div>
          </div>
        </div>

        <div className="relative aspect-square translate-x-10 overflow-hidden rounded-lg bg-gray-100 lg:col-span-2">
          <Image
            src={data.ImageLarge}
            fill
            alt={data.ArtworkTitle}
            className="object-contain"
            unoptimized
          />
        </div>
      </div>
    </div>
  );
};

export default SingleArtwork;
