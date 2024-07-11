import { useEffect, useState } from "react";
import { useParams } from "wouter";
import { Album, Song } from "../../util/types";
import {
  fetchAlbumCoverByTitle,
  fetchAlbum,
  fetchSongsByAlbum,
} from "../../util/fetchers";

export const AlbumPage = () => {
  const { id } = useParams();
  const [album, setAlbum] = useState<Album | null>(null);
  const [albumImageUrl, setAlbumImageUrl] = useState<string>("");
  const [songs, setSongs] = useState<Song[]>([]);

  useEffect(() => {
    if (!id) return;

    fetchAlbum(id).then((Album) => {
      setAlbum(Album);
      fetchAlbumCoverByTitle(Album.title, Album.artistName).then((imageUrl) => {
        setAlbumImageUrl(imageUrl);
      });
      fetchSongsByAlbum(id).then((songs) => {
        setSongs(songs);
      });
    });
  }, [id]);

  return (
    <div className="flex flex-col items-center justify-center p-4">
      <div className="flex flex-row items-center justify-start w-full max-w-4xl">
        <img
          src={albumImageUrl}
          alt={album?.title}
          className="w-1/3 min-w-[400px] h-auto rounded-lg shadow-lg"
        />
        <div className="flex flex-col ml-4 justify-center items-center">
          <h1 className="text-4xl font-bold ml-4">{album?.title}</h1>
          <p className="text-2xl">{album?.artistName}</p>
        </div>
      </div>
      <div className="flex flex-col justify-start items-start max-w-4xl pt-10">
        <div className="flex flex-col">
          <h2 className="text-2xl font-semibold mb-4">Description</h2>
          <p className="">{album?.description}</p>
        </div>
      </div>
      <div className="w-full max-w-4xl mt-8">
        <h2 className="text-2xl font-semibold mb-4">Songs</h2>
        <ul>
          {songs.map((song) => (
            <li key={song.id} className="mb-2">
              <div className="flex flex-row justify-start items-center p-2 bg-gray-100/20 rounded-lg shadow px-4 border-[1.5px] border-black/40">
                <div className="flex justify-between w-full">
                  <span className="font-medium">{song.title}</span>
                  <span className=" opacity-80"> {song.length}</span>
                </div>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};
