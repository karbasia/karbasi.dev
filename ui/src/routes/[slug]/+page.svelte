<script lang="ts">
	import hljs from 'highlight.js';
	import 'highlight.js/styles/github-dark.css';
	import { onMount } from 'svelte';

	let { data } = $props();
	onMount(() => {
		hljs.highlightAll();
	});
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<article class="prose dark:prose-invert max-w-none">
	{@html data.post.content}
</article>
