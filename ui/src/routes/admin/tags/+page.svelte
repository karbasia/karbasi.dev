<script lang="ts">
	const { data } = $props();
</script>

<div class="flex">
	<table class="table-auto">
		<thead>
			<tr>
				<th class="px-6 py-4 text-left">ID</th>
				<th class="px-6 py-4 text-left">Name</th>
				<th class="px-6 py-4 text-left">Created Date</th>
				<th class="px-6 py-4 text-left">Updated Date</th>
				<th class="px-6 py-4 text-left">Deleted Date</th>
				<th class="px-6 py-4 text-left">Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each data.tags || [] as tag}
				<tr>
					<td class="px-6 py-2 text-left">{tag.id}</td>
					<td class="px-6 py-2 text-left">{tag.name}</td>
					<td class="px-6 py-2 text-left">{new Date(tag.created_at).toLocaleDateString()}</td>
					<td class="px-6 py-2 text-left">{new Date(tag.updated_at).toLocaleDateString()}</td>
					<td class="px-6 py-2 text-left"
						>{tag.deleted_at ? new Date(tag.deleted_at).toLocaleDateString() : ''}</td
					>
					<td class="px-6 py-2 text-left">...</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
