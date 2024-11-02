import { json, LoaderFunction } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { Errors } from "~/components/shared";
import { Button } from "~/components/ui/button";
import { authenticator } from "~/lib/auth.server";

type Loader = {
  readonly guests: string[];
};

export const loader: LoaderFunction = async ({ request }) => {
  const auth = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  let guests: string[];
  try {
    const res = await fetch(`${process.env.WEDDING_BACK_URL}/user/guests`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${auth.accessToken}`,
      },
    });

    guests = await res.json();
  } catch {
    throw new Error("Error al cargar los acompanantes");
  }

  return json<Loader>({ guests });
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
