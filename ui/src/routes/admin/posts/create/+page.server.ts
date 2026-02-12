import { superValidate, fail } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { error, redirect } from '@sveltejs/kit';

import type { PageServerLoad, Actions } from './$types';
import { createRequest } from '$lib/server/api';
import { type Post, postCreateSchema } from '$lib/models/post';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import type { Tag } from '$lib/models/tag';

export const load: PageServerLoad = async ({ locals }) => {
	if (!locals.user) {
		return error(403, 'Only the admin can create posts');
	}

	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/tags',
	};
	const tags = await createRequest<Tag[]>(reqParams);
	if ('error' in tags) return error(tags.code, tags.error);

	return {
		tags,
		form: await superValidate(zod(postCreateSchema)),
	};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, zod(postCreateSchema));
		if (!form.valid) {
			return fail(400, {
				form,
			});
		}
		const params: RequestParams = {
			method: httpRequestEnum.enum.POST,
			path: '/posts',
			body: form.data,
			auth: event.locals.token,
		};
		const postData = await createRequest<Post>(params);
		if ('error' in postData) return fail(postData.code);
		redirect(303, '/admin');
	},
};
