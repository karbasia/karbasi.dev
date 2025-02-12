import { redirect, type Handle } from '@sveltejs/kit';
import type { UserCore } from '$lib/models/user';
import { createRequest, HttpRequest, type RequestParams } from '$lib/server/api';
import type { RefreshTokenResult } from '$lib/models/common';
import { clearAuth, setAccessTokenCookie } from '$lib/server/token';
import { sequence } from '@sveltejs/kit/hooks';

const protectedRoutes = ['/profile', '/admin', '/auth/sign-out'];

const verifyLoggedInUser: Handle = async ({ event, resolve }) => {
	try {
		const accessToken = event.cookies.get('accessToken');
		if (accessToken) {
			// Validate the Access Token
			const params: RequestParams = {
				method: HttpRequest.GET,
				path: '/users/me',
				auth: accessToken,
			};
			const data = await createRequest<UserCore>(params);
			if ('error' in data) {
				// Access Token is invalid. Try using the refresh token to get a new one
				const refreshToken = event.cookies.get('refreshToken');
				if (refreshToken) {
					const refreshParams: RequestParams = {
						method: HttpRequest.POST,
						path: '/auth/refresh',
						body: JSON.stringify({ refresh_token: refreshToken }),
					};
					const refreshData = await createRequest<RefreshTokenResult>(refreshParams);
					if ('error' in refreshData) {
						// Refresh token is invalid. Clear out all cookies
						clearAuth(event.cookies, event.locals);
					} else {
						setAccessTokenCookie(event.cookies, refreshData.access_token);
						event.locals.user = refreshData.user_info;
					}
				}
			} else {
				event.locals.user = data;
			}
		}
	} catch (e) {
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
