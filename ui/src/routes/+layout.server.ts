import type { LayoutServerLoad } from './$types';
import type { Post } from '$lib/models/post';
import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import { error } from '@sveltejs/kit';

export const load: LayoutServerLoad = async ({ locals }) => {
	const params: RequestParams = {
		method: HttpRequest.GET,
		path: '/posts',
	};
	const data = await createRequest<Post[]>(params);
	if ('error' in data) return error(data.code, data.error);

	return {
		posts: data,
		user: locals.user,
	};
};
