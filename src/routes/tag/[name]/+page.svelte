<script lang="ts">
  import PostCard from '$lib/components/post/PostCard.svelte';

  export let data;
</script>

{#each data.posts as post}
  <PostCard post={post}></PostCard>
{/each}