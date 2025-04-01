import { API_URL } from '$env/static/private';
import type { RequestParams } from '$lib/models/api';
import type { ErrorMessage } from '$lib/models/common';

export const createRequest = async <T = object>(
	params: RequestParams,
): Promise<T | ErrorMessage> => {
	const opts: RequestInit = {};
	let headers: HeadersInit = {};

	const url = new URL(`${API_URL}${params.path}`);
	if (params.query) {
		Object.entries(params.query).map((v, _) => url.searchParams.append(v[0], v[1]));
	}

	if (params.body) {
		headers = { 'Content-Type': 'application/json' };
		opts.body = params.body;
	}

	if (params.auth) headers = { ...headers, Authorization: `Bearer ${params.auth}` };

	opts.method = params.method;
	opts.headers = headers;

	const response = await fetch(url, opts);

	if (!response.ok) return (await response.json()) as ErrorMessage;

	const res = await response.json();
	console.log(url, opts);
	console.log(res.data);
	return res.data as T;
};
