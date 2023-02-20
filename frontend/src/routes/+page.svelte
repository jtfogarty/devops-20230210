<script lang='ts'>
	import { gitlabData } from '$lib/store'; 
	import { onMount } from 'svelte';
  
	let dataString = '';
  
	onMount(() => {
	  // Subscribe to changes in the store
	  gitlabData.subscribe((value) => {
		// Convert the updated store data to a string
		dataString = JSON.stringify(value);
	  });
	});
  </script>
  
	
	<!-- Your content here -->
	<div class="container p-10 space-y-4 h-full" style="background-color: #374C6D">
		<h1>Homepage</h1>
		<hr />
		<button >
			Show Store 
		</button>
		{ dataString }
	</div>