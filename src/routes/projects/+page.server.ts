import { error } from '@sveltejs/kit';
import { pb } from '$lib/server/pb';
import { ProfileType, type ProfileRecord } from '$lib/server/pb.d';

export async function load() {
	const projects: ProfileRecord[] = await pb
		.collection('profiles')
		.getFullList({ 
      sort: 'created', 
      filter: `type = "${ProfileType.Project}"` 
    });

	if (!projects) {
		throw error(400, 'Something went wrong');
	}

	return {
		projects: structuredClone(projects)
	};
}
