<script lang="ts">
	import { onMount } from 'svelte';
	import {
		AllCommunityModule,
		ModuleRegistry,
		createGrid,
		type GridApi,
		type GridOptions,
		type RowClickedEvent,
	} from 'ag-grid-community';

	import { mode } from 'mode-watcher';

	ModuleRegistry.registerModules([AllCommunityModule]);
	type T = $$Generic;
	let element: HTMLElement;

	let {
		defaultGridOptions,
		rowData,
		onRowClicked,
		api = $bindable(),
	}: {
		defaultGridOptions: GridOptions;
		rowData: T[];
		onRowClicked?: (event: RowClickedEvent<any, any>) => void;
		api?: GridApi;
	} = $props();

	onMount(() => {
		api = createGrid(element, {
			...defaultGridOptions,
			onRowClicked,
			rowClass: onRowClicked ? 'cursor-pointer' : 'cursor-auto',
		});
		api.setGridOption('rowData', rowData);
	});

	$effect(() => {
		if (api) {
			api.setGridOption('rowData', rowData);
		}
	});

	$effect(() => {
		document.body.dataset.agThemeMode = $mode;
	});
</script>

<div class="h-full" bind:this={element}></div>
