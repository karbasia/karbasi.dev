import type { LayoutServerLoad } from './$types';
import type { Post } from '$lib/models/post';
import { createRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const load: LayoutServerLoad = async ({ locals }) => {
	const params: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/posts',
	};
	if (locals.user) {
		params.query = { showDeleted: 'true' };
	}
	const posts = await createRequest<Post[]>(params);
	if ('error' in posts) return error(posts.code, posts.error);

	return {
		posts,
		user: locals.user,
	};
};
