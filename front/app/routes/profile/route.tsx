import { json, LoaderFunction, redirect } from "@remix-run/node";
import { Link, Outlet, useLoaderData, useLocation } from "@remix-run/react";
import { getManagementToken, logto } from "~/lib/auth.server";
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { Back } from "~/icons";
import { Button } from "~/components/ui/button";
import { TodoItem } from "./todo-item";

type Loader = {
  readonly id: string;
  readonly username: string;
  readonly primaryEmail?: string;
  readonly primaryPhone?: string;
  readonly name?: string;
  readonly avatar?: string;
  readonly customData?: unknown;
  readonly identities?: unknown;
  readonly lastSignInAt?: number;
  readonly createdAt?: number;
  readonly updatedAt?: number;
  readonly profile?: unknown;
  readonly applicationId?: string;
  readonly isSuspended?: boolean;
  readonly hasPassword?: boolean;
};

export const loader: LoaderFunction = async ({ request }) => {
  const context = await logto.getContext({ getAccessToken: false })(request);

  if (!context.isAuthenticated) {
    return redirect("/");
  }

  const userId = context.claims?.sub;
  const { access_token: managementToken } = await getManagementToken();

  const userReq = await fetch(
    `${process.env.LOGTO_ENDPOINT}/api/users/${userId}`,
    {
      method: "GET",
      headers: { Authorization: `Bearer ${managementToken}` },
    }
  );

  const user = await userReq.json();
  return json<Loader>(user);
};

export default function Profile() {
  const data = useLoaderData<Loader>();

  return (
    <div className="flex min-h-svh flex-col bg-sky-300 p-8">
      <Link to={"/"}>
        <div className="flex flex-row items-center gap-2">
          <Back className="size-8" />
          <p>Volver</p>
        </div>
      </Link>
      <h1 className="mt-4 font-sand text-4xl font-bold underline decoration-2 underline-offset-4">
        Perfil
      </h1>
      <div className="mt-4 flex flex-row items-center gap-5 rounded-lg bg-sky-950 p-4 shadow-md shadow-sky-800">
        <Avatar className="size-14">
          <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <p className="text-lg font-medium text-white">
          Hola, <span>{data.name ?? data.username}!</span>
        </p>
      </div>
      <Outlet />
      <h3 className="mt-2 font-sand text-2xl font-medium underline decoration-2 underline-offset-4">
        Tareas
      </h3>
      <div className="mt-4 flex flex-col gap-5 rounded-lg bg-sky-700 p-4 shadow-md shadow-sky-800">
        <ul>
          <TodoItem>Completar el perfil</TodoItem>
          <TodoItem>Anadir acompanantes</TodoItem>
        </ul>
      </div>
      <div className="mt-8 flex flex-row justify-around">
        <Link className="flex w-1/2 justify-center" to={"/profile/edit"}>
          <Button className="w-2/3 min-w-min">Editar perfil</Button>
        </Link>
        <GuestButton />
      </div>
    </div>
  );
}

const GuestButton = () => {
  const location = useLocation();

  if (location.pathname === "/profile/guests") {
    return (
      <Link className="flex w-1/2 justify-center" to={"/profile"}>
        <Button className="w-2/3 min-w-min">Ocultar</Button>
      </Link>
    );
  }

  return (
    <Link className="flex w-1/2 justify-center" to={"/profile/guests"}>
      <Button className="w-2/3 min-w-min">Acompanantes</Button>
    </Link>
  );
};
