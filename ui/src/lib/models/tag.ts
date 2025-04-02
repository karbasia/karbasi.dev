import { z } from 'zod';

export const tagCoreSchema = z.object({
	id: z.number(),
	name: z.string(),
});

const tagSchema = tagCoreSchema.extend({
	post_count: z.number().optional(),
	created_at: z.string().datetime(),
	updated_at: z.string().datetime(),
	deleted_at: z.string().datetime().nullable(),
});

export type TagCore = z.infer<typeof tagCoreSchema>;
export type Tag = z.infer<typeof tagSchema>;
