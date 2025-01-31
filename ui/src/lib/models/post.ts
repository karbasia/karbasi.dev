import type { Tag } from './tag';
import type { UserCore } from './user';

export interface Post {
	id: number;
	title: string;
	slug: string;
	content: string;
	active: 0 | 1;
	posted_at: string;
	created_by: UserCore;
	tags: Tag[] | null;
	created_at: string;
	updated_at: string;
	deleted_at: string | null;
}
