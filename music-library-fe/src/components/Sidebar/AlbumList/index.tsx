import { Album } from "../../../util/types";
import { InfoCard } from "../../InfoCard";

export const AlbumList = ({ albums }: { albums: Album[] }) => {
  return (
    <div className="flex flex-col px-4 pt-4 max-h-[30vh] w-full border-[2px] border-black/30 rounded-xl bg-black/20 shadow-inner shadow-black/30">
      <h2 className="text-2xl font-semibold pb-1">Albums</h2>
      <hr />
      <div className="flex flex-col gap-4 overflow-y-auto w-full">
        <span className="h-4" />
        {albums.map((album) => (
          <a href={`/album/${album.id}`} key={album.id}>
            <InfoCard
              title={album.title}
              subtitle={album.artistName}
              hasImage
              className="cursor-pointer border-[1.5px] border-black/40"
            />
          </a>
        ))}
        {albums.length === 0 && <p className="text-center">No albums found</p>}
        <span className="h-4" />
      </div>
    </div>
  );
};
