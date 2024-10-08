import { getRootKey } from "@/app/auth/auth.client";
import { apiKeyperUrl } from "./config";

interface CreateApiKeyProps {
  apiId: string;
  keyName: string;
  prefix?: string;
  permissions?: string;
  rateLimitConfig?: {
    Type?: string;
    Limit: number;
    Period: string;
    Window: string;
  };
}


export async function createApiKey(createApiKeyProps: CreateApiKeyProps) {
  const rootKey = await getRootKey();
  const response = await fetch(`${apiKeyperUrl}/apiKey`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${rootKey}`,
    },
    body: JSON.stringify({
      "apiId": createApiKeyProps.apiId,
      "name": createApiKeyProps.keyName,
      "prefix": createApiKeyProps.prefix,
      "permissions": createApiKeyProps.permissions,
      "rateLimit": createApiKeyProps.rateLimitConfig,
    }),
  })

  const data = await response.json();

  if (!response.ok) {
    throw new Error("Failed to create a new api key");
  }
}