import { useEffect, useState } from "react";
import {
  fetchAlbumCoverByTitle,
  fetchArtistPhotoByName,
} from "../../util/fetchers";

type Props = {
  title: string;
  subtitle?: string;
  hasImage?: boolean;
  className?: string;
};

export const InfoCard = ({ title, subtitle, hasImage, className }: Props) => {
  const [imageUrl, setImageUrl] = useState<string>("");

  useEffect(() => {
    if (hasImage && subtitle) {
      fetchAlbumCoverByTitle(title, subtitle).then((url: string) => {
        setImageUrl(url);
      });
    }

    if (hasImage && !subtitle) {
      fetchArtistPhotoByName(title).then((url: string) => {
        setImageUrl(url);
      });
    }
  }, [hasImage, subtitle, title]);

  return (
    <div
      className={
        "flex justify-start rounded-xl p-2.5 bg-white/20 shadow-md gap-4 max-h-[5.6rem] w-full" +
          className ?? ""
      }
    >
      {hasImage && (
        <img src={imageUrl} alt={title} className="w-16 h-16 rounded-xl" />
      )}

      <div
        className={`flex flex-col ${
          !hasImage && !subtitle ? "items-center" : ""
        } ${subtitle ? "justify-between" : "justify-center"}`}
      >
        <h3 className="text-lg font-bold">{title}</h3>
        {subtitle && <p className="text-sm">{subtitle}</p>}
      </div>
    </div>
  );
};
