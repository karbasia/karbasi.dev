import { z } from 'zod';
import { userCoreSchema } from './user';

const errorMessageSchema = z.object({
	error: z.string(),
	code: z.number(),
});

const loginResultSchema = z.object({
	access_token: z.string(),
	access_token_expiry: z.string().datetime(),
	refresh_token: z.string(),
	refresh_token_expiry: z.string().datetime(),
	user_info: userCoreSchema,
});

const refreshTokenResultSchema = loginResultSchema.omit({
	refresh_token: true,
	refresh_token_expiry: true,
});

export type ErrorMessage = z.infer<typeof errorMessageSchema>;
export type LoginResult = z.infer<typeof loginResultSchema>;
export type RefreshTokenResult = z.infer<typeof refreshTokenResultSchema>;
