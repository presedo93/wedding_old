import { json, LoaderFunction, type MetaFunction } from "@remix-run/node";
import { authenticator } from "~/lib/auth.server";

import { Cover } from "./cover";
import { useLoaderData } from "@remix-run/react";
import { NavBar } from "./nav-bar";
// import { SpotifyList } from './spotify-list'

export const meta: MetaFunction = () => {
  return [
    { title: "Laura & Rene" },
    { name: "description", content: "Our wedding!" },
  ];
};

type LoaderResponse = {
  readonly mail?: string;
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request);
  return json<LoaderResponse>({ mail: user?.email });
};

export default function Index() {
  const data = useLoaderData<LoaderResponse>();

  return (
    <div className="flex flex-col items-center">
      <NavBar isAuth={data.mail !== undefined} />
      <Cover />
      <div className="h-12" />
      <p>{data?.mail}</p>
      {/* <SpotifyList /> */}
      <div className="h-[1000px]" />
    </div>
  );
}
