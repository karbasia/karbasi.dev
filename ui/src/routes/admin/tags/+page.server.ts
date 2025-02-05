import { type Tag } from '$lib/models/tag';
import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import type { PageServerLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async () => {
	const params: RequestParams = {
		method: HttpRequest.GET,
		path: '/tags',
	};
	const data = await createRequest<Tag[]>(params);
	if ('error' in data) return error(data.code, data.error);

	return {
		tags: data,
	};
};
