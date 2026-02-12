<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import ChevronLeft from 'lucide-svelte/icons/chevron-left';
	import ChevronRight from 'lucide-svelte/icons/chevron-right';

	interface Props {
		page: number;
		totalPages: number;
		baseUrl: string;
	}

	let { page, totalPages, baseUrl }: Props = $props();

	const prevHref = $derived(page === 2 ? baseUrl || '/' : `${baseUrl}?page=${page - 1}`);
	const nextHref = $derived(`${baseUrl}?page=${page + 1}`);
</script>

{#if totalPages > 1}
	<nav class="mt-8 flex items-center justify-center gap-4">
		{#if page > 1}
			<Button variant="outline" size="sm" href={prevHref}>
				<ChevronLeft class="h-4 w-4" />
				Previous
			</Button>
		{/if}
		<span class="text-muted-foreground text-sm">
			Page {page} of {totalPages}
		</span>
		{#if page < totalPages}
			<Button variant="outline" size="sm" href={nextHref}>
				Next
				<ChevronRight class="h-4 w-4" />
			</Button>
		{/if}
	</nav>
{/if}
