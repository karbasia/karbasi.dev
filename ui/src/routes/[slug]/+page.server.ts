import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Post } from '$lib/models/post';

export const load: PageServerLoad = async ({ params, locals }) => {
	const reqParams: RequestParams = {
		method: HttpRequest.GET,
		path: `/posts/${params.slug}`,
	};
	const data = await createRequest<Post>(reqParams);
	if ('error' in data) return error(data.code, data.error);

	// Only the admin can open the deleted posts
	if ((data.deleted_at || data.active === 0) && !locals.user) {
		return error(404, 'The requested resource could not be found');
	}
	return {
		post: data,
	};
};
