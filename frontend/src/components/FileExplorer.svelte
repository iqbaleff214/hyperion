<script lang="ts">
    import { fileManagement } from "../util/shared.svelte";
    import FileExtensionIcon from "./FileExtensionIcon.svelte";
    import FolderIcon from "./FolderIcon.svelte";

    type Prop = {
        readFile: (path: string) => void,
    };

    let {
        readFile,
    }: Prop = $props();

    let openFolders = $state<Record<string, boolean>>({});

    const toggleFolder = (path: string) => {
        openFolders[path] = !openFolders[path];
    };

    const textClass = (name: string, value: string): string => {
        if (value in fileManagement.obfuscated)
            return 'text-lime-600 hover:text-lime-700';

        if (fileManagement.currentPath === value)
            return 'dark:text-neutral-200 text-neutral-800 hover:dark:text-neutral-200 hover:text-neutral-800';

        return 'text-neutral-500 hover:dark:text-neutral-200 hover:text-neutral-800';
    }
</script>

<div class="py-1.5 px-3 mb-1 flex flex-col items.start border-b border-black/15 dark:border-white/15">
    <h5 class="font-medium text-gray-500 dark:text-gray-400 truncate">File Explorer</h5>
    {#if (fileManagement.folderPath)}
        <small class="text-gray-400 dark:text-gray-500 truncate">{fileManagement.folderPath}</small>
    {/if}
</div>

<div class="pt-1.5 px-3 overflow-y-auto overflow-x-hidden flex-1 text-black dark:text-gray-300">
    {@render node(fileManagement.tree)}
</div>

{#snippet node(folders)}
    <ul>
        {#each Object.entries(folders) as [name, value]}
            <li>
                {#if typeof value === "object"}
                    <button class="cursor-pointer w-full text-start truncate flex items-center gap-x-1 text-sm font-light text-neutral-500 hover:dark:text-neutral-200 hover:text-neutral-800"
                            onclick={() => toggleFolder(name)}>
                        <FolderIcon opened={openFolders[name]}/>
                        <span class="truncate">{name}</span>
                    </button>
                    {#if openFolders[name]}
                        <div class="ps-2">
                            {@render node(value)}
                        </div>
                    {/if}
                {:else if typeof value === "string"}
                    <button class="cursor-pointer w-full text-start truncate flex items-center gap-x-1 text-sm font-light { textClass(name, value) }"
                            onclick={() => readFile(value)}>
                        <FileExtensionIcon path={name}/>
                        <span class="truncate">{name}</span>
                    </button>
                {/if}
            </li>
        {/each}
    </ul>
{/snippet}