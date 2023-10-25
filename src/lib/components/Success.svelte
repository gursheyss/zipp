<script lang="ts">
	export let data: { fileURL: string; password: string };
</script>

<div class="flex flex-col items-center justify-center w-full h-32 relative z-0 p-4">
	<p class="text-2xl text-white mb-4">your file was successfully uploaded!</p>
	<a
		href={'https://zipp-gursheyss.vercel.app' + data.fileURL}
		class="text-pink text-lg hover:text-darkpink transition-colors duration-200"
	>
		{'https://zipp-gursheyss.vercel.app' + data.fileURL}
	</a>
	<br />
	{#if data.password}
		<p class="text-white mt-2">your password is: {data.password}</p>
	{:else}
		<p class="text-white mt-2">your password is: the same one you used to upload it</p>
	{/if}
</div>
