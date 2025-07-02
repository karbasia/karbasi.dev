<script lang="ts">
	import { Badge } from '$lib/components/ui/badge';
	import { Card, CardHeader, CardTitle, CardContent } from '$lib/components/ui/card';
	import type { Tag } from '$lib/models/tag';
	import { TagIcon } from 'lucide-svelte';

	const {
		tags,
	}: {
		tags: Tag[];
	} = $props();
</script>

<Card class="border-secondary">
	<CardHeader>
		<CardTitle class="flex items-center gap-2">
			<TagIcon class="h-5 w-5" />
			Tags
		</CardTitle>
	</CardHeader>
	<CardContent>
		<div class="flex flex-col gap-2">
			{#each tags as tag}
				<div class="flex items-center justify-between">
					<a href={`/tags/${tag.name}`}>
						<Badge variant="outline" class="hover:bg-secondary cursor-pointer">
							{tag.name}
						</Badge>
					</a>
					<span class="text-muted-foreground text-sm"
						>{tag.post_count} {tag.post_count! === 1 ? 'post' : 'posts'}</span
					>
				</div>
			{/each}
		</div>
	</CardContent>
</Card>
