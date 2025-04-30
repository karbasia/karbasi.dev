<script lang="ts">
	import mermaid from 'mermaid';
	import { Button } from '$lib/components/ui/button';
	import { onMount } from 'svelte';
	import hljs from 'highlight.js';
	let { data } = $props();

	onMount(() => {
		if (data.post.content) {
			if (data.post.content.indexOf('</pre>') > 0) {
				hljs.highlightAll();
			}
			if (data.post.content.indexOf('class="language-mermaid"') > 0) {
				mermaid.initialize({
					startOnLoad: false,
					theme: 'neutral',
				});
				mermaid.run({
					querySelector: 'code.language-mermaid',
				});
			}
		}
	});
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div class="mx-auto w-full">
	{#if data.user}
		<span class="mb-4 flex"
			><a href={`/blog/${data.post.slug}/edit`}><Button variant="outline">Edit</Button></a></span
		>
	{/if}
	{@html data.post.content}
</div>
