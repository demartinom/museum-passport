import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
  const { searchParams } = new URL(req.url);
  const query = searchParams.toString();

  const res = await fetch(`${process.env.API_URL}/api/search?${query}`);
  const data = await res.json();

  return NextResponse.json(data);
}
