<script lang="ts">
    import logo from "/src/assets/images/hyperion-blue.png";
    import {fileManagement, theme} from "./util/shared.svelte";
    import Sidebar from "./components/Sidebar.svelte";
    import Content from "./components/Content.svelte";
    import Preview from "./components/Preview.svelte";
    import FileExplorer from "./components/FileExplorer.svelte";
    import {buildTree, removeCommonPrefix} from "./util/helpers";
    import {MessageInfoDialog, OpenDirectoryDialog, OpenMultipleFilesDialog} from "../wailsjs/go/filesystem/Dialog";
    import {ReadFile} from "../wailsjs/go/filesystem/FileManager";
    import {onDestroy, onMount} from "svelte";

    let isExplorerOpen: boolean = $state<boolean>(true);

    const toggleExplorer = () => {
        isExplorerOpen = !isExplorerOpen;
    };

    const openFile = () => {
        OpenMultipleFilesDialog().then((paths: string[]) => {
            closeFolder();
            const {commonPrefix, paths: items} = removeCommonPrefix(paths);
            fileManagement.tree = buildTree(items);
            fileManagement.folderPath = commonPrefix;
            if (paths.length > 0) {
                for (const [key, value] of Object.entries(fileManagement.tree)) {
                    if (typeof value === 'string') {
                        fileManagement.currentPath = value;
                        readFile(value);
                        break;
                    }
                }
            }
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });
    };

    const readFile = (path: string) => {
        ReadFile(fileManagement.folderPath + '/' + path).then((content: string) => {
            fileManagement.files[path] = content;
            fileManagement.currentPath = path;
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });
    };

    const closeFile = (path: string) => {
        delete fileManagement.files[path];

        if (!fileManagement.files[fileManagement.currentPath]) {
            const paths: string[] = Object.keys(fileManagement.files);
            fileManagement.currentPath = path.length > 0 ? paths[0] : null;
        }
    };

    const openFolder = () => {
        OpenDirectoryDialog().then((paths: string[]) => {
            closeFolder();
            const {commonPrefix, paths: items} = removeCommonPrefix(paths);
            fileManagement.tree = buildTree(items);
            fileManagement.folderPath = commonPrefix;
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });
    };

    const closeFolder = () => {
        fileManagement.tree = {};
        fileManagement.files = {};
        fileManagement.folderPath = '';
        fileManagement.currentPath = '';
    };

    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const handleThemeChange = (event: MediaQueryListEvent) => {
        theme.dark = event.matches
    };
    mediaQuery.addEventListener('change', handleThemeChange);
    onDestroy(() => {
        mediaQuery.removeEventListener('change', handleThemeChange);
    });

    onMount(() => {
        window?.runtime?.EventsOn("menu:open", openFile);
        window?.runtime?.EventsOn("menu:open-folder", openFolder);
        window?.runtime?.EventsOn("menu:close", () => closeFile(fileManagement.currentPath));
        window?.runtime?.EventsOn("menu:close-folder", closeFolder);
    });
</script>

<div class="flex h-full w-full transition-colors duration-300 overflow-hidden bg-neutral-100 dark:bg-neutral-800">
    <Sidebar
            folderFunction={toggleExplorer}
            folderActive={isExplorerOpen}/>

    <Content {isExplorerOpen}>
        {#snippet sideContent()}
            <FileExplorer {readFile}/>
        {/snippet}

        {#snippet mainContent()}
            {#if (Object.keys(fileManagement.tree).length === 0)}
                <div class="flex flex-grow w-full h-full p-20 overflow-hidden dark:text-white items-center max-w-screen-md mx-auto">
                    <div class="flex flex-col w-full gap-2">
                        <div class="text-4xl font-medium">Hyperion</div>
                        <div class="opacity-50 font-light">Fast and Secure</div>
                        <div class="text-lg font-medium mt-5">Start</div>
                        <button class="inline-block text-start text-sm cursor-pointer text-blue-500" onclick={openFile}>
                            Open...
                        </button>
                        <button class="inline-block text-start text-sm cursor-pointer text-blue-500" onclick={openFolder}>
                            Open Folder...
                        </button>
                    </div>
                    <img class="min-w-1 mx-auto my-auto object-scale-down max-h-full max-h-full grayscale opacity-25 pointer-events-none"
                         src={logo}
                         alt="hyperion"/>
                </div>
                <button onclick={openFile} class="text-black dark:text-white">Open File</button>
                <button onclick={openFolder} class="text-black dark:text-white">Open Folder</button>
            {:else if (Object.keys(fileManagement.files).length > 0)}
                <Preview {closeFile}/>
            {:else}
                <div class="flex items-center justify-center h-full">
                    <img class="min-w-1 mx-auto my-auto object-scale-down max-h-full max-h-full grayscale opacity-25 pointer-events-none"
                         src={logo} alt="hyperion"/>
                </div>
            {/if}
        {/snippet}
    </Content>
</div>