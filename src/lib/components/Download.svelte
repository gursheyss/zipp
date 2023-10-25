<script lang="ts">
	export let data: { id: string; exists: boolean };

	let password: string;
	$: password;

	const download = async () => {
		const response = await fetch(`/api/download?id=${data.id}&password=${password}`);
		if (!response.ok) {
			const error = await response.text(); // or response.json() if the error is a JSON object
			console.log(error);
		} else {
			const blob = await response.blob();
			const contentDisposition = response.headers.get('Content-Disposition');
			const filenameMatch = contentDisposition
				? contentDisposition.match(/filename="(.*?)"/)
				: null;
			const filename = filenameMatch ? filenameMatch[1] : 'default_filename';
			const newBlob = new Blob([blob], { type: blob.type });
			const url = URL.createObjectURL(newBlob);
			const link = document.createElement('a');
			link.href = url;
			link.setAttribute('download', filename);
			document.body.appendChild(link);
			link.click();
			document.body.removeChild(link);
		}
	};
</script>

{#if data.exists}
	<h1 class="text-4xl font-obviouslywide font-bold mb-4 text-pink">zipp</h1>

	<div class="flex items-center justify-center w-full relative z-0">
		<div>
			<input
				class="text-center mt-2 w-full font-medium bg-white hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow"
				placeholder="password"
				autocomplete="false"
				bind:value={password}
			/>
			<button
				form="downloadform"
				class="mt-2 bg-white w-full font-medium hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow"
				type="submit"
				on:click={download}>download</button
			>
		</div>
	</div>
{:else}
	<div class="flex flex-col justify-center items-center h-screen m-auto gap-2 sm:w-[34rem] px-4">
		<h1 class="text-2xl font-obviouslywide font-bold mb-4 text-pink">file not found :(</h1>
	</div>
{/if}
