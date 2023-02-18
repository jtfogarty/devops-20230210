<script>
	import { IconHexagons } from '@tabler/icons-svelte';
  	export let tree = {};
  	export let level = 0;
	export let onClick = null; // click handler
	/* const {key, value} = tree;
	let expanded = false; */
  function toggle(event, name, node) {
	//expanded = !expanded;
	if( node != 'mod' ){
	    event.target.parentNode.querySelector('ul').classList.toggle("hidden");
	}
	if (node === 'mod'){
		console.log(name)
	}
	onClick(tree, level);
/* 	if (!value && !!onClick){
		
	}
 */  }
</script>
<ul style="padding-left: {level * 10}px">
  {#each Object.entries(tree) as [key, value]}
	{#if Array.isArray(value)}
		{#each value as item}
			{#if item.name}
				<li class="pl-[{level * 5}rem] flex items-center">
				<span class="flex items-center space-x-2"
					on:click={(event) => toggle(event, item.name, 'mod')}
					on:keydown={(event) => { if (event.key === 'Enter') toggle(); }}
					on:keyup={(event) => { if (event.key === 'Escape') toggle(); }}
					on:keypress={(event) => { if (event.key === ' ') { event.preventDefault(); toggle(); } }} >
					<IconHexagons class="w-4 h-4 text-red-500 stroke-1" />
					<span>{item.name}</span> <!-- This is where the modules are printed -->
				</span>
				</li>
			{/if}
		{/each}
	{:else}
		<li style="padding-left: {level * 5}px">
			<span on:click={(event) => toggle(event, '', 'folder')}>{key}</span> <!-- This is where the folder names are printed are printed -->
			{#if typeof value === 'object'}
				<ul class="hidden" style="padding-left: {5}rm">
					<svelte:self tree={value} level={level + 1} {onClick}/>
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
</style>