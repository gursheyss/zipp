<script lang="ts">
	import Download from '$lib/components/Download.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="flex flex-col justify-center items-center h-screen m-auto gap-2 sm:w-[34rem] px-4">
	<h1 class="text-4xl font-obviouslywide font-bold mb-4 text-pink">zipp</h1>
	<Download {data} />
</div>
