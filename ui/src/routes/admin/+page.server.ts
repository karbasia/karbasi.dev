import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import { createRequest } from '$lib/server/api';
import type { Tag } from '$lib/models/tag';
import { type FileSchema, fileUploadSchema } from '$lib/models/file';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

export const load: PageServerLoad = async ({ locals }) => {
	const tagParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/tags?showDeleted=1',
		auth: locals.token,
	};
	const tags = await createRequest<Tag[]>(tagParams);
	if ('error' in tags) return error(tags.code, tags.error);

	const fileParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: '/files',
		auth: locals.token,
	};
	const files = await createRequest<FileSchema[]>(fileParams);
	if ('error' in files) return error(files.code, files.error);

	return {
		tags,
		files,
		uploadForm: await superValidate(zod(fileUploadSchema)),
	};
};
