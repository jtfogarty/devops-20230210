<script lang='ts'>
	import { createEventDispatcher } from 'svelte';
	import { IconHexagons } from '@tabler/icons-svelte';
  	export let tree = {};
  	export let level = 0;

	const dispatch = createEventDispatcher();

	function toggle(event: (MouseEvent & { currentTarget: EventTarget & HTMLSpanElement; }) |
	 undefined, node: string | undefined, value: unknown) {
	
	if( node != 'mod' ){
	    (event!.target as Element).parentNode!.querySelector('ul')!.classList.toggle("hidden");
	}
	
	if (node === 'mod') {
		let data = (value as {document_url: string}[])[0].document_url;
		console.log("child clicked" , data);
		dispatch("NavClick", data);
		//console.log(value)
		//console.log((value as {name: string}[])[0].name);
		//console.log((value as {git_url: string}[])[0].git_url);
		//console.log((value as {document_url: string}[])[0].document_url);
    }
  }
</script>
<ul style="padding-left: {level * 10}px">
  {#each Object.entries(tree) as [key, value]}
	{#if Array.isArray(value)}
		{#each value as item}
			{#if item.name}
				<li class="pl-[{level * 5}rem] flex items-center">
				<span class="flex items-center space-x-2"
					on:click={(event) => toggle(event, 'mod', value)}
					on:keydown={(event) => { if (event.key === 'Enter') toggle(undefined, 'mod', value); }}
					on:keyup={(event) => { if (event.key === 'Escape') toggle(undefined, 'mod', value); }}
					on:keypress={(event) => { if (event.key === ' ') { event.preventDefault(); toggle(undefined, 'mod', value); } }} >
					<IconHexagons class="w-4 h-4 text-red-500 stroke-1" />
					<span>{item.name}</span> <!-- This is where the modules are printed -->
				</span>
				</li>
			{/if}
		{/each}
	{:else}
		<li style="padding-left: {level * 5}px">
			<span on:click={(event) => toggle(event, key, 'folder')}
				on:keydown={(event) => { if (event.key === 'Enter') toggle(undefined, 'mod', value); }} 
				on:keyup={(event) => { if (event.key === 'Escape') toggle(undefined, 'mod', value); }}
				on:keypress={(event) => { if (event.key === ' ') toggle(undefined, 'mod', value); }}>{key}</span> <!-- This is where the folder names are printed are printed -->
			{#if typeof value === 'object'}
				<ul class="hidden" style="padding-left: {5}rm">
					<svelte:self tree={value} level={level + 1} />
				</ul>
			{:else}
				{value}			<!-- This is never used -->
			{/if}
		</li>
	{/if}
  {/each}
</ul>
<style>
  .hidden {
    display: none;
  }
  span {
    cursor: pointer;
  }
/*   li:not(.hidden):hover {
	background: #007ba7
} */

</style>