import { useRemixForm, getValidatedFormData } from "remix-hook-form";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import * as zod from "zod";
import {
  ActionFunctionArgs,
  json,
  LoaderFunction,
  redirect,
} from "@remix-run/node";
import { Link, Form } from "@remix-run/react";
import { Checkbox } from "~/components/ui/checkbox";
import { Label } from "~/components/ui/label";
import { Errors } from "~/components";
import { authenticator, getAuthTokens } from "~/lib/auth.server";
import { fetchAPI } from "~/lib/fetch.server";
import { guestSchema } from "~/lib/schemas";
import { Guest } from "~/lib/models";

type FormData = zod.infer<typeof guestSchema>;
const resolver = zodResolver(guestSchema);

type Loader = {
  readonly guest: Guest | undefined;
};

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.isAuthenticated(request, {
    failureRedirect: "/",
  });

  let guest: Guest | undefined;
  const { accessToken, headers } = await getAuthTokens(user, request);

  try {
    guest = await fetchAPI<Guest>("/profiles", { accessToken });
  } catch {
    throw new Error("Ha habido un error al cargar el perfil");
  }

  return json<Loader>({ guest }, { headers });
};

export default function EditGuest() {
  const {
    handleSubmit,
    formState: { errors },
    watch,
    register,
    setValue,
  } = useRemixForm<FormData>({
    mode: "onSubmit",
    resolver,
  });

  const is_vegetarian = watch("is_vegetarian");
  const needs_transport = watch("needs_transport");

  return (
    <>
      <h3 className="mb-2 mt-6 font-sand text-xl font-medium underline decoration-2 underline-offset-4">
        Nuevo Acompanante
      </h3>
      <Form onSubmit={handleSubmit} method="post" className="space-y-2">
        <div className="mt-2">
          <Label>Nombre</Label>
          <Input placeholder="Nombre..." {...register("name")} />
        </div>

        <div className="mt-2">
          <Label>Num. de telefono</Label>
          <Input placeholder="697 44 90 80" {...register("phone")} />
        </div>

        <div className="mt-2">
          <Label>Alergias</Label>
          <Input
            placeholder="Marisco, frutos secos..."
            {...register("allergies")}
          />
        </div>

        <div className="mt-4 flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4 shadow">
          <Checkbox
            name="is_vegetarian"
            onClick={() => setValue("is_vegetarian", !is_vegetarian)}
          />
          <div className="space-y-1 leading-none">
            <Label>Quieres menu vegetariano?</Label>
          </div>
        </div>

        <div className="mt-4 flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4 shadow">
          <Checkbox
            {...register("needs_transport")}
            name="needs_transport"
            onClick={() => setValue("needs_transport", !needs_transport)}
          />
          <div className="space-y-1 leading-none">
            <Label>Quieres ir y volver en autobus?</Label>
          </div>
        </div>

        {Object.entries(errors).map(([key, value]) => (
          <div key={key} className="rounded-md bg-destructive p-2">
            <p className="text-sm text-destructive-foreground">
              {value.message}
            </p>
          </div>
        ))}

        <div className="flex flex-row justify-center space-x-3">
          <Link className="w-1/2" to={"/profile/info"}>
            <Button variant={"destructive"} className="w-full min-w-min">
              Cancelar
            </Button>
          </Link>
          <Button type="submit" className="w-1/2 min-w-min">
            Submit
          </Button>
        </div>
      </Form>
    </>
  );
}

export const action = async ({ request }: ActionFunctionArgs) => {
  const { errors, data, receivedValues } = await getValidatedFormData<FormData>(
    request,
    resolver
  );

  if (errors) {
    return json({ errors, receivedValues });
  }

  const user = await authenticator.isAuthenticated(request);
  if (!user) throw new Error("Ha habido un error al autenticar al usuario");

  const { accessToken, headers } = await getAuthTokens(user, request);
  await fetchAPI<FormData>("/guests", {
    accessToken,
    body: data,
    method: "POST",
  });

  return redirect("/profile/info", { headers });
};

export function ErrorBoundary() {
  return <Errors />;
}
