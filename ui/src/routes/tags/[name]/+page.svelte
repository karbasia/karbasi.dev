<script lang="ts">
	import PostCard from '$lib/components/post/PostCard.svelte';

	const { data } = $props();
</script>

<span class="flex-1">
	{#if data.posts.length === 0}
		<p>This tag does not contain any posts.</p>
	{/if}
	{#each data.posts as post}
		{#if !post.deleted_at && post.active}
			<PostCard {post}></PostCard>
		{/if}
	{/each}
</span>
