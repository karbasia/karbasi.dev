<script lang="ts">
	import { goto } from '$app/navigation';
	import type { GridOptions, RowClickedEvent } from 'ag-grid-community';
	import type { Tag } from '$lib/models/tag';
	import AgGrid from '../grid/AGGrid.svelte';

	const {
		tags,
	}: {
		tags: Tag[];
	} = $props();

	const defaultGridOptions: GridOptions<Tag> = {
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

	const handleRowClicked = (e: RowClickedEvent<Tag, any>) => {
		goto(`/tag/${e.data?.id}`);
	};
</script>

<div class="h-[90vh]">
	<AgGrid rowData={tags} {defaultGridOptions} onRowClicked={handleRowClicked} />
</div>
