import React from "react";
import Link from "next/link";
import Image from "next/image";
import { Art } from "@/types/search";

export default function ArtworkCard({ item }: { item: Art }) {
  return (
    <Link href={`/art/${item.ID}`} className="group">
      <div className="relative aspect-3/4 w-full overflow-hidden rounded-lg bg-stone-100">
        <Image
          src={item.ImageSmall}
          alt={item.ArtworkTitle}
          fill
          unoptimized
          sizes="(max-width: 640px) 50vw, (max-width: 768px) 33vw, (max-width: 1024px) 25vw, 20vw"
          className="object-contain transition-transform group-hover:scale-105"
        />
      </div>

      <h3 className="mt-1.5 line-clamp-2 text-xs font-semibold sm:mt-2 sm:text-sm">
        {item.ArtworkTitle}
      </h3>
      <p className="mt-0.5 text-[11px] text-stone-400 sm:text-xs">
        {item.Museum}
      </p>
    </Link>
  );
}
