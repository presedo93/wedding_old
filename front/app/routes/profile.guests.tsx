import { json, LoaderFunction } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { Errors } from "~/components/shared";
import { Button } from "~/components/ui/button";
import { authenticator, getAuthTokens } from "~/lib/auth.server";

type Loader = {
  readonly guests: string[];
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const { accessToken, headers } = await getAuthTokens(user, request);

  const res = await fetch(`${process.env.BACKEND_API_URL}/user/guests`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`);
  }

  const guests = await res.json();
  console.log("GUESTS", guests);

  return json<Loader>({ guests }, { headers });
};

export default function Guests() {
  const data = useLoaderData<Loader | undefined>();

  return (
    <>
      <h3 className="mt-6 font-sand text-2xl font-medium underline decoration-2 underline-offset-4">
        Acompanantes
      </h3>
      <div className="flex flex-col items-center justify-center">
        <span className="my-6 text-sm">
          No has anadido ningun acompanate aun!
        </span>
      </div>
      <Link className="flex w-full justify-center" to={"/profile/new-guest"}>
        <Button className="w-3/4 min-w-min">Nuevo acompanante</Button>
      </Link>
    </>
  );
}

export function ErrorBoundary() {
  return <Errors />;
}
