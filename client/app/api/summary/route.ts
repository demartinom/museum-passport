import { NextRequest, NextResponse } from "next/server";

const baseUrl = process.env.API_URL;

export async function GET(req: NextRequest) {
  const id = new URL(req.url).searchParams.get("id");

  const res = await fetch(`${baseUrl}/api/summary?id=${id}`);
  const data = await res.json();

  return NextResponse.json(data);
}
