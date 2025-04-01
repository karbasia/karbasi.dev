import type { LoginResult, RefreshTokenResult } from '$lib/models/common';
import { type Cookies } from '@sveltejs/kit';

export const setCookie = (cookies: Cookies, name: string, token: string, expiry?: Date): void => {
	cookies.set(name, token, {
		httpOnly: true,
		path: '/',
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'strict',
		expires: expiry,
	});
};

export const setAuth = (
	cookies: Cookies,
	locals: App.Locals,
	loginResponse: LoginResult | RefreshTokenResult,
): void => {
	locals.user = loginResponse.user_info;
	setCookie(
		cookies,
		'accessToken',
		loginResponse.access_token,
		new Date(loginResponse.access_token_expiry),
	);
	setCookie(cookies, 'userInfo', JSON.stringify(loginResponse.user_info));
	if ('refresh_token' in loginResponse) {
		setCookie(
			cookies,
			'refreshToken',
			loginResponse.refresh_token,
			new Date(loginResponse.refresh_token_expiry),
		);
	}
};

export const clearAuth = (cookies: Cookies, locals: App.Locals): void => {
	locals.user = null;
	cookies.delete('accessToken', { path: '/' });
	cookies.delete('refreshToken', { path: '/' });
	cookies.delete('userInfo', { path: '/' });
};
