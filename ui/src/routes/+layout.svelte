<script lang="ts">
	import '../app.css';
	import NavBar from '$lib/components/layout/NavBar.svelte';
	import { ModeWatcher } from 'mode-watcher';

	let { children } = $props();
</script>

<ModeWatcher defaultMode={'dark'} />

<NavBar></NavBar>

<main class="container mx-auto flex flex-1 gap-8 px-4 pt-20">
	{@render children()}
</main>
