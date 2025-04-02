<script lang="ts">
	import PostsGrid from '$lib/components/post/PostsGrid.svelte';
	import * as Tabs from '$lib/components/ui/tabs/index.js';

	const { data } = $props();
</script>

<Tabs.Root value="posts" class="w-full">
	<Tabs.List class="grid w-full grid-cols-2">
		<Tabs.Trigger value="posts">Posts</Tabs.Trigger>
		<Tabs.Trigger value="tags">Tags</Tabs.Trigger>
	</Tabs.List>
	<Tabs.Content value="posts">
		<PostsGrid posts={data.posts} />
	</Tabs.Content>
	<Tabs.Content value="tags">
		<p>tags</p>
	</Tabs.Content>
</Tabs.Root>
