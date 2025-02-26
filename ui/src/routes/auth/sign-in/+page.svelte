<script lang="ts">
	import type { PageProps } from './$types';

	let { form }: PageProps = $props();
</script>

<div class="flex min-h-screen items-center justify-center">
	<div class="dark: bg-surface-700 bg-etch-50 w-full max-w-md rounded-xl p-8 shadow-lg">
		<h2 class="text-midnight-900 mb-6 text-center text-2xl font-bold">Sign In</h2>

		<form method="POST" class="space-y-4">
			{#if form?.incorrect}<div class="text-rose-700">Sign in failed</div>{/if}
			<div>
				<label class="mb-1 block text-sm font-medium" for="email">Email</label>
				<input
					id="email"
					type="email"
					name="email"
					required
					value={form?.email ?? ''}
					class="border-midnight-700 focus:border-secondary-500 focus:ring-secondary-500 w-full rounded-lg border px-4 py-2 outline-none transition-all focus:ring-2"
					placeholder="me@karbasi.dev"
				/>
			</div>

			<div>
				<label class="mb-1 block text-sm font-medium" for="password">Password</label>
				<input
					id="password"
					type="password"
					name="password"
					required
					class="border-midnight-700 focus:border-secondary-500 focus:ring-secondary-500 w-full rounded-lg border px-4 py-2 outline-none transition-all focus:ring-2"
					placeholder="••••••••"
				/>
			</div>

			<button
				class="bg-secondary-700 hover:bg-secondary-800 w-full rounded-lg py-2.5 font-medium text-white transition-colors"
			>
				Sign In
			</button>
		</form>
	</div>
</div>
