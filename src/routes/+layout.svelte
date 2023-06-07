<script lang='ts'>
	// The ordering of these imports is critical to your app working properly
	import '../theme.postcss';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/skeleton.css';
	// Most of your app wide CSS should be put in this file
	import '../app.postcss';
	import 'highlight.js/styles/github-dark.css';

	import { AppBar, AppShell, LightSwitch } from '@skeletonlabs/skeleton';
	import { page } from '$app/stores';
</script>

<AppShell regionPage="relative" slotPageHeader="sticky top-0 z-10">
	<svelte:fragment slot="header">
		<AppBar padding="0">
			<svelte:fragment slot="lead">
				<a href="/" class="btn text-xl lowercase">Karbasi</a>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a class="btn btn-sm {$page.url.pathname === '/about' ? 'variant-filled-primary' : 'bg-initial'}" href="/about">About</a>
				<a class="btn btn-sm {$page.url.pathname === '/projects' ? 'variant-filled-primary' : 'bg-initial'}" href="/projects">Projects</a>
				<LightSwitch />
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<div class="container mx-auto flex justify-center mt-3">
		<slot />
	</div>
	<svelte:fragment slot="footer">
		<div class="mx-auto flex justify-center text-xs">Running on recycled hardware from my closet</div>
	</svelte:fragment>
</AppShell>