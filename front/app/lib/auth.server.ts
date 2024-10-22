import { createCookieSessionStorage } from "@remix-run/node";
import { makeLogtoRemix } from "@logto/remix";

const sessionStorage = createCookieSessionStorage({
  cookie: {
    name: "logto-session",
    maxAge: 14 * 24 * 60 * 60,
    secrets: ["s3cr3t"],
  },
});

export const logto = makeLogtoRemix(
  {
    endpoint: process.env.LOGTO_ENDPOINT!,
    appId: process.env.LOGTO_APP_ID!,
    appSecret: process.env.LOGTO_APP_SECRET!,
    baseUrl: process.env.LOGTO_BASE_URL!,
    resources: [
      process.env.LOGTO_WEDDING_RESOURCE!,
      process.env.LOGTO_MANAGEMENT_RESOURCE!,
    ],
  },
  { sessionStorage }
);

export function logtoAuthHeader() {
  const appId = process.env.LOGTO_MANAGEMENT_ID!;
  const appSecret = process.env.LOGTO_MANAGEMENT_SECRET!;

  return `Basic ${Buffer.from(`${appId}:${appSecret}`).toString("base64")}`;
}

interface ManagementResponse {
  access_token: string;
  expires_in: number;
  token_type: string;
  scope: string;
}

export async function getManagementToken() {
  const res = await fetch(`${process.env.LOGTO_ENDPOINT}/oidc/token`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
      Authorization: logtoAuthHeader(),
    },
    body: new URLSearchParams({
      grant_type: "client_credentials",
      resource: process.env.LOGTO_MANAGEMENT_RESOURCE!,
      scope: "all",
    }).toString(),
  });

  const json = await res.json();
  return json as ManagementResponse;
}
