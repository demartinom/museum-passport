export interface SearchResult {
  ResultsLength: number;
  Art: Art[];
}

export interface Art {
  id: string;
  ArtMedium: string;
  ArtistName: string;
  ArtworkTitle: string;
  ArtworkType: string;
  DateCreated: string;
  ID: string;
  ImageLarge: string;
  ImageSmall: string;
  Museum: string;
  PublicDomain: boolean;
}
