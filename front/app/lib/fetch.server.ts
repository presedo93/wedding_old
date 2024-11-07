const { BACKEND_API_URL } = process.env;

export async function fetchAPI<T>(
  endpoint: string,
  {
    body,
    method,
    accessToken,
  }: {
    body?: unknown;
    method?: string;
    accessToken: string;
  }
): Promise<T | undefined> {
  const res = await fetch(`${BACKEND_API_URL}${endpoint}`, {
    method: method || "GET",
    body: JSON.stringify(body),
    headers: {
      ContentType: "application/json",
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!res.ok) {
    throw new Error(`HTTP ${res.status} for ${method} ${endpoint}!`);
  }

  return await res.json();
}
