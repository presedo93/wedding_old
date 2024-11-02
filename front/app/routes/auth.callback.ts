import type { LoaderFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";

import { authenticator } from "~/lib/auth.server";
import { sessionStorage } from "~/lib/session.server";

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.authenticate("oauth2", request);

  // Get the current session
  const session = await sessionStorage.getSession(
    request.headers.get("Cookie")
  );

  // Store authenticated user details in session
  session.set(authenticator.sessionKey as "user", user);

  // Commit the session
  const headers = new Headers({
    "Set-Cookie": await sessionStorage.commitSession(session),
  });

  // Redirect to the application root with updated session
  return redirect("/", { headers });
};
