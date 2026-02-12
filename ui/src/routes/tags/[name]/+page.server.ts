import { error } from '@sveltejs/kit';

import type { PageServerLoad } from './$types';
import { createPaginatedRequest } from '$lib/server/api';
import { type Post } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const load: PageServerLoad = async ({ params, url }) => {
	const page = url.searchParams.get('page') || '1';

	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: `/tags/${params.name}`,
		query: { page, page_size: '10' },
	};
	const result = await createPaginatedRequest<Post>(reqParams);
	if ('error' in result) return error(result.code, result.error);

	return {
		posts: result.items,
		pagination: result.pagination,
	};
};
