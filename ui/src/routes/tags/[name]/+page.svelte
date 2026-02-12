<script lang="ts">
	import PostCard from '$lib/components/post/PostCard.svelte';
	import Pagination from '$lib/components/layout/Pagination.svelte';
	import { page } from '$app/state';

	const { data } = $props();
</script>

<span class="flex-1">
	{#if data.posts.length === 0}
		<p>This tag does not contain any posts.</p>
	{/if}
	{#each data.posts as post}
		<PostCard {post}></PostCard>
	{/each}
	<Pagination page={data.pagination.page} totalPages={data.pagination.total_pages} baseUrl={`/tags/${page.params.name}`} />
</span>
