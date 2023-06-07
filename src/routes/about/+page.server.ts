import { error } from '@sveltejs/kit';
import { pb } from '$lib/server/pb';
import { ProfileType, type ProfileRecord } from '$lib/server/pb.d';

export async function load() {
	const profileItems: ProfileRecord[] = await pb
		.collection('profiles')
		.getFullList({ 
      sort: '-to, responsibilities.order',
			expand: 'responsibilities',
      filter: `type = "work" || type = "education"` 
    });

	if (!profileItems) {
		throw error(400, 'Something went wrong');
	}

	profileItems.filter(item => item.type === ProfileType.Work).map((item) => {
		item.expand.responsibilities.sort((a, b) => a.order - b.order);
	});

	return {
		work: profileItems.filter(item => item.type === ProfileType.Work).map(item => structuredClone(item)),
		education: profileItems.filter(item => item.type === ProfileType.Education).map(item => structuredClone(item))
	};
}