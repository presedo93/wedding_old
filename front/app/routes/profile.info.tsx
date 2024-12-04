import { ActionFunctionArgs, json, LoaderFunction } from "@remix-run/node";
import { Link, redirect, useLoaderData } from "@remix-run/react";
import { Errors, GuestCard } from "~/components";
import { Button } from "~/components/ui/button";
import { authenticator, tokenizer } from "~/lib/auth.server";
import { TodoItem } from "../components/todo-item";
import { Guest } from "~/lib/models";
import { fetchAPI } from "~/lib/fetch.server";

type Loader = {
  readonly guests: Guest[];
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  const accessToken = await tokenizer(request, user);
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
  return json<Loader>({ guests });
};

export default function Guests() {
  const data = useLoaderData<Loader>();

  return (
    <>
      <h3 className="mt-6 font-sand text-xl font-medium underline decoration-2 underline-offset-4">
        Acompanantes
      </h3>
      <div className="my-2 flex flex-col items-center justify-center">
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

export const action = async ({ request }: ActionFunctionArgs) => {
  const id = new URL(request.url).searchParams.get("id");

  console.log(id);

  const user = await authenticator.isAuthenticated(request);
  if (!user) throw new Error("Ha habido un error al autenticar al usuario");

  const headers = new Headers();
  const accessToken = await tokenizer(request, user, { headers });
  const url = `/guests/${id}`;

  await fetchAPI<FormData>(url, {
    accessToken,
    method: "delete",
  });

  return redirect("/profile/info", { headers });
};

const GuestsList = ({ guests }: { guests: Guest[] }) => {
  return guests.map((g, i) => <GuestCard guest={g} key={i} />);
};

const NoGuests = () => (
  <span className="my-6 text-sm">No has anadido ningun acompanate aun!</span>
);

export function ErrorBoundary() {
  return <Errors />;
}
