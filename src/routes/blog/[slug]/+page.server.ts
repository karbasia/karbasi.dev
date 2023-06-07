import { error } from '@sveltejs/kit';
import { pb } from '$lib/server/pb';
import type { PostRecord } from '$lib/server/pb.d';
import type { ListResult } from 'pocketbase';

export async function load({ params }) {
	const post: ListResult<PostRecord> = await pb
		.collection('posts')
		.getList(1, 1, { filter: `slug = "${params.slug}"`, expand: 'tags' });
    

	if (!post || post.totalItems > 1) {
		throw error(400, 'Something went wrong');
	} else if (post.totalItems === 0) {
		throw error(404, 'Post not found');
	}

	return {
		post: structuredClone(post.items[0])
	};
}
