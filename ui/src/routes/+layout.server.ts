import type { LayoutServerLoad } from './$types';
import { createRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import type { Tag } from '$lib/models/tag';

export const load: LayoutServerLoad = async ({ locals }) => {
	const tagParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/tags/counts',
		query: { page_size: '1000' },
		auth: locals.token,
	};
	const tagPosts = await createRequest<Tag[]>(tagParams);
	if ('error' in tagPosts) return error(tagPosts.code, tagPosts.error);

	return {
		tagPosts,
		user: locals.user,
	};
};
