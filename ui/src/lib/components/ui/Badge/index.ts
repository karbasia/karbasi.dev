import { type VariantProps, tv } from 'tailwind-variants';
export { default as Badge } from './Badge.svelte';

export const badgeVariants = tv({
	base: 'cursor-pointer focus:ring-ring inline-flex select-none items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2',
	variants: {
		variant: {
			default:
				'dark:bg-secondary-700 bg-secondary-100 text-primary-foreground hover:bg-surface-50 dark:hover:bg-secondary-500 border-transparent',
			secondary:
				'dark:bg-tertiary-700 bg-tertiary-100 text-secondary-foreground hover:bg-tertiary-50 dark:hover:bg-tertiary-500 border-transparent',
			destructive:
				'bg-destructive text-destructive-foreground hover:bg-destructive/80 border-transparent',
			outline: 'text-foreground',
		},
	},
	defaultVariants: {
		variant: 'default',
	},
});

export type Variant = VariantProps<typeof badgeVariants>['variant'];
