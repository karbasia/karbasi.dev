<script lang="ts">
	import PostCard from '$lib/components/ui/Post/PostCard.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head>
	<title>Karbasi.dev</title>
</svelte:head>

{#each data.posts as post}
	{#if !post.deleted_at && post.active === 1}
		<PostCard {post}></PostCard>
	{/if}
{/each}
