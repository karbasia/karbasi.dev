<script lang="ts">
	import Editor from '$lib/components/ui/Editor/Editor.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<article class="prose dark:prose-invert max-w-none">
	{@html data.post.content}
	<Editor post={data.post} editable={false}></Editor>
</article>
