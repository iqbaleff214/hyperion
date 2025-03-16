<script lang="ts">
    import {theme} from '../util/shared.svelte';
    import folderIcon from './icons/folder.svg';
    import folderDarkIcon from './icons/folder-dark.svg';
    import folderOpenIcon from './icons/folder-open.svg';
    import folderOpenDarkIcon from './icons/folder-open-dark.svg';

    let {opened}: { opened: boolean } = $props();
</script>

{#if (theme.dark)}
    {#if (opened)}
        <img src={folderOpenDarkIcon} alt='opened folder icon' class="icon">
    {:else}
        <img src={folderDarkIcon} alt='closed folder icon' class="icon">
    {/if}
{:else}
    {#if (opened)}
        <img src={folderOpenIcon} alt='opened folder icon' class="icon">
    {:else}
        <img src={folderIcon} alt='closed folder icon' class="icon">
    {/if}
{/if}


<style>
    img {
        width: 12px;
        height: 12px;
    }
</style>