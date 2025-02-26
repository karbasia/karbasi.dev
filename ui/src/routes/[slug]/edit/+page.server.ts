import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';
import type { Post } from '$lib/models/post';

export const load: PageServerLoad = async ({ params, locals }) => {
	if (!locals.user) {
		return error(403, 'Only the admin can edit posts');
	}
	const reqParams: RequestParams = {
		method: HttpRequest.GET,
		path: `/posts/${params.slug}`,
	};
	const data = await createRequest<Post>(reqParams);
	if ('error' in data) return error(data.code, data.error);

	return {
		post: data,
	};
};

export const actions = {
	default: async (event) => {
		const data = await event.request.formData();
		const postId = data.get('id');

		const params: RequestParams = {
			method: HttpRequest.POST,
			path: `/posts/${postId}`,
			body: JSON.stringify(data),
		};
	},
} satisfies Actions;
