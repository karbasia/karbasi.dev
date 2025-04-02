<script lang="ts">
	import { Edra } from '$lib/components/edra/shadcn/index.js';
	import type { Content, Editor } from '@tiptap/core';

	let editor = $state<Editor>();

	let { data } = $props();
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div class="mx-auto w-full">
	<Edra
		class="h-[90vh] w-full overflow-auto"
		bind:editor
		content={data.post.content}
		editable={false}
	/>
</div>
