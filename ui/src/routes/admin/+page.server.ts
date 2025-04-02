import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import { createRequest } from '$lib/server/api';
import type { Tag } from '$lib/models/tag';

export const load: PageServerLoad = async ({ locals }) => {
	const tagParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/tags?showDeleted=1',
		auth: locals.token,
	};
	const tags = await createRequest<Tag[]>(tagParams);
	if ('error' in tags) return error(tags.code, tags.error);

	return {
		tags,
	};
};
