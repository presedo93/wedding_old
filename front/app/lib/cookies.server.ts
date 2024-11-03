import { createCookie } from "@remix-run/node"; // or cloudflare/deno

export const refreshCookie = createCookie("cognito_refresh", {
  maxAge: 14 * 24 * 60 * 60, // 14 days
  sameSite: "lax",
  path: "/",
  httpOnly: true,
  secrets: ["s3cr3t"],
  secure: process.env.NODE_ENV === "production", // enabled only in prod
});
