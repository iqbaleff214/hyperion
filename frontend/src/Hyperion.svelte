<script lang="ts">
    import logo from "/src/assets/images/hyperion-blue.png";
    import {fileManagement, theme} from "./util/shared.svelte";
    import Sidebar from "./components/Sidebar.svelte";
    import Content from "./components/Content.svelte";
    import Preview from "./components/Preview.svelte";
    import FileExplorer from "./components/FileExplorer.svelte";
    import {buildTree, flattenTree, removeCommonPrefix} from "./util/helpers";
    import {ReadFile, WriteFile} from "../wailsjs/go/filesystem/FileManager";
    import {Obfuscate} from "../wailsjs/go/obfuscator/Obfuscator";
    import {MessageInfoDialog, OpenDirectoryDialog, OpenMultipleFilesDialog} from "../wailsjs/go/filesystem/Dialog";
    import {onDestroy, onMount} from "svelte";
    import {EventsEmit} from "../wailsjs/runtime";
    import PreviewAction from "./components/PreviewAction.svelte";

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
            EventsEmit("files", true);
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

    const saveFile = (path: string) => {
        if (!(path in fileManagement.obfuscated))
            return;

        const obfuscated: string = fileManagement.obfuscated[path];
        WriteFile(fileManagement.folderPath + '/' + path, obfuscated).then(() => {
            fileManagement.files[path] = obfuscated;

            delete fileManagement.obfuscated[path];
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });

    };

    const saveAll = () => {
        const paths = Object.keys(fileManagement.obfuscated);
        Promise.all(paths.map(path =>
            WriteFile(fileManagement.folderPath + '/' + path, fileManagement.obfuscated[path])
                .then(() => {
                    fileManagement.files[path] = fileManagement.obfuscated[path];
                    delete fileManagement.obfuscated[path];
                })
                .catch((error: any) => {
                    MessageInfoDialog('Error', error);
                })
        )).then(() => {
            console.log('All files saved successfully');
        }).catch(() => {
            console.log('Some files failed to save');
        });
    };

    const closeFile = (path: string) => {
        if (path === '')
            return;

        delete fileManagement.files[path];

        if (!fileManagement.files[fileManagement.currentPath]) {
            const paths: string[] = Object.keys(fileManagement.files);
            fileManagement.currentPath = paths.length > 0 ? paths[0] : null;
        }
    };

    const openFolder = () => {
        OpenDirectoryDialog().then((paths: string[]) => {
            closeFolder();
            const {commonPrefix, paths: items} = removeCommonPrefix(paths);
            fileManagement.tree = buildTree(items);
            fileManagement.folderPath = commonPrefix;
            EventsEmit("files", true);
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });
    };

    const closeFolder = () => {
        fileManagement.tree = {};
        fileManagement.files = {};
        fileManagement.obfuscated = {};
        fileManagement.folderPath = '';
        fileManagement.currentPath = '';
    };

    const obfuscate = (path: string) => {
        Obfuscate(fileManagement.folderPath + '/' + path).then((content: string) => {
            fileManagement.obfuscated[path] = content;
        }).catch((error: any) => {
            MessageInfoDialog('Error', error);
        });
    };

    const obfuscateAll = () => {
        const paths: string[] = flattenTree(fileManagement.tree);
        Promise.all(
            paths.map(path =>
                Obfuscate(fileManagement.folderPath + '/' + path)
                    .then((content: string) => fileManagement.obfuscated[path] = content)
                    .catch((error: any) => MessageInfoDialog('Error', error))
            )
        );
    };

    const undo = (path: string) => {
        delete fileManagement.obfuscated[path];
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

        window?.runtime?.EventsOn("menu:save", () => saveFile(fileManagement.currentPath));
        window?.runtime?.EventsOn("menu:save-as", () => null);
        window?.runtime?.EventsOn("menu:save-all", saveAll);

        window?.runtime?.EventsOn("run:undo", () => undo(fileManagement.currentPath));
        window?.runtime?.EventsOn("run:undo-all", () => fileManagement.obfuscated = {});

        window?.runtime?.EventsOn("run:obfuscate", () => obfuscate(fileManagement.currentPath));
        window?.runtime?.EventsOn("run:obfuscate-all", obfuscateAll);
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
                <Preview {closeFile}>
                    <PreviewAction
                        obfuscate={() => obfuscate(fileManagement.currentPath)}
                        cancel={() => undo(fileManagement.currentPath)}
                    />
                </Preview>
            {:else}
                <div class="flex items-center justify-center h-full">
                    <img class="min-w-1 mx-auto my-auto object-scale-down max-h-full max-h-full grayscale opacity-25 pointer-events-none"
                         src={logo} alt="hyperion"/>
                </div>
            {/if}
        {/snippet}
    </Content>
</div>