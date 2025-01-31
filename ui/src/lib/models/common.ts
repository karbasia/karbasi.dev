import type { UserCore } from './user';

export interface ErrorMessage {
	error: string;
	code: number;
}

export interface LoginResult {
	access_token: string;
	access_token_expiry: string;
	refresh_token: string;
	refresh_token_expiry: string;
	user_info: UserCore;
}

export interface RefreshTokenResult {
	access_token: string;
	access_token_expiry: string;
	user_info: UserCore;
}
