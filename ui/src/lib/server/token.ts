import type { LoginResult } from '$lib/models/common';
import { type Cookies } from '@sveltejs/kit';

export const setAccessTokenCookie = (cookies: Cookies, accessToken: string): void => {
	cookies.set('accessToken', accessToken, {
		httpOnly: true,
		path: '/',
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'strict',
	});
};

export const setRefreshTokenCookie = (cookies: Cookies, refreshToken: string): void => {
	cookies.set('refreshToken', refreshToken, {
		httpOnly: true,
		path: '/',
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'strict',
	});
};

export const setAuth = (cookies: Cookies, locals: App.Locals, loginResponse: LoginResult): void => {
	locals.user = loginResponse.user_info;
	setAccessTokenCookie(cookies, loginResponse.access_token);
	setRefreshTokenCookie(cookies, loginResponse.refresh_token);
};

export const clearAuth = (cookies: Cookies, locals: App.Locals): void => {
	locals.user = null;
	cookies.delete('accessToken', { path: '/' });
	cookies.delete('refreshToken', { path: '/' });
};
