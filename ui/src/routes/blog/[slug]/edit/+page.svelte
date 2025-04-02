<script lang="ts">
	import type { Content, Editor } from '@tiptap/core';
	import type { Transaction } from '@tiptap/pm/state';
	import { Edra, EdraToolbar, EdraBubbleMenu } from '$lib/components/edra/shadcn/index.js';
	import DragHandle from '$lib/components/edra/drag-handle.svelte';

	let content = $state<Content>();
	let editor = $state<Editor>();
	let showToolBar = $state(true);

	const onUpdate = (props: { editor: Editor; transaction: Transaction }) => {
		content = props.editor.getJSON();
	};
	let { data, form } = $props();
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div class="mx-auto w-full">
	{#if editor && showToolBar}
		<div class="overflow-auto rounded-t border-x border-t p-1">
			<EdraToolbar {editor} />
		</div>
		<EdraBubbleMenu {editor} />
		<DragHandle {editor} />
	{/if}
	<div class="rounded-b border">
		<Edra
			class="h-[90vh] w-full overflow-auto"
			bind:editor
			content={data.post.content}
			{onUpdate}
		/>
	</div>
</div>
