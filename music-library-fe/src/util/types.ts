type Artist = {
  id: string;
  name: string;
};

type Album = {
  id: string;
  title: string;
  description: string;
  // artistId: string;
  artistName: string;
};

type Song = {
  id: string;
  title: string;
  length: string;
  // albumId: string;
  artistName: string;
  albumTitle: string;
};

type SearchResponse = {
  artists?: Artist[];
  albums?: Album[];
  songs?: Song[];
};

export type { Artist, Album, Song, SearchResponse };
