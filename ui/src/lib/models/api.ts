import { z } from 'zod';

export const httpRequestEnum = z.enum(['GET', 'POST', 'PUT', 'PATCH', 'DELETE']);

const requestParams = z.object({
	method: httpRequestEnum,
	path: z.string(),
	query: z.record(z.string(), z.string()).optional(),
	body: z.object({}).passthrough().optional(),
	auth: z.string().optional(),
	formData: z.instanceof(FormData).optional(),
});

export type RequestParams = z.infer<typeof requestParams>;
