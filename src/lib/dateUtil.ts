export const formatDate = (date: Date) => {
  return new Date(date).toLocaleString(undefined, {
		weekday: 'long',
		year: 'numeric',
		month: 'long',
		day: 'numeric'
	});
};

export const getYear = (date?: Date) => {
	return date ? new Date(date).getFullYear() : new Date().getFullYear();
}