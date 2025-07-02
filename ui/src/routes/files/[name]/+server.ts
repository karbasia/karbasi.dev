import { httpRequestEnum, type RequestParams } from '$lib/models/api';
import { createRawRequest } from '$lib/server/api';
import { error } from '@sveltejs/kit';

export async function GET({ params }) {
	const reqParams: RequestParams = {
		method: httpRequestEnum.enum.GET,
		path: `/files/${params.name}`,
	};
	const fileRes = await createRawRequest(reqParams);
	if ('error' in fileRes) return error(fileRes.code, fileRes.error);
	return fileRes;
}
