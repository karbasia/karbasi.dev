export interface PaginationMeta {
	page: number;
	page_size: number;
	total_items: number;
	total_pages: number;
}

export interface PaginatedResponse<T> {
	items: T[];
	pagination: PaginationMeta;
}
