<script lang="ts">
	import { Editor } from '@tiptap/core';
	import type { Level } from '@tiptap/extension-heading';
	import Badge from '../Badge/Badge.svelte';

	let {
		editor,
	}: {
		editor: Editor | undefined;
	} = $props();
</script>

<div class="border-b p-2">
	<div class="flex flex-wrap gap-2">
		<!-- Headings -->
		<div class="flex gap-1">
			{#each Array.from({ length: 6 }) as _, i}
				{@const level = (i + 1) as Level}
				<Badge
					variant={editor?.isActive('heading', { level }) ? 'secondary' : 'outline'}
					onclick={() => editor?.chain().focus().toggleHeading({ level }).run()}
					>H{level}
				</Badge>
			{/each}
		</div>

		|

		<!-- Text styles -->
		<div class="flex gap-1">
			<Badge
				variant={editor?.isActive('bold') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleBold().run()}
			>
				Bold
			</Badge>

			<Badge
				variant={editor?.isActive('italic') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleItalic().run()}
			>
				Italics
			</Badge>

			<Badge
				class="line-through"
				variant={editor?.isActive('strike') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleStrike().run()}
			>
				Strikethrough
			</Badge>
		</div>

		|

		<!-- Code -->
		<div class="flex gap-1">
			<Badge
				variant={editor?.isActive('code') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleCode().run()}
			>
				Code snippet
			</Badge>

			<Badge
				variant={editor?.isActive('codeBlock') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleCodeBlock().run()}
			>
				Code block
			</Badge>
		</div>

		|

		<!-- Lists -->
		<div class="flex gap-1">
			<Badge
				variant={editor?.isActive('bulletList') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleBulletList().run()}
			>
				Bullet list
			</Badge>

			<Badge
				variant={editor?.isActive('orderedList') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleOrderedList().run()}
			>
				Ordered list
			</Badge>
		</div>

		|

		<!-- Block quotes and horizontal rule -->
		<div class="flex gap-1">
			<Badge
				variant={editor?.isActive('blockquote') ? 'secondary' : 'outline'}
				onclick={() => editor?.chain().focus().toggleBlockquote().run()}
			>
				Block quote
			</Badge>

			<Badge variant="outline" onclick={() => editor?.chain().focus().setHorizontalRule().run()}>
				Horizontal Rule
			</Badge>
		</div>

		|

		<!-- Undo/Redo -->
		<div class="flex gap-1">
			<Badge
				variant={!editor?.can().undo() ? 'destructive' : 'outline'}
				onclick={() => editor?.chain().focus().undo().run()}
			>
				Undo
			</Badge>

			<Badge
				variant={!editor?.can().redo() ? 'destructive' : 'outline'}
				onclick={() => editor?.chain().focus().redo().run()}
			>
				Redo
			</Badge>
		</div>
	</div>
</div>
