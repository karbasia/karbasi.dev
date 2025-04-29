<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import type { Editor } from '@tiptap/core';
	import type { Transaction } from '@tiptap/pm/state';
	import { Edra, EdraToolbar } from '$lib/components/edra/shadcn';
	import DragHandle from '$lib/components/edra/drag-handle.svelte';
	import { superForm } from 'sveltekit-superforms';
	import { postFormSchema } from '$lib/models/post';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { Checkbox } from '$lib/components/ui/checkbox';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Trash } from 'lucide-svelte';

	const { data } = $props();

	const form = superForm(data.form, {
		dataType: 'json',
		validators: zodClient(postFormSchema),
	});

	const { form: formData, enhance } = superForm(data.form, {
		dataType: 'json',
	});

	let editor = $state<Editor>();
	let showToolBar = $state(true);
	let tags = $derived($formData.tags);

	const onUpdate = (props: { editor: Editor; transaction: Transaction }) => {
		$formData.content = props.editor.getHTML();
	};

	const addTag = () => {
		tags.push({ id: 0, name: '' });
		$formData.tags = tags;
	};

	const removeTag = (tagIndex: number) => {
		tags.splice(tagIndex, 1);
		$formData.tags = tags;
	};
</script>

<svelte:head>
	<title>Karbasi.dev | {data.post.title}</title>
</svelte:head>

<div class="mx-auto w-full">
	<form method="POST" use:enhance>
		<Form.Field {form} name="title">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Title</Form.Label>
					<Input {...props} bind:value={$formData.title} />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="slug">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Slug</Form.Label>
					<Input {...props} bind:value={$formData.slug} />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="headline">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Headline</Form.Label>
					<Input {...props} bind:value={$formData.headline} />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="posted_at">
			<Form.Control>
				{#snippet children({ props })}
					<Form.Label>Posted Date</Form.Label>
					<Input {...props} type="date" bind:value={$formData.posted_at} />
				{/snippet}
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field
			{form}
			name="active"
			class="border-input flex flex-row items-start space-x-3 space-y-0 rounded-md border p-4"
		>
			<Form.Control>
				{#snippet children({ props })}
					<Checkbox {...props} bind:checked={$formData.active} />
					<div class="space-y-1 leading-none">
						<Form.Label>Post Active?</Form.Label>
					</div>
				{/snippet}
			</Form.Control>
		</Form.Field>
		<Form.Fieldset {form} name="tags" class="space-y-0">
			<div class="mt-4">
				<Form.Legend class="text-base">Tags</Form.Legend>
			</div>
			<div class="space-y-2">
				{#each tags as _, i}
					<Form.Field {form} name="tags[{i}]" class="flex flex-row">
						<Form.Control>
							{#snippet children({ props })}
								<Form.Label class="hidden">{$formData.tags[i].name}</Form.Label>
								<Input
									{...props}
									bind:value={$formData.tags[i].name}
									readonly={!!$formData.tags[i].id}
								/>
								<Button variant="destructive" class="ml-2" onclick={() => removeTag(i)}
									><Trash /></Button
								>
							{/snippet}
						</Form.Control>
						<Form.FieldErrors />
					</Form.Field>
				{/each}
				<Button onclick={addTag}>Add Tag</Button>
			</div>
		</Form.Fieldset>
		<Form.Field {form} name="content" class="my-2">
			{#if editor && showToolBar}
				<div class="overflow-auto rounded-t border-x border-t p-1">
					<EdraToolbar {editor} />
				</div>
				<DragHandle {editor} />
			{/if}
			<div class="rounded-b border">
				<Edra
					class="h-[50vh] w-full overflow-auto"
					bind:editor
					content={$formData.content}
					{onUpdate}
				/>
			</div>
		</Form.Field>
		<Form.Button>Save</Form.Button>
	</form>
</div>
