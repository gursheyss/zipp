import { API_URL } from '$env/static/private';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const id = params.id;
	const response = await fetch(API_URL + `/check?id=${id}`, {
		method: 'GET'
	});
	const exists = (await response.text()) === 'true';
	return {
		exists: exists
	};
};

export const actions = {
	download: async ({ params, request }) => {
		const form = await request.formData();
		const password = form.get('password') as string;
		const id = params.id;
		const response = await fetch(API_URL + `/download?id=${id}&password=${password}`, {
			method: 'GET'
		});

		const decryptedFiles = await response.json();
		return {
			decryptedFiles,
			id: id
		};
	}
};
