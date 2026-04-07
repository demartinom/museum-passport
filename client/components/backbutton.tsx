"use client";

import { useRouter } from "next/navigation";
import { ArrowLeftIcon } from "lucide-react";
export default function BackButton() {
  const router = useRouter();

  return (
    <button
      onClick={() => router.back()}
      className="absolute top-6 left-6 flex cursor-pointer items-center gap-2 text-gray-500 transition-colors hover:text-black"
    >
      <ArrowLeftIcon size={24} />
      <span className="text-sm font-medium text-black">Back</span>
    </button>
  );
}
