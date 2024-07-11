import { Music } from "lucide-react";

export const Homepage = () => {
  return (
    <div className="flex flex-col w-full p-6 justify-center items-center opacity-30 h-[90vh] overflow-hidden">
      <Music className="h-[400px] scale-[10]" />
      <h1 className="text-9xl font-bold">Music Library</h1>
    </div>
  );
};
