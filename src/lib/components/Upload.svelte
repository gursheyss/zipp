<script lang="ts">
	import { enhance } from '$app/forms';
	import { Paperclip, Upload, X } from 'lucide-svelte';
	import { fade } from 'svelte/transition';
	import Loading from './Loading.svelte';

	let isLoading = false;
	let files: FileList | null = null;
	let fileName = '';

	function handleDragOver(event: DragEvent) {
		event.preventDefault();
	}

	function handleDrop(event: DragEvent) {
		event.preventDefault();
		if (event.dataTransfer) {
			files = checkFileSize(event.dataTransfer.files);
			fileName = getFileName(files);
		}
	}

	function handleFileSelection(event: Event) {
		const target = event.target as HTMLInputElement;
		if (target.files) {
			files = checkFileSize(target.files);
			fileName = getFileName(files);
		}
	}

	const MAX_SIZE_MB = 100;
	const MAX_SIZE_BYTES = MAX_SIZE_MB * 1024 * 1024;

	function checkFileSize(fileList: FileList): FileList | null {
		if (fileList[0].size > MAX_SIZE_BYTES) {
			alert('File size exceeds 100MB limit');
			return null;
		}
		return fileList;
	}
	function getFileName(files: FileList | null): string {
		return files && files.length > 0 ? files[0].name : '';
	}

	function clearFiles() {
		files = null;
		fileName = '';
	}

	function handleSubmit() {
		clearFiles();
		isLoading = true;
	}

	$: fileName;
</script>

<form
	id="uploadform"
	action="?/upload"
	method="POST"
	class="flex items-center justify-center w-full h-64 border-2 border-darkpink border-dashed rounded-md relative z-0 hover:border-pink"
	on:dragover={handleDragOver}
	use:enhance={() => {
		return async ({ update }) => {
			files = null;
			update({ reset: false });
		};
	}}
	on:drop={handleDrop}
	on:submit={handleSubmit}
>
	<div class="flex items-center justify-center absolute inset-0">
		<div class="flex items-center justify-center absolute inset-0">
			<Upload color="pink" class="w-5 h-5" />
			<span class="ml-2 text-pink">upload or drag here</span>
		</div>
		{#if fileName}
			<div
				class="flex items-center justify-center absolute inset-x-0 bottom-10 z-10"
				transition:fade={{ delay: 0, duration: 200 }}
			>
				<span class="ml-2 text-pink flex flex-row items-center">
					<Paperclip class="w-4 h-4 mr-1" />
					{fileName}
					<button type="button" class="flex items-center pl-1 justify-center" on:click={clearFiles}>
						<X class="w-3 h-3" color="#808080" />
					</button>
				</span>
			</div>
		{/if}
		<input
			id="file-upload"
			name="file-upload"
			type="file"
			on:change={handleFileSelection}
			class="w-full h-full opacity-0 cursor-pointer"
		/>
	</div>
</form>
<button
	form="uploadform"
	class={`mt-2 w-full font-medium hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow ${
		!fileName ? 'bg-slate-200' : 'bg-white'
	}`}
	type="submit"
	disabled={!fileName}
>
	<span class="flex justify-center items-center">
		share
		{#if isLoading}
			<div class="pl-2">
				<Loading />
			</div>
		{/if}
	</span>
</button>
<input
	form="uploadform"
	id="password"
	name="password"
	type="password"
	class="text-center mt-2 w-full font-medium bg-white hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow"
	placeholder="password (optional)"
	disabled={!fileName}
/>
