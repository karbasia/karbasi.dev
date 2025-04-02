import { createRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';
import type { Post } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const load: PageServerLoad = async ({ params, locals }) => {
	if (!locals.user) {
		return error(403, 'Only the admin can edit posts');
	}
	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: `/posts/${params.slug}`,
	};
	const post = await createRequest<Post>(reqParams);
	if ('error' in post) return error(post.code, post.error);

	return {
		post,
	};
};

export const actions = {
	default: async (event) => {
		const data = await event.request.formData();
		const postId = data.get('id');

		const params: RequestParams = {
			method: httpRequestEnum.enum.POST,
			path: `/posts/${postId}`,
			body: JSON.stringify(data),
		};
	},
} satisfies Actions;
