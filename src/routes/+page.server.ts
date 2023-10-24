import { API_URL } from "$env/static/private";

export const actions = {
	upload: async ({ request }) => {
		const form = await request.formData();
		const file = form.get('file-upload') as File;
		const password = form.get('password') as string;

		const data = new FormData();
		data.append('file', file);
		data.append('password', password);
		console.log(password);

		const response = await fetch(API_URL + '/upload', {
			method: 'POST',
			body: data
		});

		const responseData = await response.json();
		console.log(responseData);
	}
};
