import { Authenticator } from "remix-auth";
import { OAuth2Strategy } from "remix-auth-oauth2";
import { sessionStorage } from "~/lib/session.server";
import { refreshCookie } from "./cookies.server";

// Define the type for authenticated user data
export type User = {
  accessToken: string;
  refreshToken?: string;
  expiresAt: number;
  id: string;
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
      id: info.sub,
      email: info?.email,
      email_verified: info?.email_verified === "true",
    };
  }
);

// Use the OAuth2 strategy with the authenticator
authenticator.use(strategy);

function isTokenExpired(expiresAt: number): boolean {
  return expiresAt < Date.now();
}

async function getRefreshedAccessToken(
  req: Request
): Promise<{ accessToken: string; headers?: Headers }> {
  const cookie = await refreshCookie.parse(req.headers.get("Cookie"));
  const session = await sessionStorage.getSession(req.headers.get("Cookie"));

  const tokens = await strategy.refreshToken(cookie);
  const { refresh_token, access_token } = tokens;
  session.set(authenticator.sessionKey as "user", {
    accessToken: access_token,
    ...session,
  });

  const headers = new Headers();
  headers.append("Set-Cookie", await sessionStorage.commitSession(session));
  headers.append("Set-Cookie", await refreshCookie.serialize(refresh_token));

  return { accessToken: access_token, headers };
}

export async function getAuthTokens(
  user: User,
  req: Request
): Promise<{ accessToken: string; headers?: Headers }> {
  if (isTokenExpired(user.expiresAt)) {
    console.log("Token expired, refreshing...");
    return await getRefreshedAccessToken(req);
  }
  return { accessToken: user.accessToken };
}
