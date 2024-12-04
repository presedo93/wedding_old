import type { ActionFunction, LoaderFunction } from "@remix-run/node";
import { redirect } from "@remix-run/node";

import { sessionStorage } from "~/lib/session.server";

const handleLogout = async (request: Request) => {
  const session = await sessionStorage.getSession(
    request.headers.get("Cookie")
  );

  const { COGNITO_APP_ID, COGNITO_POOL_URL, COGNITO_LOGOUT_URL } = process.env;
  const cognitoLogout = new URL(`${COGNITO_POOL_URL!}/logout`);

  cognitoLogout.searchParams.set("client_id", COGNITO_APP_ID!);
  cognitoLogout.searchParams.set("logout_uri", `${COGNITO_LOGOUT_URL!}`);

  return redirect(cognitoLogout.toString(), {
    headers: {
      "Set-Cookie": await sessionStorage.destroySession(session),
    },
  });
};

export const loader: LoaderFunction = async ({ request }) => {
  return await handleLogout(request);
};

export const action: ActionFunction = async ({ request }) => {
  return await handleLogout(request);
};
