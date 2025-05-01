<script lang="ts">
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import hljs from 'highlight.js';
	import mermaid from 'mermaid';
	import { Button } from '$lib/components/ui/button';

	let { data } = $props();

	let darkMode = $state(true);

	onMount(() => {
		if (browser) {
			darkMode = document.documentElement.classList.contains('dark');
		}
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
	<article class="prose dark:prose-invert h-full min-w-full cursor-auto *:outline-none">
		{@html data.post.content}
	</article>
</div>
