import { redirect } from "@remix-run/react";
import { Authenticator } from "remix-auth";
import { OAuth2Strategy } from "remix-auth-oauth2";

import { sessionStorage } from "~/lib/session.server";
import { refreshCookie } from "./cookies.server";

// Define the type for authenticated user data
export type User = {
  accessToken: string;
  expiresAt: number;
  email: string;
  email_verified?: boolean;
};

export type Authenticated = User & { refreshToken?: string };

// User info from Cognito
type CognitoUser = {
  sub: string;
  email_verified: string;
  email: string;
  username: string;
};

// Initialize the authenticator with session storage
export const authenticator = new Authenticator<User>(sessionStorage);

// Configure OAuth2 strategy with Cognito
const strategy = new OAuth2Strategy(
  {
    clientId: process.env.COGNITO_APP_ID!,
    clientSecret: process.env.COGNITO_APP_SECRET!,
    redirectURI: process.env.COGNITO_CALLBACK_URL!,
    authorizationEndpoint: `${process.env.COGNITO_POOL_URL}/oauth2/authorize`,
    tokenEndpoint: `${process.env.COGNITO_POOL_URL}/oauth2/token`,
  },
  async ({ tokens }): Promise<Authenticated> => {
    const response = await fetch(
      `${process.env.COGNITO_POOL_URL}/oauth2/userInfo`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${tokens.access_token}`,
          ContentType: "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error(`Authorization error`);
    }

    const info: CognitoUser = await response.json();
    return {
      accessToken: tokens.access_token,
      refreshToken: tokens.refresh_token,
      expiresAt: Date.now() + (tokens.expires_in ?? 0) * 1000,
      email: info?.email,
      email_verified: info?.email_verified === "true",
    };
  }
);

// Use the OAuth2 strategy with the authenticator
authenticator.use(strategy);

export async function tokenizer(
  request: Request,
  user: User,
  {
    headers = new Headers(),
  }: { headers?: Headers; shouldRefresh?: boolean } = {}
) {
  if (user.expiresAt < Date.now()) await refreshTokens(request, user, headers);

  return user.accessToken;
}

async function refreshTokens(request: Request, user: User, headers: Headers) {
  console.log("Token expired, refreshing...");
  const cookie = request.headers.get("Cookie");

  const refresh = await refreshCookie.parse(cookie);
  const session = await sessionStorage.getSession(cookie);

  const tokens = await strategy.refreshToken(refresh);
  const { access_token, expires_in } = tokens;

  session.set(authenticator.sessionKey, {
    ...user,
    accessToken: access_token,
    expiresAt: Date.now() + (expires_in ?? 0) * 1000,
  });

  headers.append("Set-Cookie", await sessionStorage.commitSession(session));
  if (request.method === "GET") throw redirect(request.url, { headers });

  return access_token;
}
