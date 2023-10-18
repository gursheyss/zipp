<script lang="ts">
	import { Paperclip, Upload } from 'lucide-svelte';

	let files: FileList | null = null;
	let fileName = '';

	function handleDragOver(event: DragEvent) {
		event.preventDefault();
	}

	function handleDrop(event: DragEvent) {
		event.preventDefault();
		if (event.dataTransfer) {
			files = event.dataTransfer.files;
			fileName = getFileName(files);
		}
	}

	function getFileName(files: FileList | null): string {
		return files && files.length > 0 ? files[0].name : '';
	}

	$: fileName = getFileName(files);
</script>

<div
	class="flex items-center justify-center w-full h-64 border-2 border-darkpink border-dashed rounded-md relative"
	on:dragover={handleDragOver}
	on:drop={handleDrop}
	role="button"
	tabindex="0"
>
	<div class="flex items-center justify-center absolute inset-0">
		<div class="flex items-center justify-center absolute inset-0">
			<Upload color="pink" />
			<span class="ml-2 text-pink">upload or drag here</span>
		</div>
		{#if fileName}
			<div class="flex items-center justify-center absolute inset-x-0 bottom-10">
				<span class="ml-2 text-pink flex flex-row items-center">
					<Paperclip class="w-4 h-4 mr-1" />
					{fileName}
				</span>
			</div>
		{/if}
		<input type="file" bind:files class="w-full h-full opacity-0 cursor-pointer" />
	</div>
</div>
<button
	class="mt-2 w-full font-medium bg-white hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow {fileName
		? ''
		: 'bg-slate-300'}"
	disabled={!fileName}>share</button
>
