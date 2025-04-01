<script lang="ts">
	import { Badge } from "$lib/components/ui/badge";
	import { Card, CardHeader, CardTitle, CardContent } from "$lib/components/ui/card";
	import { TagIcon } from "lucide-svelte";

	type Props = {
		tagCounts: Record<string, number>;
	};

	let { tagCounts }: Props = $props();
</script>

<Card class="sticky top-20">
	<CardHeader>
		<CardTitle class="flex items-center gap-2">
			<TagIcon class="h-5 w-5" />
			Tags
		</CardTitle>
	</CardHeader>
	<CardContent>
		<div class="flex flex-col gap-2">
			{#each Object.entries(tagCounts) as [tag, count]}
				<div class="flex items-center justify-between">
					<Badge variant="outline" class="cursor-pointer hover:bg-secondary">
						{tag}
					</Badge>
					<span class="text-sm text-muted-foreground">{count} posts</span>
				</div>
			{/each}
		</div>
	</CardContent>
</Card>
