<script lang="ts">
    import { fileManagement } from "../util/shared.svelte";
    import PreviewTab from "./PreviewTab.svelte";

    type Prop = {
        closeFile: (path: string) => void,
        children: () => any,
    };

    let {
        closeFile, children
    }: Prop = $props();

    let paths: string[] = $derived(Object.keys(fileManagement.files));

    const closeTab = (index: number) => {
        const path = paths[index];
        closeFile(path);
    };
</script>

<div class="flex flex-col overflow-y-auto grow bg-neutral-50 dark:bg-neutral-900 overflow-x-hidden h-full">
    <PreviewTab {paths} {closeTab}/>

    {#if (fileManagement.currentPath)}
    <div class="grow overflow-y-auto relative w-full">
        {@render children?.()}
        <div class="text-sm font-mono">
            {#each (fileManagement.obfuscated[fileManagement.currentPath] ?? fileManagement.files[fileManagement.currentPath]).split('\n') as row, i}
                <div class="flex gap-x-5 hover:bg-black/5 dark:hover:bg-white/10">
                    <div class="min-w-[30px] text-end text-neutral-400 dark:text-neutral-600">{i+1}</div>
                    <pre class="whitespace-pre-wrap text-neutral-600 dark:text-neutral-400 break-words">{row}</pre>
                </div>
            {/each}
        </div>
    </div>
    {/if}
</div>