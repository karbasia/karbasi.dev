<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import type { Post } from '$lib/models/post';

	let element: Element;
	let editor: Editor | undefined = $state();
	let { post, editable }: { post: Post; editable: boolean } = $props();

	onMount(() => {
		editor = new Editor({
			element: element,
			extensions: [StarterKit],
			content: post.content,
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
