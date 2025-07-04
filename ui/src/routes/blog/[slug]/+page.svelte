<script lang="ts">
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';
	import hljs from 'highlight.js';
	import mermaid from 'mermaid';
	import { Button } from '$lib/components/ui/button';
	import { Badge } from '$lib/components/ui/badge';
	import { formatDate } from '$lib/util/date';

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
	<div class="mb-2 text-center text-3xl">{data.post.title}</div>
	<div class="mb-4 flex flex-col items-center justify-center border-b border-secondary pb-2">
		<span class="mb-2 text-center">{formatDate(data.post.posted_at ?? data.post.created_at)}</span>
		<div class="flex flex-wrap justify-center gap-2">
			{#each data.post.tags as tag}
				<a href={`/tags/${tag.name}`}><Badge variant="outline">{tag.name}</Badge></a>
			{/each}
		</div>
	</div>
	{#if data.user}
		<span class="mb-4 flex"
			><a href={`/blog/${data.post.slug}/edit`}><Button variant="outline">Edit</Button></a></span
		>
	{/if}
	<article class="prose h-full min-w-full cursor-auto dark:prose-invert *:outline-none">
		{@html data.post.content}
	</article>
</div>
