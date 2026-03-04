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
    <div>
      {data.ArtworkTitle}
      <Image
        src={data.ImageLarge}
        width={1000}
        height={1000}
        alt="art"
        unoptimized
      ></Image>
    </div>
  );
};

export default SingleArtwork;
