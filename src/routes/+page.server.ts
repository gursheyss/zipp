import { API_URL } from '$env/static/private';
import { randomBytes } from 'crypto';
import ShortUniqueId from 'short-unique-id';

export const actions = {
	upload: async ({ request }) => {
		try {
			const form = await request.formData();
			const file = form.get('file-upload') as File;
			let password = form.get('password') as string;
			let gennedPassword = false;

			// generate random pass if its not provided
			if (!password) {
				password = randomBytes(16).toString('hex');
				gennedPassword = true;
			}

			// generate unique id for file
			const { randomUUID } = new ShortUniqueId({ length: 10 });
			const id = randomUUID();

			const data = new FormData();
			data.append('id', id);
			data.append('file', file);
			data.append('password', password);

			const response = await fetch(API_URL + '/upload', {
				method: 'POST',
				body: data
			});

			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			return {
				fileURL: `/file/${id}`,
				...(gennedPassword ? { password: `${password}` } : {})
			};
		} catch (error) {
			console.error('An error occurred:', error);
			throw error;
		}
	}
};
