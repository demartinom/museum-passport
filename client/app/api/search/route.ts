import { NextRequest, NextResponse } from "next/server";

const baseUrl = process.env.API_URL;

export async function GET(req: NextRequest) {
  const { searchParams } = new URL(req.url);
  const query = searchParams.toString();

  const res = await fetch(`${baseUrl}/api/search?${query}`);
  const data = await res.json();

  return NextResponse.json(data);
}
