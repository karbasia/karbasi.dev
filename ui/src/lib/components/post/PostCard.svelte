<script lang="ts">
	import {
		Card,
		CardHeader,
		CardTitle,
		CardDescription,
		CardContent,
	} from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { CalendarIcon } from 'lucide-svelte';
	import type { Post } from '$lib/models/post';
	import { formatDate } from '$lib/util/date';

	let { post }: { post: Post } = $props();
</script>

<div class="flex flex-col gap-6">
	<Card class="border-secondary">
		<CardHeader>
			<CardTitle>
				<a href={`/blog/${post.slug}`} class="hover:text-primary">
					{post.title}
				</a>
			</CardTitle>
			<CardDescription class="flex items-center gap-2">
				<CalendarIcon class="h-4 w-4" />
				<span>{formatDate(post.posted_at ?? post.created_at)}</span>
				<div class="flex flex-wrap gap-2">
					{#each post.tags as tag}
						<Badge variant="secondary">{tag.name}</Badge>
					{/each}
				</div>
			</CardDescription>
		</CardHeader>
		<CardContent>
			<p class="mb-4">{post.slug}</p>
		</CardContent>
	</Card>
</div>
