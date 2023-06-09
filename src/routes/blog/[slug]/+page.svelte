<script lang="ts">
		import { onMount } from 'svelte';
	import { formatDate } from '$lib/dateUtil';
	import { afterNavigate } from '$app/navigation';
	import hljs from 'highlight.js';

	export let data;
	let previousPage : string = '/' ;

	afterNavigate(({ from }) => {
		if (from && from.url) {
			previousPage = from.url.pathname + from.url.search;	
		}
	});

	onMount(() => {
		if (data.post.body.indexOf("<pre>") > 0) {
			hljs.highlightAll();
		}
  });
</script>

<svelte:head>
	<title>{data.post.title} - Karbasi.dev</title>
</svelte:head>


<div class="card w-full">
	<header class="card-header">
		<h2 class="sm:text-3xl text-xl">{data.post.title}</h2>
		<div class="text-secondary-600 dark:text-secondary-100">{formatDate(data.post.created)}</div>
		<div class="flex flex-row">
			{#each data.post.expand.tags as tag}
				<div class="card px-1 mr-2 variant-filled-primary"><a href="/tag/{tag.name}">{tag.name}</a></div>
			{/each}
		</div>
	</header>
	<section class="p-4">{@html data.post.body}</section>
	<footer class="card-footer text-tertiary-600 dark:text-tertiary-100"><a href="{previousPage}">Go Back</a></footer>
</div>