import { useEffect, useState } from "react";
import { Searchbar } from "../SearchBar";
import { AlbumList } from "./AlbumList";
import {
  fetchAlbums,
  fetchArtists,
  fetchSearch,
  fetchSongs,
} from "../../util/fetchers";
import { Album, Artist, Song } from "../../util/types";
import { ArtistList } from "./ArtistList";
import { useDebounce } from "use-debounce";
import { Spinner } from "../Spinner";
import { SongList } from "./SongList";

export const Sidebar = () => {
  const [search, setSearch] = useState<string>("");
  const [debouncedSearch] = useDebounce(search, 1000);
  const [spinner, setSpinner] = useState<boolean>(false);

  const [albums, setAlbums] = useState<Album[]>([]);
  const [artists, setArtists] = useState<Artist[]>([]);
  const [songs, setSongs] = useState<Song[]>([]);

  useEffect(() => {
    fetchAlbums().then((res) => setAlbums(res));
    fetchArtists().then((res) => setArtists(res));
  }, []);

  useEffect(() => {
    if (debouncedSearch !== "") {
      fetchSearch(debouncedSearch).then((res) => {
        setAlbums(res.albums ?? []);
        setArtists(res.artists ?? []);
        setSongs(res.songs ?? []);
      });
    } else {
      fetchAlbums().then((res) => setAlbums(res));
      fetchArtists().then((res) => setArtists(res));
    }
    setSpinner(false);
  }, [debouncedSearch]);

  return (
    <div className="w-[26rem] h-[100vh] bg-gray-100/20 flex flex-col justify-start z-50 rounded-r-xl p-4 pb-0 border-[1.5px] border-black/40">
      <Searchbar
        value={search}
        onChange={(e) => {
          setSearch(e.target.value);
          setSpinner(true);
        }}
      />
      <div className="flex flex-col justify-start gap-10 py-4 pb-0 overflow-y-auto">
        {spinner && (
          <div className="w-full flex justify-center items-center">
            <Spinner />
          </div>
        )}
        <ArtistList artists={artists} />
        <AlbumList albums={albums} />
        {songs && songs.length > 0 && <SongList songs={songs} />}
        <span className="h-2" />
      </div>
    </div>
  );
};
