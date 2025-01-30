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

export function removeAuth(cookies: Cookies): void {
	cookies.delete('accessToken', { path: '/' });
	cookies.delete('refreshToken', { path: '/' });
}
