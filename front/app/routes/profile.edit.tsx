import { useRemixForm, getValidatedFormData } from "remix-hook-form";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import * as zod from "zod";
import { ActionFunctionArgs, json, redirect } from "@remix-run/node";
import { Link, Form, useOutletContext } from "@remix-run/react";
import { Label } from "~/components/ui/label";
import { Errors } from "~/components";
import { authenticator, getAuthTokens } from "~/lib/auth.server";
import { fetchAPI } from "~/lib/fetch.server";
import { profileSchema } from "~/lib/schemas";

type FormData = zod.infer<typeof profileSchema>;
const resolver = zodResolver(profileSchema);

export const action = async ({ request }: ActionFunctionArgs) => {
  const { errors, data, receivedValues } = await getValidatedFormData<FormData>(
    request,
    resolver
  );

  if (errors) {
    return json({ errors, receivedValues });
  }

  const user = await authenticator.isAuthenticated(request);
  if (!user) return redirect("/");

  const { accessToken, headers } = await getAuthTokens(user, request);
  await fetchAPI<FormData>("/profiles", {
    accessToken,
    method: "POST",
    body: { id: user.id, ...data },
  });

  return redirect("/profile/info", { headers });
};

export default function NewProfile() {
  const profile = useOutletContext();

  const {
    handleSubmit,
    formState: { errors },
    register,
  } = useRemixForm<FormData>({
    mode: "onSubmit",
    resolver,
  });

  return (
    <div className="mt-4 flex flex-col rounded-md border border-sky-950 p-6">
      {!profile && (
        <p className="mb-2 text-sm">
          Aun no has creado tu perfil, rellena algunos datos para disfrutar de
          la web!
        </p>
      )}
      <Form onSubmit={handleSubmit} method="post" className="space-y-2">
        <div>
          <Label>Nombre</Label>
          <Input placeholder="Nombre..." {...register("name")} />
        </div>

        <div className="mt-2">
          <Label>Num. de telefono</Label>
          <Input placeholder="697449080" {...register("phone")} />
        </div>

        <div className="mt-2">
          <Label>Mi email</Label>
          <Input placeholder="email@dominio.es" {...register("email")} />
        </div>

        {Object.entries(errors).map(([key, value]) => (
          <div key={key} className="rounded-md bg-destructive p-2">
            <p className="text-sm text-destructive-foreground">
              {value.message}
            </p>
          </div>
        ))}

        <div className="flex flex-row justify-center space-x-3 pt-4">
          <Link className="w-1/2" to={"/profile/info"}>
            <Button variant={"destructive"} className="w-full min-w-min">
              Cancelar
            </Button>
          </Link>
          <Button type="submit" className="w-1/2 min-w-min">
            Guardar
          </Button>
        </div>
      </Form>
    </div>
  );
}

export function ErrorBoundary() {
  return <Errors />;
}
