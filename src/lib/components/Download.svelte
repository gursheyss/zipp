<script lang="ts">
	export let data: any;
</script>

{#if data.exists}
	<form
		id="downloadform"
		action="?/download"
		method="GET"
		class="flex items-center justify-center w-full relative z-0"
	>
		<div>
			<input
				form="downloadform"
				id="password"
				name="password"
				class="text-center mt-2 w-full font-medium bg-white hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow"
				placeholder="password"
				autocomplete="false"
			/>
			<button
				form="downloadform"
				class="mt-2 bg-white w-full font-medium hover:bg-opacity-95 transition rounded-[6px] depth-white py-2 box-shadow"
				type="submit">download</button
			>
		</div>
	</form>
{:else}
	<div class="flex flex-col justify-center items-center h-screen m-auto gap-2 sm:w-[34rem] px-4">
		<h1 class="text-2xl font-obviouslywide font-bold mb-4 text-pink">file not found :(</h1>
	</div>
{/if}
