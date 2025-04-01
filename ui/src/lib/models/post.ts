import { z } from 'zod';
import { tagCoreSchema } from './tag';
import { userCoreSchema } from './user';

const postSchema = z.object({
	id: z.number(),
	title: z.string(),
	slug: z.string(),
	content: z.string().optional(),
	active: z.union([z.literal(0), z.literal(1)]),
	posted_at: z.string().datetime().nullable(),
	created_by: userCoreSchema,
	tags: z.array(tagCoreSchema),
	created_at: z.string().datetime(),
	updated_at: z.string().datetime(),
	deleted_at: z.string().datetime().nullable(),
});

export type Post = z.infer<typeof postSchema>;
