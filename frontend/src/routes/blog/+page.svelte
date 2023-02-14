<script lang="ts">
    import { Modal, modalStore } from '@skeletonlabs/skeleton';
    import type { ModalSettings } from '@skeletonlabs/skeleton';

    
function triggerAlert(): void {
	const alert: ModalSettings = {
		type: 'alert',
		title: 'Example Alert',
		body: 'This is an example modal.',
		// Optionally override buttont text
		buttonTextCancel: 'Cancel'
	};
	modalStore.trigger(alert);
}
				
</script>

<div class="container p-10 space-y-4 h-full" style="background-color: #374C6D">
    <button class="btn variant-ghost-surface" on:click={triggerAlert}>Trigger</button>
</div>
<pre>queue: {JSON.stringify($modalStore, null, 2)}</pre>