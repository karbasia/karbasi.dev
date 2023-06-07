import { error } from '@sveltejs/kit';
import { pb } from '$lib/server/pb';
import type { PostRecord } from '$lib/server/pb.d';
import type { ListResult } from 'pocketbase';

export async function load({ params }) {
	const posts: ListResult<PostRecord> = await pb
		.collection('posts')
		.getList(1, 10, 
			{
				sort: '-created', 
				expand: 'tags' 
			});

	if (!posts) {
		throw error(400, 'Something went wrong');
	}

	return {
		posts: structuredClone(posts.items)
	};
}
