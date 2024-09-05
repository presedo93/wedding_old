import type { MetaFunction } from "@remix-run/node";

import { Cover } from "./cover";
// import { SpotifyList } from './spotify-list'

export const meta: MetaFunction = () => {
  return [
    { title: "Laura & Rene" },
    { name: "description", content: "Our wedding!" },
  ];
};

export default function Index() {
  return (
    <div className="flex flex-col items-center">
      <Cover />
      <div className="h-12" />
      {/* <SpotifyList /> */}
      {/* <div className="h-[1000px]" /> */}
    </div>
  );
}
