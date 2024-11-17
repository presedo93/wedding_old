import * as zod from "zod";

const phoneNumberSchema = zod
  .string()
  .trim()
  .regex(/^\s*$|^\+?(\d\s*){8,16}$/, "Numero de telefono invalido")
  .transform((value) => value.replace(/\s/g, ""));

const allergiesSchema = zod
  .union([zod.string().trim(), zod.string().trim().array()])
  .transform((v) => (Array.isArray(v) ? v : v.split(",").filter(Boolean)));

export const guestSchema = zod.object({
  id: zod.number().optional(),
  name: zod.string().min(1, "El nombre es necesario"),
  is_vegetarian: zod.coerce.boolean().default(false),
  needs_transport: zod.coerce.boolean().default(false),
  allergies: allergiesSchema,
  phone: phoneNumberSchema,
});

export const profileSchema = zod.object({
  id: zod.string().uuid().optional(),
  name: zod.string().min(1, "El nombre es necesario"),
  phone: phoneNumberSchema,
  email: zod.string().email(),
  // picture_url: zod.string().url(), // TODO: Add this field
});
