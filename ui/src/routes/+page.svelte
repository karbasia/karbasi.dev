<script lang="ts">
	import TagSidebar from '$lib/components/layout/TagSidebar.svelte';
	import PostCard from '$lib/components/post/PostCard.svelte';

	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>Karbasi.dev</title>
</svelte:head>

<section class="flex-1">
	{#each data.posts as post}
		{#if !post.deleted_at && post.active}
			<PostCard {post}></PostCard>
		{/if}
	{/each}
</section>

<aside class="ml-4 hidden w-64 sm:block">
	<TagSidebar tags={data.tagPosts} />
</aside>
