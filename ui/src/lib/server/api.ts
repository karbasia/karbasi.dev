import { env } from '$env/dynamic/private';
import type { RequestParams } from '$lib/models/api';
import type { ErrorMessage } from '$lib/models/common';

export const createRequest = async <T = object>(
	params: RequestParams,
): Promise<T | ErrorMessage> => {
	const opts: RequestInit = {};
	let headers: HeadersInit = {};

	const url = new URL(`${env.API_URL}${params.path}`);
	if (params.query) {
		Object.entries(params.query).map((v, _) => url.searchParams.append(v[0], v[1]));
	}

	if (params.body) {
		headers = { 'Content-Type': 'application/json' };
		opts.body = JSON.stringify(params.body);
	} else if (params.formData) {
		opts.body = params.formData;
	}

	if (params.auth) headers = { ...headers, Authorization: `Bearer ${params.auth}` };

	opts.method = params.method;
	opts.headers = headers;

	const response = await fetch(url, opts);

	if (!response.ok) return (await response.json()) as ErrorMessage;

	const res = await response.json();
	return res.data as T;
};

// Used for calling the file download endpoint which requires the entire response object
export const createRawRequest = async (params: RequestParams): Promise<Response | ErrorMessage> => {
	const opts: RequestInit = {};
	let headers: HeadersInit = {};

	const url = new URL(`${env.API_URL}${params.path}`);
	if (params.query) {
		Object.entries(params.query).map((v, _) => url.searchParams.append(v[0], v[1]));
	}

	if (params.body) {
		headers = { 'Content-Type': 'application/json' };
		opts.body = JSON.stringify(params.body);
	} else if (params.formData) {
		opts.body = params.formData;
	}

	if (params.auth) headers = { ...headers, Authorization: `Bearer ${params.auth}` };

	opts.method = params.method;
	opts.headers = headers;

	const response = await fetch(url, opts);
	if (!response.ok) return (await response.json()) as ErrorMessage;

	return response;
};
