<script lang="ts">
	import type { GridOptions, RowClickedEvent } from 'ag-grid-community';
	import AgGrid from '../grid/AGGrid.svelte';
	import type { FileSchema } from '$lib/models/file';
	import { goto } from '$app/navigation';

	const {
		files,
	}: {
		files: FileSchema[];
	} = $props();

	const defaultGridOptions: GridOptions<FileSchema> = {
		defaultColDef: {
			filter: true,
		},
		columnDefs: [
			{ field: 'id', headerName: 'ID', maxWidth: 100 },
			{ field: 'name', headerName: 'Name', flex: 1 },
			{
				headerName: 'Created Date',
				valueGetter: (r) => (r.data?.created_at ? new Date(r.data?.created_at) : null),
				cellDataType: 'date',
			},
			{
				headerName: 'Deleted Date',
				valueGetter: (r) => (r.data?.deleted_at ? new Date(r.data?.deleted_at) : null),
				cellDataType: 'date',
			},
		],
	};

	const handleRowClicked = (e: RowClickedEvent<FileSchema, any>) => {
		goto(`/files/${e.data?.name}`);
	};
</script>

<div class="h-[calc(90vh-80px)]">
	<AgGrid rowData={files} {defaultGridOptions} onRowClicked={handleRowClicked} />
</div>
