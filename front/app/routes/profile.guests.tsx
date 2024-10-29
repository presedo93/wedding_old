import { json, LoaderFunction, redirect } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { Button } from "~/components/ui/button";
import { logto } from "~/lib/auth.server";

type Loader = {
  readonly id: string;
};

export const loader: LoaderFunction = async ({ request }) => {
  const context = await logto.getContext({ getAccessToken: true })(request);

  if (!context.isAuthenticated) {
    return redirect("/");
  }

  // const userId = context.accessToken;

  return json<Loader | undefined>({ id: "1" });
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
