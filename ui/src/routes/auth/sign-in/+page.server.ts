import { fail, redirect } from '@sveltejs/kit';
import type { LoginResult } from '$lib/models/common';
import { createRequest } from '$lib/server/api';
import type { Actions } from './$types';
import { setAuth } from '$lib/server/token';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

export const actions = {
	default: async (event) => {
		const data = await event.request.formData();
		const email = data.get('email');
		const password = data.get('password');

		const params: RequestParams = {
			method: httpRequestEnum.enum.POST,
			path: '/auth/login',
			body: JSON.stringify({ email, password }),
		};

		const loginResponse = await createRequest<LoginResult>(params);
		if ('error' in loginResponse) {
			return fail(loginResponse.code, { email, incorrect: true });
		}
		setAuth(event.cookies, event.locals, loginResponse);
		redirect(303, event.url.searchParams.get('redirectTo') ?? '/admin');
	},
} satisfies Actions;
