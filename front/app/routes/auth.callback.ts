import type { LoaderFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";

import { authenticator } from "~/lib/auth.server";
import { refreshCookie } from "~/lib/cookies.server";
import { sessionStorage } from "~/lib/session.server";

export const loader: LoaderFunction = async ({ request }) => {
  const user = await authenticator.authenticate("oauth2", request);
  const { refreshToken, ...rest } = user;

  // Get the current session
  const session = await sessionStorage.getSession(
    request.headers.get("Cookie")
  );

  // Store authenticated user details in session
  session.set(authenticator.sessionKey, rest);

  // Commit the session
  const headers = new Headers();
  headers.append("Set-Cookie", await sessionStorage.commitSession(session));
  headers.append("Set-Cookie", await refreshCookie.serialize(refreshToken));

  // Redirect to the application root with updated session
  return redirect("/", { headers });
};
