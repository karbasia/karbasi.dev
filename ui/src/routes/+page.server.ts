import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import { createRequest } from '$lib/server/api';
import { fileUploadSchema } from '$lib/models/file';
import type { Actions } from '@sveltejs/kit';
import { message, fail, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';

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
