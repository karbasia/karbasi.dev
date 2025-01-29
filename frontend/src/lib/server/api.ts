import { API_URL } from '$env/static/private';
import type { ErrorMessage } from '$lib/models/common';

export enum HttpRequest {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  PATCH = 'PATCH',
  DELETE = 'DELETE',
}

export interface RequestParams {
  method: HttpRequest;
  path: string;
  body?: string;
  auth?: string;
}

export const createRequest = async <T=object>(params: RequestParams): Promise<T | ErrorMessage> => {
  const opts: RequestInit = {};
  let headers: HeadersInit = {};

  if (params.body) {
    headers = { 'Content-Type': 'application/json' };
    opts.body = params.body;
  }

  if (params.auth) headers = { ...headers, Authorization: `Bearer ${params.auth}` };

  opts.method = params.method;
  opts.headers = headers;

  const response = await fetch(`${API_URL}${params.path}`, opts);

  if (!response.ok) return (await response.json()) as ErrorMessage;

  const res =  await response.json();
  return res.data as T;
}