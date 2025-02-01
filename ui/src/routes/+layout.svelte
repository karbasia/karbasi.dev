<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
	import { Menu, X } from 'lucide-svelte';

	let { data, children } = $props();
	let menuOpen = $state(false);

	console.log(data);

	const handleMenuClick = () => (menuOpen = !menuOpen);
	const handleMenuClose = () => (menuOpen = false);
</script>

<header class="bg-sea-300 fixed top-0 w-full">
	<nav class="container mx-auto py-3">
		<div class="flex items-center justify-between">
			<a href="/" class="text-2xl font-bold">karbasi.dev</a>
			<div class="hidden space-x-4 md:block">
				<a href="/about" class={page.url.pathname === '/about' ? 'font-bold' : 'font-semibold'}
					>About
				</a>
				<a
					href="/projects"
					class={page.url.pathname === '/projects' ? 'font-bold' : 'font-semibold'}
					>Projects
				</a>
				{#if data.user}
					<a href="/admin" class={page.url.pathname === '/admin' ? 'font-bold' : 'font-semibold'}
						>Admin
					</a>
				{/if}
			</div>
			<div class="md:hidden">
				<button onclick={handleMenuClick}><Menu /></button>
			</div>
		</div>
		<div
			class="bg-sea-300 fixed inset-y-0 right-0 w-48 transform flex-col transition-transform duration-300 ease-in-out
			{menuOpen ? '' : 'translate-x-full'}"
		>
			<div class="flex flex-row justify-end px-6 py-3">
				<button onclick={handleMenuClose}>
					<X />
				</button>
			</div>
			<div class="flex flex-col">
				<a
					href="/about"
					class="px-6 py-2 {page.url.pathname === '/about' ? 'font-bold' : 'font-semibold'}"
					>About
				</a>
				<a
					href="/projects"
					class="px-6 py-2 {page.url.pathname === '/projects' ? 'font-bold' : 'font-semibold'}"
					>Projects
				</a>
				{#if data.user}
					<a
						href="/admin"
						class="px-6 py-2 {page.url.pathname === '/admin' ? 'font-bold' : 'font-semibold'}"
						>Admin
					</a>
				{/if}
			</div>
		</div>
	</nav>
</header>
<div class="container mx-auto mt-14 flex flex-col justify-center">
	{@render children()}
</div>
