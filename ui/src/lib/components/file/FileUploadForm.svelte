<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import { fileProxy, superForm, type SuperValidated } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { fileUploadSchema, type FileUploadFormSchema } from '$lib/models/file';
	import { invalidate } from '$app/navigation';

	let {
		uploadForm,
	}: {
		uploadForm: SuperValidated<FileUploadFormSchema>;
	} = $props();

	const form = superForm(uploadForm, {
		validators: zodClient(fileUploadSchema),
		onUpdated({ form }) {
			if (form.valid) {
				invalidate('files');
			}
		},
	});

	const { form: formData, enhance } = form;
	const file = fileProxy(formData, 'file');
</script>

<form
	method="POST"
	class="grid w-full grid-cols-12 items-end gap-4"
	enctype="multipart/form-data"
	use:enhance
	action="/?/fileUpload"
>
	<Form.Field {form} name="file" class="col-span-10">
		<Form.Control>
			<Form.Label>Upload File</Form.Label>
			<Input name="file" type="file" bind:files={$file} />
		</Form.Control>
		<Form.FieldErrors />
	</Form.Field>
	<Form.Button class="col-span-2 mb-2">Upload</Form.Button>
</form>
