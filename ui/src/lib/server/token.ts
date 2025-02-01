import type { LoginResult } from '$lib/models/common';
import { type Cookies } from '@sveltejs/kit';

export function setAccessTokenCookie(cookies: Cookies, accessToken: string, expiry: Date): void {
	cookies.set('accessToken', accessToken, {
		httpOnly: true,
		path: '/',
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'strict',
		expires: expiry,
	});
}

export function setRefreshTokenCookie(cookies: Cookies, refreshToken: string, expiry: Date): void {
	cookies.set('refreshToken', refreshToken, {
		httpOnly: true,
		path: '/',
		secure: process.env.NODE_ENV === 'production',
		sameSite: 'strict',
		expires: expiry,
	});
}

export function setAuth(cookies: Cookies, locals: App.Locals, loginResponse: LoginResult): void {
	locals.user = loginResponse.user_info;
	setAccessTokenCookie(
		cookies,
		loginResponse.access_token,
		new Date(loginResponse.access_token_expiry),
	);
	setRefreshTokenCookie(
		cookies,
		loginResponse.refresh_token,
		new Date(loginResponse.refresh_token_expiry),
	);
}

export function clearAuth(cookies: Cookies, locals: App.Locals): void {
	locals.user = null;
	cookies.delete('accessToken', { path: '/' });
	cookies.delete('refreshToken', { path: '/' });
}
