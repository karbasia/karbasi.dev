<script lang="ts">
	import '../app.css';
	import NavBar from '$lib/components/layout/NavBar.svelte';
	import { ModeWatcher } from 'mode-watcher';
	import Footer from '$lib/components/layout/Footer.svelte';

	let { children } = $props();
</script>

<ModeWatcher defaultMode={'dark'} />

<NavBar></NavBar>

<main class=" mx-auto flex flex-1 px-4 pb-14 pt-16">
	{@render children()}
</main>

<Footer />
