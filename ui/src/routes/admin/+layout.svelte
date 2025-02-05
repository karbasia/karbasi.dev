<script lang="ts">
	import { page } from '$app/state';
	let { children } = $props();
	const menuItems = [
		{ name: 'Posts', path: '/admin/posts' },
		{ name: 'Tags', path: '/admin/tags' },
	];
</script>

<div class="min-h-screen">
	<div class="flex flex-row pt-2">
		{#each menuItems as menuOption}
			<a href={menuOption.path} class="flex flex-1 flex-col">
				<div
					class="{page.url.pathname === menuOption.path ? 'bg-etch-50' : ''}
                flex flex-1 flex-col items-center border-2 py-2 pb-2 pl-16 text-lg
                {page.url.pathname === menuOption.path ? 'font-bold' : 'font-semibold'}"
				>
					{menuOption.name}
				</div>
			</a>
		{/each}
	</div>
	<div>
		{@render children()}
	</div>
</div>
