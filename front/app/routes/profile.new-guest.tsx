import { useRemixForm, getValidatedFormData } from "remix-hook-form";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import * as zod from "zod";
import { ActionFunctionArgs, json, redirect } from "@remix-run/node";
import { Link, Form } from "@remix-run/react";
import { Checkbox } from "~/components/ui/checkbox";
import { Label } from "~/components/ui/label";
import { Errors } from "~/components/shared";

const allergiesSchema = zod
  .union([zod.string().trim(), zod.string().trim().array()])
  .transform((v) => (Array.isArray(v) ? v : v.split(",").filter(Boolean)));

const phoneNumberSchema = zod
  .string()
  .regex(/^(\+?[1-9]\d{8,16})?$/, "Numero de telefono invalido");

const schema = zod.object({
  name: zod.string().min(1, "El nombre es necesario"),
  is_vegetarian: zod.coerce.boolean().default(false),
  needs_transport: zod.coerce.boolean().default(false),
  allergies: allergiesSchema,
  phone: phoneNumberSchema,
});

type FormData = zod.infer<typeof schema>;
const resolver = zodResolver(schema);

export const action = async ({ request }: ActionFunctionArgs) => {
  const { errors, data, receivedValues } = await getValidatedFormData<FormData>(
    request,
    resolver
  );

  if (errors) {
    return json({ errors, receivedValues });
  }

  const accessToken = await getAccessToken(request);
  console.log("Access Token", accessToken);

  try {
    await fetch(`${process.env.LOGTO_WEDDING_RESOURCE}/guests`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
      body: JSON.stringify(data),
    });
  } catch {
    throw new Error("Error al crear el usuario");
  }

  return redirect("/profile/guests");
};

export default function GuestsNew() {
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
      <h3 className="mb-2 mt-6 font-sand text-2xl font-medium underline decoration-2 underline-offset-4">
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
          <Link className="w-1/2" to={"/profile/guests"}>
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

export function ErrorBoundary() {
  return <Errors />;
}
