<script lang="ts">
	import Footer from '$lib/components/Footer.svelte';
	import Success from '$lib/components/Success.svelte';
	import Upload from '$lib/components/Upload.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<div class="flex flex-col justify-center items-center h-screen m-auto gap-2 sm:w-[34rem] px-4">
	<h1 class="text-4xl font-obviouslywide font-bold mb-4 text-pink">zipp</h1>
	{#if form?.fileURL}
		<Success data={form} />
	{:else}
		<Upload />
	{/if}
	<Footer />
</div>
