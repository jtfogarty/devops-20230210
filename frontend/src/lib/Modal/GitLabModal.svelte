<script lang="ts">
	// Props
	/** Exposes parent props to this component. */
	export let parent: any;

	// Stores
	//import { writable } from 'svelte/store';
	import { modalStore } from '@skeletonlabs/skeleton';
	import { gitlabData, updateGitlabData } from '$lib/store';


	/* const gitlabData = writable({
    name: '',
    server: '',
    token: ''
	}); */

	// We've created a custom submit function to pass the response and close the modal.
	function onFormSubmit(): void {
		if ($modalStore[0].response) $modalStore[0].response(gitlabData);
		modalStore.close();
	}

	// Base Classes
	const cBase = 'space-y-4';
	const cForm = 'border border-surface-500 p-4 space-y-4 rounded-container-token';
</script>

<!-- @component This example creates a simple form modal. -->

<div class="modal-example-form {cBase}">
	<!-- Enable for debugging: -->
	<!-- <pre>{JSON.stringify(gitlabData, null, 2)}</pre> -->
	<form class="modal-form {cForm}">
		<label class="label">
			<span>User Name</span>
			<input class="input text-center" type="text" bind:value={$gitlabData.name} placeholder="Enter your user name..." />
		</label>
		<label class="label">
			<span>GitLab Server</span>
			<input class="input text-center" type="text" bind:value={$gitlabData.server} placeholder="gitlab.com" />
		</label>
		<label class="label">
			<span>GitLab Personal Access Token</span>
			<input class="input text-center" type="text" bind:value={$gitlabData.token} placeholder="Enter GitLab Token..." />
		</label>
	</form>
	<!-- prettier-ignore -->
	<footer class="modal-footer {parent.regionFooter}">
        <button class="btn {parent.buttonNeutral}" on:click={parent.onClose}>{parent.buttonTextCancel}</button>
        <button class="btn {parent.buttonPositive}" on:click={() => {
			updateGitlabData($gitlabData);
			onFormSubmit();
		  }}>Submit</button>
    </footer>
</div>