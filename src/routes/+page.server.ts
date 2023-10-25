import { API_URL } from '$env/static/private';
import { randomBytes } from 'crypto';
import ShortUniqueId from 'short-unique-id';

export const actions = {
	upload: async ({ request }) => {
		try {
			const form = await request.formData();
			const file = form.get('file-upload') as File;
			let password = form.get('password') as string;

			// generate random pass if its not provided
			if (!password) {
				password = randomBytes(32).toString('hex');
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

			const responseData = await response.json();
			console.log(responseData);

			if (!form.get('password')) {
				responseData.password = password;
			}

			responseData.fileURL = `file/${id}`;

			return responseData;
		} catch (error) {
			console.error('An error occurred:', error);
			throw error;
		}
	}
};
