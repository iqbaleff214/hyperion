<script lang="ts">
    import SidebarItem from "./SidebarItem.svelte";
    import folderIcon from "./icons/folder.svg";
    import folderDarkIcon from "./icons/folder-dark.svg";

    type Prop = {
        folderFunction: () => void,
        folderActive?: boolean
    };

    let {
        folderFunction = () => null, folderActive = false,
    }: Prop = $props();
</script>

<div class="flex flex-col gap-y-1 border-r border-black/15 dark:border-white/15 p-1.5">
    <SidebarItem onclick={folderFunction} active={folderActive}>
        {#snippet darkIcon()}
            <img src={folderDarkIcon} alt="folder icon">
        {/snippet}
        {#snippet lightIcon()}
            <img src={folderIcon} alt="folder icon">
        {/snippet}
    </SidebarItem>
</div>

<style>
    img {
        width: 18px !important;
        height: 18px !important;
        min-width: 18px;
        min-height: 18px;
    }
</style>
