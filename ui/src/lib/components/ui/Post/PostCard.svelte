<script lang="ts">
	import type { Post } from '$lib/models/post';
	import Badge from '../badge/badge.svelte';
	let { post }: { post: Post } = $props();
</script>

<div class="flex flex-col pb-5">
	<h2>{post.title}</h2>
	<div>Posted at {new Date(post.posted_at).toLocaleDateString()}</div>
	<div>
		{#each post.tags || [] as tag}
			<Badge class="mx-1">{tag.name}</Badge>
		{/each}
	</div>
	<div>{post.content.length > 50 ? `${post.content.substring(0, 50)}...` : post.content}</div>
</div>
