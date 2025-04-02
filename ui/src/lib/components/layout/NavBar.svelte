<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { BookOpenIcon, TagIcon } from 'lucide-svelte';
	import LightSwitch from './LightSwitch.svelte';

	const menuItems = [
		{ label: 'About', icon: BookOpenIcon, href: '/about' },
		{ label: 'Projects', icon: TagIcon, href: '/projects' },
	];
</script>

<nav class="bg-background/80 fixed left-0 right-0 top-0 z-50 backdrop-blur-sm">
	<div class="mx-auto flex h-16 items-center justify-between px-4">
		<a href="/" class="text-xl font-bold">karbasi</a>

		<ul class="flex items-center gap-2">
			{#each menuItems as item}
				<li>
					<Button variant="ghost" size="sm">
						<a href={item.href} class="flex items-center gap-2">
							<svelte:component this={item.icon} class="h-4 w-4" />
							<span class="hidden sm:inline">{item.label}</span>
						</a>
					</Button>
				</li>
			{/each}
			<li><LightSwitch /></li>
		</ul>
	</div>
</nav>
