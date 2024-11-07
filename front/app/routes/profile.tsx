import { json, LoaderFunction } from "@remix-run/node";
import { Link, Outlet, redirect, useLoaderData } from "@remix-run/react";
import { authenticator, getAuthTokens } from "~/lib/auth.server";
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { Back } from "~/icons";
import { Button } from "~/components/ui/button";
import { fetchAPI } from "~/lib/fetch.server";
import { Profile } from "~/lib/models";

type Loader = {
  readonly email: string;
  readonly profile: Profile | undefined;
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  let profile: Profile | undefined;
  const { accessToken, headers } = await getAuthTokens(user, request);

  try {
    profile = await fetchAPI<Profile>("/profiles", { accessToken });
  } catch {
    profile = undefined;
  }

  if (!profile && !request.url.includes("/profile/edit")) {
    return redirect("/profile/edit", { headers });
  }

  return json<Loader>({ email: user.email, profile }, { headers });
};

export default function EditProfile() {
  const data = useLoaderData<Loader>();

  return (
    <div className="flex min-h-svh flex-col bg-sky-200 p-8">
      <Link to={"/"}>
        <div className="flex flex-row items-center gap-2">
          <Back className="size-8" />
          <p>Volver</p>
        </div>
      </Link>
      <h1 className="mt-4 font-sand text-4xl font-bold">Mi perfil</h1>
      <div className="mt-4 flex flex-row items-center gap-5 rounded-lg bg-sky-950 p-4 shadow-md shadow-sky-800">
        <Avatar className="size-14">
          <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <p className="text-lg text-white">
          Hola,
          <span className="font-bold">{data.profile?.name ?? data.email}</span>!
        </p>
      </div>
      <Outlet context={data.profile} />
      <Link className="mt-8 flex justify-center" to={"/profile/edit"}>
        <Button className="w-2/3 min-w-min">Editar perfil</Button>
      </Link>
    </div>
  );
}
