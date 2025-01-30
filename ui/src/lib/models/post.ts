import type { Tag } from './tag';
import type { CoreUser } from './user';

export interface Post {
	id: number;
	title: string;
	slug: string;
	content: string;
	active: 0 | 1;
	posted_at: Date;
	created_by: CoreUser;
	tags?: Tag[];
	created_at: Date;
	updated_at: Date;
	deleted_at?: Date;
}
