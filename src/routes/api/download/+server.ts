import { API_URL } from '$env/static/private';
import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ request }) => {
    const requestUrl = new URL(request.url);
    const id = requestUrl.searchParams.get('id');
    const password = requestUrl.searchParams.get('password');

    const response = await fetch(`${API_URL}/download?id=${id}&password=${password}`, {
        method: 'GET'
    });

    if (!response.ok) {
        return new Response('Error fetching file', { status: response.status });
    } else {
        const blob = await response.blob();
        const contentDisposition = response.headers.get('Content-Disposition');
        if (contentDisposition === null) {
            return new Response('Error fetching file headers', { status: 500 });
        }
        console.log(contentDisposition)
        const headers = {
            'Content-Disposition': contentDisposition,
            'Content-Type': blob.type
        };
        return new Response(blob, { status: 200, headers });
    }
};
