import { fail, redirect } from '@sveltejs/kit';
import type { LoginResult } from '$lib/models/common';
import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import type { Actions } from './$types';
import { setAccessTokenCookie, setAuth, setRefreshTokenCookie } from '$lib/server/token';

export const actions = {
	default: async (event) => {
		const data = await event.request.formData();
		const email = data.get('email');
		const password = data.get('password');

		const params: RequestParams = {
			method: HttpRequest.POST,
			path: '/auth/login',
			body: JSON.stringify({ email, password }),
		};

		const loginResponse = await createRequest<LoginResult>(params);
		if ('error' in loginResponse) {
			return fail(loginResponse.code, { email, incorrect: true });
		}
		setAuth(event.cookies, event.locals, loginResponse);
		redirect(303, '/admin');
		return { success: true };
	},
} satisfies Actions;
