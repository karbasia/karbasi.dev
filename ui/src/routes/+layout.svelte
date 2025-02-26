<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import { Menu, X } from 'lucide-svelte';

	let { data, children } = $props();
	let menuOpen = $state(false);

	const handleMenuClick = () => (menuOpen = !menuOpen);
	const handleMenuClose = () => (menuOpen = false);

	const navItems = [
		{ name: 'About', path: '/about', adminOnly: false },
		{ name: 'Projects', path: '/projects', adminOnly: false },
		{ name: 'Admin', path: '/admin', adminOnly: true },
	];
</script>

<header class="fixed top-0 w-full">
	<nav class="container mx-auto py-3">
		<div class="flex items-center justify-between">
			<a href="/" class="text-2xl font-bold">karbasi.dev</a>
			<div class="hidden space-x-4 md:block">
				{#each navItems as navItem}
					{#if !navItem.adminOnly || (data.user && navItem.adminOnly)}
						<a
							href={navItem.path}
							class={page.url.pathname === navItem.path ? 'font-bold' : 'font-semibold'}
							>{navItem.name}
						</a>
					{/if}
				{/each}
			</div>
			<div class="md:hidden">
				<button onclick={handleMenuClick}><Menu /></button>
			</div>
		</div>
		<div
			class="bg-etch-200 dark:bg-surface-800 fixed inset-y-0 right-0 w-48 transform flex-col transition-transform duration-300 ease-in-out
			{menuOpen ? '' : 'translate-x-full'}"
		>
			<div class="flex flex-row justify-end py-3">
				<button onclick={handleMenuClose}>
					<X />
				</button>
			</div>
			<div class="flex flex-col">
				{#each navItems as navItem}
					{#if !navItem.adminOnly || (data.user && navItem.adminOnly)}
						<a
							href={navItem.path}
							class="px-6 py-2 {page.url.pathname === navItem.path ? 'font-bold' : 'font-semibold'}"
							>{navItem.name}
						</a>
					{/if}
				{/each}
			</div>
		</div>
	</nav>
</header>
<div class="container mx-auto mt-14 flex flex-col justify-center">
	{@render children()}
</div>
