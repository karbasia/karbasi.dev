import {z} from 'zod';

export const userCoreSchema = z.object({
	id: z.number(),
	full_name: z.string()
});

export type UserCore = z.infer<typeof userCoreSchema>;