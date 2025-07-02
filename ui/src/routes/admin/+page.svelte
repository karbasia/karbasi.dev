<script lang="ts">
	import * as Tabs from '$lib/components/ui/tabs/index.js';
	import PostsGrid from '$lib/components/post/PostsGrid.svelte';
	import TagsGrid from '$lib/components/tag/TagsGrid.svelte';
	import FilesPage from './_components/FilesPage.svelte';

	const { data } = $props();
</script>

<Tabs.Root value="posts" class="w-full">
	<Tabs.List class="grid w-full grid-cols-3">
		<Tabs.Trigger value="posts">Posts</Tabs.Trigger>
		<Tabs.Trigger value="tags">Tags</Tabs.Trigger>
		<Tabs.Trigger value="files">Files</Tabs.Trigger>
	</Tabs.List>
	<Tabs.Content value="posts">
		<PostsGrid posts={data.posts} />
	</Tabs.Content>
	<Tabs.Content value="tags">
		<TagsGrid tags={data.tags} />
	</Tabs.Content>
	<Tabs.Content value="files">
		<FilesPage files={data.files} uploadForm={data.uploadForm} />
	</Tabs.Content>
</Tabs.Root>
