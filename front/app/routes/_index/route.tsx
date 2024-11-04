import { json, LoaderFunction, type MetaFunction } from "@remix-run/node";
import { User, authenticator } from "~/lib/auth.server";

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
  readonly auth: User | null;
};

export const loader: LoaderFunction = async ({ request }) => {
  const auth = await authenticator.isAuthenticated(request);

  return json<LoaderResponse>({ auth });
};

export default function Index() {
  const data = useLoaderData<LoaderResponse>();

  return (
    <div className="flex flex-col items-center">
      <NavBar isAuth={data.auth !== null} />
      <Cover />
      <div className="h-12" />
      <p>{data.auth?.email}</p>
      {/* <SpotifyList /> */}
      <div className="h-[1000px]" />
    </div>
  );
}
