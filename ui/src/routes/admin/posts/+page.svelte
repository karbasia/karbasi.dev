<script lang="ts">
	const { data } = $props();
</script>

<div class="flex">
	<table class="table-auto">
		<thead>
			<tr>
				<th class="px-6 py-4 text-left">ID</th>
				<th class="px-6 py-4 text-left">Title</th>
				<th class="px-6 py-4 text-left">Created Date</th>
				<th class="px-6 py-4 text-left">Updated Date</th>
				<th class="px-6 py-4 text-left">Deleted Date</th>
				<th class="px-6 py-4 text-left">Actions</th>
			</tr>
		</thead>
		<tbody>
			{#each data.posts || [] as post}
				<tr>
					<td class="px-6 py-2 text-left">{post.id}</td>
					<td class="px-6 py-2 text-left">{post.title}</td>
					<td class="px-6 py-2 text-left">{new Date(post.created_at).toLocaleDateString()}</td>
					<td class="px-6 py-2 text-left">{new Date(post.updated_at).toLocaleDateString()}</td>
					<td class="px-6 py-2 text-left"
						>{post.deleted_at ? new Date(post.deleted_at).toLocaleDateString() : ''}</td
					>
					<td class="px-6 py-2 text-left"><a href={`/blog/${post.slug}`}>Open</a></td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
