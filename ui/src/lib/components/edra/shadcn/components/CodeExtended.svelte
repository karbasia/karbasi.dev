<script lang="ts">
	import { NodeViewWrapper, NodeViewContent } from 'svelte-tiptap';
	import type { NodeViewProps } from '@tiptap/core';
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	const { editor, node, updateAttributes, extension }: NodeViewProps = $props();
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
	import { Check, Copy, ChevronDown } from 'lucide-svelte';

	let preRef: HTMLPreElement;

	let isCopying = $state(false);

	const languages: string[] = extension.options.lowlight.listLanguages().sort();

	let defaultLanguage = $state(node.attrs.language);

	function copyCode() {
		isCopying = true;
		navigator.clipboard.writeText(preRef.innerText);
		setTimeout(() => {
			isCopying = false;
		}, 1000);
	}
</script>

<NodeViewWrapper
	class="code-wrapper bg-muted dark:bg-muted/20 group relative rounded p-6"
	draggable={false}
	spellcheck={false}
>
	<div class="code-wrapper-tile" contenteditable="false">
		{#if editor.isEditable}
			<DropdownMenu.Root>
				<DropdownMenu.Trigger
					contenteditable="false"
					class={buttonVariants({
						variant: 'ghost',
						size: 'sm',
						class: 'text-muted-foreground h-4 rounded px-1 py-2 text-xs',
					})}
					>{defaultLanguage}
					<ChevronDown class="!size-2" />
				</DropdownMenu.Trigger>
				<DropdownMenu.Content class="h-96 w-40 overflow-auto" contenteditable="false">
					{#each languages as language}
						<DropdownMenu.Item
							contenteditable="false"
							data-current={defaultLanguage === language}
							class="data-[current=true]:bg-muted"
							textValue={language}
							onclick={() => {
								defaultLanguage = language;
								updateAttributes({ language: defaultLanguage });
							}}
						>
							<span>{language}</span>
							{#if defaultLanguage === language}
								<Check class="ml-auto" />
							{/if}
						</DropdownMenu.Item>
					{/each}
				</DropdownMenu.Content>
			</DropdownMenu.Root>
		{:else}
			<span class="text-muted-foreground inline-flex h-4 rounded px-1 pb-2 text-xs">
				{defaultLanguage}
			</span>
		{/if}
		<Button variant="ghost" class="text-muted-foreground size-4 py-0" onclick={copyCode}>
			{#if isCopying}
				<Check class="size-3 text-green-500" />
			{:else}
				<Copy class="size-3" />
			{/if}
		</Button>
	</div>
	<pre bind:this={preRef} draggable={false}>
		<NodeViewContent as="code" class={`language-${defaultLanguage}`} {...node.attrs} />	
	</pre>
</NodeViewWrapper>
