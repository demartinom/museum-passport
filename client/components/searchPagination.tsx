import {
  Pagination,
  PaginationItem,
  PaginationNext,
  PaginationPrevious,
  PaginationContent,
} from "./pagination";

interface PaginationProps {
  pageNumber: number;
  goToPage: (value: number) => void;
  maxPageNext: number;
}

export default function SearchPagination({
  pageNumber,
  goToPage,
  maxPageNext,
}: PaginationProps) {
  return (
    <Pagination className="mt-10">
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious
            href="#"
            onClick={(e) => {
              e.preventDefault();
              if (pageNumber > 1) {
                goToPage(pageNumber - 1);
              }
            }}
            className={pageNumber <= 1 ? "pointer-events-none opacity-50" : ""}
          />
        </PaginationItem>
        <PaginationItem>
          <PaginationNext
            href="#"
            onClick={(e) => {
              e.preventDefault();
              if (pageNumber < maxPageNext) {
                goToPage(pageNumber + 1);
              }
            }}
            className={
              pageNumber >= maxPageNext ? "pointer-events-none opacity-50" : ""
            }
          />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
