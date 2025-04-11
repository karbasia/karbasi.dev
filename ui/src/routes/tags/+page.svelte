<script lang="ts">
	import TagSidebar from '$lib/components/layout/TagSidebar.svelte';

	const { data } = $props();
</script>

<aside class="w-full">
	<TagSidebar tags={data.tagPosts} />
</aside>
