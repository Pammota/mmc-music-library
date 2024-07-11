import { Route, Switch } from "wouter";
import { Homepage } from "./pages/Homepage";
import { Sidebar } from "./components/Sidebar";
import { ArtistPage } from "./pages/ArtistPage";
import { AlbumPage } from "./pages/AlbumPage";

export default function App() {
  return (
    <main className="h-screen flex justify-between">
      <Sidebar />
      <div className="flex flex-col w-full p-4 overflow-hidden">
        <div className="overflow-y-auto">
          <Switch>
            <Route path="/" component={Homepage} />
            <Route path="/artist/:id" component={ArtistPage} />
            <Route path="/album/:id" component={AlbumPage} />
            <Route>404: No such page!</Route>
          </Switch>
        </div>
      </div>
    </main>
  );
}
