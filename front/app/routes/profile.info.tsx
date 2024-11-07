import { json, LoaderFunction } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { Errors } from "~/components";
import { Button } from "~/components/ui/button";
import { authenticator, getAuthTokens } from "~/lib/auth.server";
import { TodoItem } from "../components/todo-item";

type Loader = {
  readonly guests: { name: string }[];
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const { accessToken, headers } = await getAuthTokens(user, request);

  const res = await fetch(`${process.env.BACKEND_API_URL}/profiles/guests`, {
    method: "GET",
    headers: {
      ContentType: "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`);
  }

  const guests = await res.json();
  return json<Loader>({ guests }, { headers });
};

export default function Guests() {
  const data = useLoaderData<Loader>();

  return (
    <>
      <h3 className="mt-6 font-sand text-xl font-medium underline decoration-2 underline-offset-4">
        Acompanantes
      </h3>
      <div className="flex flex-col items-center justify-center">
        {data.guests.length ? (
          <GuestsList guests={data.guests} />
        ) : (
          <NoGuests />
        )}
      </div>
      <Link className="flex w-full justify-center" to={"/profile/guest"}>
        <Button className="w-2/3 min-w-min">Nuevo acompanante</Button>
      </Link>
      <h3 className="mt-2 font-sand text-xl font-medium underline decoration-2 underline-offset-4">
        Tareas
      </h3>
      <div className="mt-4 flex flex-col gap-5 rounded-lg bg-sky-700 p-4 shadow-md shadow-sky-800">
        <ul>
          <TodoItem>Completar el perfil</TodoItem>
          <TodoItem>Anadir acompanantes</TodoItem>
        </ul>
      </div>
    </>
  );
}

const GuestsList = ({ guests }: { guests: { name: string }[] }) => {
  return guests.map((g, i) => <div key={i}>Guest: {g.name}</div>);
};

const NoGuests = () => (
  <span className="my-6 text-sm">No has anadido ningun acompanate aun!</span>
);

export function ErrorBoundary() {
  return <Errors />;
}
