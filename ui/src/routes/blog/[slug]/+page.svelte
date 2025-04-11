<script lang="ts">
	import { Edra } from '$lib/components/edra/shadcn/index.js';
	import { Button } from '$lib/components/ui/button';
	import type { Editor } from '@tiptap/core';

	let editor = $state<Editor>();

	let { data } = $props();
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
	<Edra
		class="h-[90vh] w-full overflow-auto"
		bind:editor
		content={data.post.content}
		editable={false}
	/>
</div>
