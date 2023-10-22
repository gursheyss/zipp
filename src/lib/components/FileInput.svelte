<script lang="ts">
	import { Paperclip, Upload, X } from 'lucide-svelte';
	import { fade } from 'svelte/transition';

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

	function checkFileSize(fileList: FileList): FileList | null {
		const MAX_SIZE_MB = 100;
		const MAX_SIZE_BYTES = MAX_SIZE_MB * 1024 * 1024;
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

	$: fileName;
</script>

<div
	class="flex items-center justify-center w-full h-64 border-2 border-darkpink border-dashed rounded-md relative z-0"
	on:dragover={handleDragOver}
	on:drop={handleDrop}
	role="button"
	tabindex="0"
>
	<div class="flex items-center justify-center absolute inset-0">
		<div class="flex items-center justify-center absolute inset-0">
			<Upload color="pink" class="w-5 h-5"/>
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
					<button class="flex items-center pl-1 justify-center" on:click={clearFiles}>
						<X class="w-3 h-3" color="#808080" />
					</button>
				</span>
			</div>
		{/if}
		<input
			type="file"
			bind:files
			on:change={handleFileSelection}
			class="w-full h-full opacity-0 cursor-pointer"
		/>
	</div>
</div>
<button
	class="mt-2 w-full font-medium bg-white hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow {fileName
		? ''
		: 'bg-slate-300'}"
	disabled={!fileName}>share</button
>
