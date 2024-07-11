import { Search } from "lucide-react";

export const Searchbar = ({
  value,
  onChange,
}: {
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) => {
  return (
    <div className="relative w-full">
      <Search className="absolute left-2.5 top-2 h-4 w-4 text-muted-foreground" />
      <input
        type="search"
        placeholder="Search..."
        className="w-full rounded-lg bg-black/40 py-1 pl-8 pr-2  border border-white/30 focus:border-white focus:outline-none focus:ring-0 sm:text-sm"
        value={value}
        onChange={onChange}
      />
    </div>
  );
};
