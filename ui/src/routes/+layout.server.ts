import type { LayoutServerLoad } from './$types';
import type { Post } from '$lib/models/post';
import { createRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import type { Tag } from '$lib/models/tag';

export const load: LayoutServerLoad = async ({ locals }) => {
	const postParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/posts',
		auth: locals.token,
	};
	if (locals.user) {
		postParams.query = { showDeleted: 'true' };
	}
	const posts = await createRequest<Post[]>(postParams);
	if ('error' in posts) return error(posts.code, posts.error);

	const tagParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/tags/counts',
		auth: locals.token,
	};
	const tagPosts = await createRequest<Tag[]>(tagParams);
	if ('error' in tagPosts) return error(tagPosts.code, tagPosts.error);

	return {
		posts,
		tagPosts,
		user: locals.user,
	};
};
