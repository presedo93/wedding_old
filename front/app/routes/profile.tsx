import { json, LoaderFunction } from "@remix-run/node";
import { Link, Outlet, redirect, useLoaderData } from "@remix-run/react";
import { authenticator, tokenizer } from "~/lib/auth.server";
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar";
import { BackIcon } from "~/icons";
import { Button } from "~/components/ui/button";
import { fetchAPI } from "~/lib/fetch.server";
import { Profile } from "~/lib/models";

type Loader = {
  readonly profile: Profile | undefined;
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  let profile: Profile | undefined;
  const accessToken = await tokenizer(request, user, { shouldRefresh: true });

  try {
    profile = await fetchAPI<Profile>("/profiles", { accessToken });
  } catch {
    profile = undefined;
  }

  if (!profile && !request.url.includes("/profile/edit")) {
    return redirect("/profile/edit");
  }

  return json<Loader>({ profile });
};

export default function EditProfile() {
  const data = useLoaderData<Loader>();

  return (
    <div className="flex min-h-svh flex-col bg-sky-200 p-8">
      <h1 className="mt-4 font-sand text-4xl font-bold">Mi perfil</h1>
      <BackButton />
      {data.profile && <ProfileCard profile={data.profile} />}
      <Outlet />
      <Link className="mt-8 flex justify-center" to={"/profile/edit"}>
        <Button className="w-2/3 min-w-min">Editar perfil</Button>
      </Link>
    </div>
  );
}

const ProfileCard = ({ profile }: { profile: Profile }) => (
  <div className="mt-4 flex flex-row items-center gap-5 rounded-lg bg-sky-950 p-4 shadow-md shadow-sky-800">
    <Avatar className="size-14">
      <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
      <AvatarFallback>CN</AvatarFallback>
    </Avatar>
    <p className="mr-4 line-clamp-1 text-ellipsis text-lg text-white">
      Hola, <span className="font-bold">{profile.name}</span>!
    </p>
  </div>
);

const BackButton = () => (
  <Link to={"/"}>
    <div className="flex flex-row items-center gap-2">
      <BackIcon className="size-8" />
      <p>Volver</p>
    </div>
  </Link>
);
