/* eslint-disable @typescript-eslint/no-explicit-any */

import { User } from "../../types";
import {
  BaseServerResponse,
  FetchError,
  HTTPMethod,
  InternalServerError,
} from "./Entities";
import { ServerEndpoints } from "./ServerEndpoints";

export default async function performRequest<T>(
  method: HTTPMethod,
  path: string,
  body: any | null = null
): Promise<T | null> {
  // Build request
  const request = new Request(path, {
    body: body === null ? null : JSON.stringify(body),
    method: HTTPMethod[method],
    headers: {
      "Content-Type": "application/json",
      "Bearer-Token": retreiveAuthToken(),
    },
  });

  // Await response and decode
  const response = await fetch(request);
  const serverResponseBody = (await response.json()) as BaseServerResponse<T>;

  // Check if we received a token in order to store it locally
  const authToken = response.headers.get("Bearer-Token");
  if (authToken) {
    storeAuthToken(authToken);
  }

  // Check for error
  if (serverResponseBody.error) {
    // If exired token error trigger refresh
    if (
      serverResponseBody.internalErrorCode === InternalServerError.ExpiredToken
    ) {
      await refreshAuthToken();
      return performRequest(method, path, body);
    }
    throw new FetchError(
      serverResponseBody.errorText || "Fetch errror occured",
      serverResponseBody.internalErrorCode,
      response.status
    );
  }

  // Store user email and password on userMe fetch
  if (path === ServerEndpoints.userMe()) {
    const userPassword = (serverResponseBody.data as User).password;
    storeUserPassword(userPassword);
  }

  // Return received data
  const data : BaseServerResponse = {
    data: serverResponseBody.data,
    hearderStatus: response.ok
  }
  return data;
}

export function performLogOut() {
  localStorage.setItem("Bearer-Token", "");
  localStorage.setItem("password", "");
}

// ---------- Private API ----------

function storeAuthToken(token: string) {
  localStorage.setItem("Bearer-Token", token);
}

function retreiveAuthToken(): string {
  return localStorage.getItem("Bearer-Token") || "";
}

function storeUserPassword(password: string) {
  localStorage.setItem("password", password);
}

function retreivePassword(): string {
  return localStorage.getItem("password") || "";
}

async function refreshAuthToken() {
  const userData = {
    password: retreivePassword(),
  };
  try {
    await performRequest(
      HTTPMethod.POST,
      ServerEndpoints.refreshToken(),
      userData
    );
  } catch (error) {
    console.log(`💥 Failed to perform auth token refresh with error: ${error}`);
    throw error;
  }
}
