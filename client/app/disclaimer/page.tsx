import BackButton from "@/components/backbutton";
import { fraunces } from "@/lib/fonts";
import Link from "next/link";
const Disclaimer = () => {
  return (
    <div className="relative mt-10">
      <BackButton />
      <h1 className={`${fraunces.className} text-center text-7xl`}>
        AI Disclaimer
      </h1>
      <p className="mt-10 p-10 text-center text-4xl">
        The content within the art description has been generated using{" "}
        <Link
          className="underline"
          target="_blank"
          href={
            "https://openai.com/index/gpt-4o-mini-advancing-cost-efficient-intelligence/"
          }
        >
          GPT‑4o mini
        </Link>
        . 100% accuracy of information can not be guaranteed. It is recommended
        to follow the link to the art page on the museum&#39;s website.
      </p>
    </div>
  );
};

export default Disclaimer;
