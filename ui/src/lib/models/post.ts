import { z } from 'zod';
import { tagCoreSchema } from './tag';
import { userCoreSchema } from './user';

export const postSchema = z.object({
	id: z.number(),
	title: z.string(),
	slug: z.string(),
	content: z.string().optional(),
	active: z.boolean(),
	posted_at: z.string().nullable().optional(),
	created_by: userCoreSchema,
	tags: z.array(tagCoreSchema),
	created_at: z.string(),
	updated_at: z.string(),
	deleted_at: z.string().optional().nullable(),
});

export const postFormSchema = postSchema.omit({
	created_by: true,
	created_at: true,
	updated_at: true,
	deleted_at: true,
});

export type Post = z.infer<typeof postSchema>;
