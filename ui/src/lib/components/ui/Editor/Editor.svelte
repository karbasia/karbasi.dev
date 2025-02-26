<script lang="ts">
	import 'highlight.js/styles/github-dark.css';
	import { onMount, onDestroy } from 'svelte';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight';
	import { all, createLowlight } from 'lowlight';

	import type { Post } from '$lib/models/post';
	import Toolbar from './Toolbar.svelte';

	let element: Element;
	let editor: Editor | undefined = $state();
	let { post, editable }: { post: Post; editable: boolean } = $props();
	const lowlight = createLowlight(all);

	onMount(() => {
		editor = new Editor({
			element: element,
			extensions: [
				StarterKit,
				CodeBlockLowlight.configure({
					lowlight,
				}),
			],
			content: post.content,
			editable: editable,
			onTransaction: ({ editor: newEditor }) => {
				// force re-render so `editor.isActive` works as expected
				editor = undefined;
				editor = newEditor;
			},
		});
	});

	onDestroy(() => {
		if (editor) {
			editor.destroy();
		}
	});
</script>

<div>
	{#if editable}
		<Toolbar {editor} />
	{/if}
	<div bind:this={element}></div>
</div>
