import type { LoaderFunction } from "@remix-run/node";

import { authenticator } from "~/lib/auth.server";

export const loader: LoaderFunction = async ({ request }) => {
  await authenticator.authenticate("oauth2", request, { successRedirect: "/" });
};
