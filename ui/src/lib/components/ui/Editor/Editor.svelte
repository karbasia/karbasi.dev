<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import Heading from '@tiptap/extension-heading';
	import type { Post } from '$lib/models/post';

	let element: Element;
	let editor: Editor | undefined = $state();
	let { post, editable }: { post: Post; editable: boolean } = $props();

	onMount(() => {
		editor = new Editor({
			element: element,
			extensions: [
				StarterKit,
				// Heading.configure({
				// 	HTMLAttributes: {
				// 		class: 'text-xl font-bold',
				// 		levels: [2],
				// 	},
				// }),
			],
			content: post.content,
			// onTransaction: () => {
			// 	// force re-render so `editor.isActive` works as expected
			// 	editor = editor;
			// },
		});
	});

	onDestroy(() => {
		if (editor) {
			console.log('destroyed');
			editor.destroy();
		}
	});
</script>

<div bind:this={element}></div>

<style>
	button.active {
		background: black;
		color: white;
	}
</style>
