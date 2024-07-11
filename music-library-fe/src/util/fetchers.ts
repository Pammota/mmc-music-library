import { Artist, Album, Song, SearchResponse } from "../util/types";

const fetcher = async (url: string) => {
  const res = await fetch(url);
  return res.json();
};

// collections
const fetchArtists = async (): Promise<Artist[]> => {
  const url = "http://localhost:8080/artists";
  return fetcher(url);
};

const fetchAlbums = async (): Promise<Album[]> => {
  const url = "http://localhost:8080/albums";
  return fetcher(url);
};

const fetchSongs = async (): Promise<Song[]> => {
  const url = "http://localhost:8080/songs";
  return fetcher(url);
};

// singles
const fetchArtist = async (id: string): Promise<Artist> => {
  const url = `http://localhost:8080/artists/${id}`;
  return fetcher(url);
};

const fetchAlbum = async (id: string): Promise<Album> => {
  const url = `http://localhost:8080/albums/${id}`;
  return fetcher(url);
};

const fetchSong = async (id: string): Promise<Song> => {
  const url = `http://localhost:8080/songs/${id}`;
  return fetcher(url);
};

// conditionals
const fetchAlbumsByArtist = async (artistId: string): Promise<Album[]> => {
  const url = `http://localhost:8080/albums/byArtist/${artistId}`;
  return fetcher(url);
};

const fetchSongsByAlbum = async (albumId: string): Promise<Song[]> => {
  const url = `http://localhost:8080/songs/byAlbum/${albumId}`;
  return fetcher(url);
};

const fetchSongsByArtist = async (artistId: string): Promise<Song[]> => {
  const url = `http://localhost:8080/songs/byArtist/${artistId}`;
  return fetcher(url);
};

const fetchSearch = async (query: string): Promise<SearchResponse> => {
  const url = `http://localhost:8080/search/${query}`;
  return fetcher(url);
};

const fetchAlbumCoverByTitle = async (
  title: string,
  artistName: string,
  size: "small" | "medium" | "large" = "medium"
): Promise<string> => {
  const url = `http://localhost:4000/${artistName}/${title}/${size}`;
  const res = await fetch(url);
  const { image } = await res.json();
  console.log(image);
  return image;
};

const token = import.meta.env.VITE_SPOTIFY_TOKEN;

const fetchArtistPhotoByName = async (artistName: string): Promise<string> => {
  const url = `https://api.spotify.com/v1/search?type=artist&q=${artistName}&decorate_restrictions=false&best_match=true&include_external=audio&limit=1`;
  console.log(url);
  const res = await fetch(url, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  const artistObj = await res.json();
  return artistObj.best_match.items[0].images[0].url;
};

export {
  // collections
  fetchArtists,
  fetchAlbums,
  fetchSongs,
  // singles
  fetchArtist,
  fetchAlbum,
  fetchSong,
  // conditionals
  fetchAlbumsByArtist,
  fetchSongsByAlbum,
  fetchSongsByArtist,
  // search
  fetchSearch,
  // covers
  fetchAlbumCoverByTitle,
  fetchArtistPhotoByName,
};
