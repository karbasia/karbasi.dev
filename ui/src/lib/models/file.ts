import { z } from 'zod';

export const fileUploadSchema = z.object({
	file: z.instanceof(File, { message: 'Please upload a file.' }),
});

export const fileSchema = z.object({
	id: z.number(),
	name: z.string(),
	created_at: z.string(),
	updated_at: z.string(),
	deleted_at: z.string().optional().nullable(),
});

export type FileUploadFormSchema = z.infer<typeof fileUploadSchema>;
export type FileSchema = z.infer<typeof fileSchema>;
