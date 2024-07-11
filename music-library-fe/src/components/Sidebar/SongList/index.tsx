import { Song } from "../../../util/types";
import { InfoCard } from "../../InfoCard";

export const SongList = ({ songs }: { songs: Song[] }) => {
  return (
    <div className="flex flex-col px-4 pt-4 max-h-[25vh] w-full border-[2px] border-black/30 rounded-xl bg-black/20 shadow-inner shadow-black/30">
      <h2 className="text-2xl font-semibold pb-1">Songs</h2>
      <hr />
      <div className="flex flex-col gap-4 overflow-y-auto w-full">
        <span className="h-4" />
        {songs.map((song) => (
          <InfoCard
            key={song.id}
            title={song.title}
            subtitle={song.albumTitle}
            hasImage
            className="cursor-pointer max-h-[4rem] border-[1.5px] border-black/40"
          />
        ))}
        {songs.length === 0 && <p className="text-center">No songs found</p>}
        <span className="h-4" />
      </div>
    </div>
  );
};
