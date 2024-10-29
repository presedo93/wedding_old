import type { LogtoContext } from "@logto/remix";
import {
  json,
  LoaderFunction,
  redirect,
  type MetaFunction,
} from "@remix-run/node";
import { logto } from "~/lib/auth.server";

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
  readonly context: LogtoContext;
};

export const loader: LoaderFunction = async ({ request }) => {
  const context = await logto.getContext({ getAccessToken: true })(request);

  if (!context.isAuthenticated) {
    return redirect("/auth/sign-in");
  }

  return json<LoaderResponse>({ context });
};

export default function Index() {
  const data = useLoaderData<LoaderResponse>();

  return (
    <div className="flex flex-col items-center">
      <NavBar />
      <Cover />
      <div className="h-12" />
      <p>{data.context.claims?.username}</p>
      {/* <SpotifyList /> */}
      <div className="h-[1000px]" />
    </div>
  );
}
