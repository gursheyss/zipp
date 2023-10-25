import { API_URL } from '$env/static/private';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const id = params.id;

	const response = await fetch(API_URL + `/check?id=${id}`, {
		method: 'GET'
	});
	const exists = (await response.text()) === 'true';
	return {
		id: id,
		exists: exists
	};
};
