import { Authenticator } from "remix-auth";
import { OAuth2Strategy } from "remix-auth-oauth2";
import { sessionStorage } from "~/lib/session.server";

// Define the type for authenticated user data
export type Authenticated = {
  accessToken: string;
  email?: string;
  email_verified?: boolean;
};

// User info from Cognito
type UserInfo = {
  sub: string;
  email_verified: string;
  email: string;
  username: string;
};

// Initialize the authenticator with session storage
export const authenticator = new Authenticator<Authenticated>(sessionStorage);

// Configure OAuth2 strategy with Cognito
export const strategy = new OAuth2Strategy(
  {
    clientId: process.env.COGNITO_APP_ID!,
    clientSecret: process.env.COGNITO_APP_SECRET!,
    redirectURI: process.env.COGNITO_CALLBACK_URL!,
    authorizationEndpoint: `${process.env.COGNITO_POOL_URL}/oauth2/authorize`,
    tokenEndpoint: `${process.env.COGNITO_POOL_URL}/oauth2/token`,
  },
  async ({ tokens }): Promise<Authenticated> => {
    try {
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

      const info: UserInfo = await response.json();
      return {
        accessToken: tokens.access_token,
        email: info?.email,
        email_verified: info?.email_verified === "true",
      };
    } catch (e) {
      throw new Error(`Authorization error: ${e}`);
    }
  }
);

// Use the OAuth2 strategy with the authenticator
authenticator.use(strategy);
