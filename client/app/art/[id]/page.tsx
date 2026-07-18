import { AISummary } from "@/components/aisummary";
import { Art } from "@/types/search";
import Image from "next/image";
import BackButton from "@/components/backbutton";
import Link from "next/link";

const SingleArtwork = async ({
  params,
}: {
  params: Promise<{ id: string }>;
}) => {
  const { id } = await params;

  let data: Art;
  try {
    const res = await fetch(`${process.env.API_URL}/api/artwork/${id}`);

    data = await res.json();
  } catch (err) {
    console.error(err);
    return <div>Something went wrong</div>;
  }

  return (
    <div className="relative flex justify-center">
      <BackButton />
      <div className="container px-4 py-8 sm:py-12">
        <div className="grid grid-cols-1 items-start gap-8 lg:grid-cols-[3fr_4fr] lg:gap-0">
          {/* Image comes first on mobile via order-1; text follows via order-2. On lg screens order is reset so the grid-column placement (text left, image right) takes over. */}
          <div className="relative order-1 aspect-square w-full overflow-hidden rounded-lg bg-gray-100 lg:order-2 lg:translate-x-10">
            <Image
              src={data.ImageLarge}
              fill
              alt={data.ArtworkTitle}
              className="object-contain"
              unoptimized
            />
          </div>

          <div className="order-2 px-4 sm:px-6 lg:order-1 lg:px-10 lg:pt-10">
            <h1 className="mb-2 text-center text-2xl font-bold sm:text-3xl lg:text-4xl">
              {data.ArtworkTitle}
            </h1>
            <p className="text-center text-lg text-gray-600 sm:text-xl">
              {data.ArtistName || "Unknown Artist"}
            </p>
            <p className="text-base text-gray-500 sm:text-lg">
              {data.DateCreated}
            </p>

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
                <Link
                  href={data.URL}
                  target="_blank"
                  className="text-blue-500 hover:underline"
                >
                  <p className="text-lg">{data.Museum}</p>
                </Link>
              </div>
              <AISummary id={id} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SingleArtwork;
