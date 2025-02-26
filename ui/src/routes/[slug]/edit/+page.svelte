<script lang="ts">
	import Editor from '$lib/components/ui/Editor/Editor.svelte';

	let { data, form } = $props();
	console.log(data.post.posted_at.split('T')[0]);
	data.post.id;
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div class="my-4 flex-row">
	<div>Title: <input value={data.post.title} /></div>
	<div>Slug: <input value={data.post.slug} /></div>
	<div>Posted date: <input type="date" value={data.post.posted_at.split('T')[0]} /></div>
	<div><input type="checkbox" checked={data.post.active === 1} /> Active?</div>
	<button
		class="dark:bg-primary-600 bg-primary-800 cursor-pointer rounded px-4"
		onclick={() => console.log('saved!')}>Save</button
	>
</div>
<article class="prose dark:prose-invert max-w-none">
	<Editor post={data.post} editable={true}></Editor>
</article>
