"use client";

import { useState, useEffect } from "react";

export const AISummary = ({ id }: { id: string }) => {
  const [summary, setSummary] = useState<string>("");
  const [loadingSummary, setLoadingSummary] = useState(true);

  useEffect(() => {
    async function fetchSummary() {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URL}/summary?id=${id}`,
        );
        const data = await res.json();
        setSummary(data.summary);
      } catch (err) {
        console.error(err);
      } finally {
        setLoadingSummary(false);
      }
    }

    fetchSummary();
  }, [id]);

  return (
    <div>
      {/* Artwork display */}

      <div className="mt-6 rounded-lg bg-gray-50 p-2">
        <h3 className="mb-2 font-semibold">About this artwork</h3>
        {loadingSummary ? (
          <p className="text-gray-500">Generating summary...</p>
        ) : (
          <p className="whitespace-pre-line text-gray-700">{summary}</p>
        )}
      </div>
    </div>
  );
};
