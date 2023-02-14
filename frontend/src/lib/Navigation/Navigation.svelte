<script>
	import {IconHexagons} from '@tabler/icons-svelte';
  export let tree = {};
  export let level = 0;
	export let onClick = null; // click handler
	
  function toggle(event) {
    event.target.parentNode.querySelector('ul').classList.toggle("hidden");
		onClick(tree, level);
  }
</script>
<ul style="padding-left: {level * 10}px">
  {#each Object.entries(tree) as [key, value]}
    
			{#if Array.isArray(value)}
				{#each value as item}
					{#if item.name}
					<li class="pl-[{level * 5}rem] flex items-center">
						<span class="flex items-center space-x-2" on:click={toggle} {onClick}>
							<IconHexagons class="w-4 h-4 text-red-500 stroke-1" />
							<span>{item.name}</span>
						</span>
					</li>
 <!-- This is where the modules are printed -->
					{/if}
				{/each}
			{:else}
				<li style="padding-left: {level * 5}px">
				<span on:click={toggle} {onClick}>{key} </span>
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