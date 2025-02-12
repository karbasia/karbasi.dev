<script lang="ts">
	import Editor from '$lib/components/ui/Editor/Editor.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div>{@html data.post.content}</div>

<Editor post={data.post} editable={false}></Editor>
