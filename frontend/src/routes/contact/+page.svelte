<script lang='ts'>
    import { page } from '$app/stores'
    const text = $page.url.searchParams.get('text');
</script>

<p>{ text }</p>