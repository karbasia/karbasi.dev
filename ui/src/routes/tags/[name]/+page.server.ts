import { error } from '@sveltejs/kit';

import type { PageServerLoad } from './$types';
import { createRequest } from '$lib/server/api';
import { type Post } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const load: PageServerLoad = async ({ params }) => {
	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: `/tags/${params.name}`,
	};
	const posts = await createRequest<Post[]>(reqParams);
	if ('error' in posts) return error(posts.code, posts.error);

	return {
		posts,
	};
};
