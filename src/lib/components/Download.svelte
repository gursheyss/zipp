<script lang="ts">
	export let data: { id: string; exists: boolean };

	let password: string;
	$: password;

	let invalidPassword = false;
	$: invalidPassword;

	import Loading from './Loading.svelte';
	let isLoading = false;
	let success = false;

	const download = async () => {
		success = false;
		isLoading = true;
		const response = await fetch(`/api/download?id=${data.id}&password=${password}`);
		if (!response.ok) {
			invalidPassword = true;
			isLoading = false;
		} else {
			invalidPassword = false;
			success = true;
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
			isLoading = false;
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
				type="password"
				bind:value={password}
			/>
			<button
				form="downloadform"
				class="mt-2 bg-white w-full font-medium hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow relative"
				type="submit"
				on:click={download}
			>
				<div class="flex justify-center items-center space-x-2">
					<span>download</span>
					{#if isLoading}
						<Loading />
					{/if}
				</div>
			</button>
		</div>
	</div>
	<div class="h-5 text-pink">
		{#if invalidPassword}
			<p>invalid password</p>
		{/if}
		{#if success}
			<p>your files have been downloaded & deleted</p>
		{/if}
	</div>
{:else}
	<div class="flex flex-col justify-center items-center h-screen m-auto gap-2 sm:w-[34rem] px-4">
		<h1 class="text-2xl font-obviouslywide font-bold mb-4 text-pink">file not found :(</h1>
	</div>
{/if}
