import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import { createRequest, createPaginatedRequest } from '$lib/server/api';
import { fileUploadSchema } from '$lib/models/file';
import type { Post } from '$lib/models/post';
import type { Actions, PageServerLoad } from './$types';
import { message, fail, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { error } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ url }) => {
	const page = url.searchParams.get('page') || '1';

	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/posts',
		query: { page, page_size: '10' },
	};
	const result = await createPaginatedRequest<Post>(reqParams);
	if ('error' in result) return error(result.code, result.error);

	return {
		posts: result.items,
		pagination: result.pagination,
	};
};

export const actions: Actions = {
	fileUpload: async (event) => {
		const formData = await event.request.formData();
		const form = await superValidate(formData, zod(fileUploadSchema));
		if (!form.valid) {
			return fail(400, { form });
		}

		const reqParams: RequestParams = {
			method: httpRequestEnum.enum.POST,
			path: `/files`,
			formData: formData,
			auth: event.locals.token,
		};
		const results = await createRequest<never>(reqParams);
		if ('error' in results) return fail(results.code, { form });

		return message(form, 'File uploaded successfully');
	},
};
