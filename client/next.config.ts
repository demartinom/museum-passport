import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "images.metmuseum.org",
      },
      {
        protocol: "https",
        hostname: "ids.lib.harvard.edu", // Harvard images
      },
      {
        protocol: "https",
        hostname: "nrs.harvard.edu", // Also Harvard
      },
      {
        protocol: "https",
        hostname: "**.harvard.edu", // Wildcard for all Harvard subdomains
      },
    ],
  },
};

export default nextConfig;
