<script lang="ts">
	import { UserCog } from 'lucide-svelte';

	const currentYear = new Date().getFullYear();

	const socialLinks = [
		{ href: 'https://github.com/karbasia', label: 'GitHub' },
		{ href: 'https://linkedin.com/in/amirkarbasi', label: 'LinkedIn' },
	];
</script>

<footer class="fixed bottom-0 left-0 right-0 backdrop-blur-sm">
	<div class="container mx-auto px-4 py-4">
		<div class="flex flex-row">
			<!-- Copyright and links -->
			<a href="/admin" class="text-muted-foreground hover:text-foreground text-sm"> <UserCog /> </a>
			<div class="flex flex-1 flex-row justify-center">
				<p class="text-muted-foreground text-sm">
					© {currentYear} Amir Karbasi. All opinions are my own.
				</p>
			</div>

			<!-- Social links -->
			<div class="items-center gap-2">
				{#each socialLinks as social}
					<a
						class="px-2"
						href={social.href}
						target="_blank"
						rel="noopener noreferrer"
						aria-label={social.label}
					>
						{social.label}
					</a>
				{/each}
			</div>
		</div>
	</div>
</footer>
