<script lang="ts">
	import PostCard from '$lib/components/post/PostCard.svelte';

  export let data;
</script>

<svelte:head>
	<title>Karbasi.dev</title>
</svelte:head>


{#each data.posts as post}
  <PostCard post={post}></PostCard>
{/each}