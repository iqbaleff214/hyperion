<script lang="ts">
    import closeIcon from './icons/close.svg';
    import closeDarkIcon from './icons/close-dark.svg';

    import {theme, fileManagement} from '../util/shared.svelte';
    import FileExtensionIcon from "./FileExtensionIcon.svelte";

    type Prop = {
        paths: string[],
        closeTab: (index: number) => void,
    };

    let {
        paths, closeTab = (i) => null,
    }: Prop = $props();


    let hovered: number = $state<number>(0);

    const onHover = (i: number) => {
        hovered = i;
    };

    const onHoverLeave = (i: number) => {
        if (hovered === i) {
            hovered = null;
        }
    };

    const onActive = (i: number) => {
        fileManagement.currentPath = paths[i];
    };
</script>

<div class="shrink-0 h-[30px] flex gap-1 dark:text-white w-full border-b border-black/15 dark:border-white/15">
    <div class="text-sm overflow-x-auto flex gap-x-1 flex-1 w-full" style="scrollbar-width: none;">
        {#each paths as tab, i}
            <div onmouseenter={() => onHover(i)} onmouseleave={() => onHoverLeave(i)} onclick={() => onActive(i)}
                 role="none"
                 class="ps-1.5 pe-6 text-sm font-light hover:dark:text-neutral-200 hover:text-neutral-800 flex items-center justify-start gap-x-1 relative truncate {fileManagement.currentPath === tab ? 'border-b-2 border-blue-500 text-neutral-800 dark:text-neutral-200' : 'text-neutral-500'} ">
                <FileExtensionIcon path={tab}/>
                <span class="truncate cursor-pointer { paths[i] in fileManagement.obfuscated ? 'text-lime-600 font-medium' : '' }">{tab?.split('/').pop()}</span>

                {#if (hovered === i)}
                    <button class="hover:bg-neutral-200 dark:hover:bg-neutral-700 rounded-full p-[1px] absolute right-1.5"
                            onclick={() => closeTab(i)}>
                        {#if (theme.dark)}
                            <img src={closeDarkIcon} alt="close tab" class="close-button">
                        {:else}
                            <img src={closeIcon} alt="close tab" class="close-button">
                        {/if}
                    </button>
                {/if}
            </div>
        {/each}
    </div>
</div>

<style>
    img.close-button {
        width: 12px;
        height: 12px;
    }
</style>