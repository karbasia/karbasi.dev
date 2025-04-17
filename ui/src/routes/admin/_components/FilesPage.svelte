<script lang="ts">
	import FilesGrid from '$lib/components/file/FilesGrid.svelte';
	import FileUploadForm from '$lib/components/file/FileUploadForm.svelte';
	import type { FileSchema, FileUploadFormSchema } from '$lib/models/file';
	import type { SuperValidated } from 'sveltekit-superforms';

	let {
		files,
		uploadForm,
	}: {
		files: FileSchema[];
		uploadForm: SuperValidated<FileUploadFormSchema>;
	} = $props();
</script>

<FileUploadForm {uploadForm} />
<FilesGrid {files} />
