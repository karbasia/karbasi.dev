import { redirect, type Handle } from '@sveltejs/kit';
import { createRequest } from '$lib/server/api';
import type { RefreshTokenResult } from '$lib/models/common';
import { clearAuth, setAuth } from '$lib/server/token';
import { sequence } from '@sveltejs/kit/hooks';
import { httpRequestEnum, type RequestParams } from '$lib/models/api';

const protectedRoutes = ['/profile', '/admin', '/auth/sign-out'];

const verifyLoggedInUser: Handle = async ({ event, resolve }) => {
	try {
		const accessToken = event.cookies.get('accessToken');
		const refreshToken = event.cookies.get('refreshToken');
		const userInfo = event.cookies.get('userInfo');
		if (accessToken && userInfo) {
			event.locals.user = JSON.parse(userInfo);
			event.locals.token = accessToken;
		} else if (!accessToken && refreshToken) {
			// Validate the Access Token
			const params: RequestParams = {
				method: httpRequestEnum.enum.POST,
				path: '/auth/refresh',
				body: JSON.stringify({ refresh_token: refreshToken }),
			};
			const res = await createRequest<RefreshTokenResult>(params);
			if ('error' in res) {
				throw new Error('Could not refresh token');
			} else {
				setAuth(event.cookies, event.locals, res);
			}
		}
	} catch (e) {
		clearAuth(event.cookies, event.locals);
		console.error(e);
	}
	return await resolve(event);
};

const protectRoutes: Handle = async ({ event, resolve }) => {
	if (protectedRoutes.some((route) => event.url.pathname.startsWith(route))) {
		const accessToken = event.cookies.get('accessToken');
		if (!accessToken) redirect(302, `/auth/sign-in/?redirectTo=${event.url.href}`);
	}
	return await resolve(event);
};

export const handle = sequence(verifyLoggedInUser, protectRoutes);
