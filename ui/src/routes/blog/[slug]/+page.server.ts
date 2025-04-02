import { createRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Post } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const load: PageServerLoad = async ({ params, locals }) => {
	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: `/posts/${params.slug}`,
	};
	const post = await createRequest<Post>(reqParams);
	if ('error' in post) return error(post.code, post.error);

	// Only the admin can open the deleted posts
	if ((post.deleted_at || post.active === 0) && !locals.user) {
		return error(404, 'The requested resource could not be found');
	}
	return {
		post,
	};
};
