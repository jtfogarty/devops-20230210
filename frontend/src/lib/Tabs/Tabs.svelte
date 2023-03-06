<script lang="ts">
    import { TabGroup, Tab } from '@skeletonlabs/skeleton';
    let tabSet: number = 0;

    export let value = {};
    let i = 0;
</script>

<TabGroup>
    <Tab bind:group={tabSet} name="tab0" value={0}>Code</Tab>
    <Tab bind:group={tabSet} name="tab1" value={1}>Document</Tab>
    <svelte:fragment slot="panel">
<!--     {#if tabSet === 0}
        {(value as {git_url: string}[])[0].git_url}
    {:else if tabSet === 1}
        {(value as {document_url: string}[])[0].document_url}
    {/if} -->
    </svelte:fragment>
</TabGroup>
