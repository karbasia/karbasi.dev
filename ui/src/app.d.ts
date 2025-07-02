// See https://svelte.dev/docs/kit/types#app.d.ts

import type { UserCore } from '$lib/models/user';

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: UserCore | null;
			token?: string;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
