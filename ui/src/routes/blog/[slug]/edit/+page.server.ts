import { superValidate, fail } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { error } from '@sveltejs/kit';

import type { PageServerLoad, Actions } from './$types';
import { createRequest } from '$lib/server/api';
import { type Post, postFormSchema, postSchema } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import type { Tag } from '$lib/models/tag';

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

	reqParams.path = '/tags';
	const tags = await createRequest<Tag[]>(reqParams);
	if ('error' in tags) return error(tags.code, tags.error);

	return {
		post,
		tags,
		form: await superValidate(post, zod(postSchema)),
	};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, zod(postFormSchema));
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}
		const params: RequestParams = {
			method: httpRequestEnum.enum.PATCH,
			path: `/posts/${form.data.id}`,
			body: form.data,
			auth: event.locals.token,
		};
		const postData = await createRequest<Post>(params);
		if ('error' in postData) return fail(postData.code);
		form.data = postData;
		return {
			form,
		};
	},
};
