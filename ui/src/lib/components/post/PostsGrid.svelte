<script lang="ts">
	import type { GridOptions, RowClickedEvent } from 'ag-grid-community';
	import type { Post } from '$lib/models/post';
	import AgGrid from '../grid/AGGrid.svelte';
	import { goto } from '$app/navigation';

	const {
		posts,
	}: {
		posts: Post[];
	} = $props();

	const defaultGridOptions: GridOptions<Post> = {
		defaultColDef: {
			filter: true,
		},
		columnDefs: [
			{ field: 'id', headerName: 'ID', maxWidth: 100 },
			{ field: 'title', headerName: 'Title', flex: 1 },
			{
				headerName: 'Posted Date',
				valueGetter: (r) => (r.data?.posted_at ? new Date(r.data?.posted_at) : null),
				cellDataType: 'date',
			},
			{
				headerName: 'Active?',
				valueGetter: (r) => r.data?.active === 1,
				cellDataType: 'boolean',
				maxWidth: 100,
			},
			{
				headerName: 'Deleted Date',
				valueGetter: (r) => (r.data?.deleted_at ? new Date(r.data?.deleted_at) : null),
				cellDataType: 'date',
			},
		],
	};

	const handleRowClicked = (e: RowClickedEvent<Post, any>) => {
		goto(`/blog/${e.data?.slug}`);
	};
</script>

<div class="h-[90vh]">
	<AgGrid rowData={posts} {defaultGridOptions} onRowClicked={handleRowClicked} />
</div>
