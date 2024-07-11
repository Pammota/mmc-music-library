import { useEffect, useState } from "react";
import { useParams } from "wouter";
import { Album, Artist } from "../../util/types";
import {
  fetchAlbumCoverByTitle,
  fetchAlbumsByArtist,
  fetchArtist,
  fetchArtistPhotoByName,
} from "../../util/fetchers";

export const ArtistPage = () => {
  const { id } = useParams();
  const [artist, setArtist] = useState<Artist | null>(null);
  const [artistImageUrl, setArtistImageUrl] = useState<string>("");
  const [albums, setAlbums] = useState<(Album & { photoUrl?: string })[]>([]);

  useEffect(() => {
    if (!id) return;

    fetchArtist(id).then((artist) => {
      setArtist(artist);
      fetchAlbumsByArtist(artist.id).then(async (albums) => {
        if (albums.length === 0) return;

        const albumsWithPhotoUrl: (Album & { photoUrl?: string })[] =
          await Promise.all(
            albums.map(
              async (album) =>
                await fetchAlbumCoverByTitle(
                  album.title,
                  album.artistName
                ).then((imageUrl) => ({ ...album, photoUrl: imageUrl }))
            )
          );

        setAlbums(albumsWithPhotoUrl);
      });

      fetchArtistPhotoByName(artist.name).then((imageUrl) => {
        setArtistImageUrl(imageUrl);
      });
    });
  }, [id]);

  console.log(albums);

  return (
    <div className="flex flex-col items-center justify-center p-4">
      <div className="flex flex-row items-center justify-start w-full max-w-4xl">
        <img
          src={artistImageUrl}
          alt={artist?.name}
          className="w-1/2 min-w-[400px] h-auto rounded-lg shadow-lg"
        />
        <h1 className="text-4xl font-bold ml-4">{artist?.name}</h1>
      </div>
      <div className="w-full max-w-4xl mt-8">
        <h2 className="text-2xl font-semibold mb-4">Albums</h2>
        <ul>
          {albums.map((album) => (
            <li key={album.id} className="mb-2">
              <a href={`/album/${album.id}`}>
                <div className="flex flex-row justify-start items-center p-2 bg-gray-100/20 rounded-lg shadow border-[1.5px] border-black/40">
                  <img
                    src={album.photoUrl}
                    alt={album.title}
                    className="w-[45px] h-[45px] rounded-lg shadow-lg"
                  />
                  <div className="ml-4">
                    <span className="font-medium">{album.title}</span>
                    <span className="text-sm text-gray-600"></span>
                  </div>
                </div>
              </a>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};
