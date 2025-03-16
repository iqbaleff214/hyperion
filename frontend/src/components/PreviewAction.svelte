<script lang="ts">
    import { theme, fileManagement } from '../util/shared.svelte';

    import copyIcon from './icons/copy.svg';
    import copyDarkIcon from './icons/copy-dark.svg';
    import copiedIcon from './icons/copied.svg';
    import copiedDarkIcon from './icons/copied-dark.svg';
    import obfuscateIcon from './icons/obfuscate.svg';
    import obfuscateDarkIcon from './icons/obfuscate-dark.svg';
    import obfuscateCancelIcon from './icons/obfuscate-cancel.svg';
    import obfuscateCancelDarkIcon from './icons/obfuscate-cancel-dark.svg';

    type Prop = {
        obfuscate: () => void,
        cancel: () => void,
    };

    let {
        obfuscate, cancel,
    }: Prop = $props();

    let isCopied: boolean = $state<boolean>(false);
    let actions = $state({
        'obfuscate': {
            icon: {
                light: obfuscateIcon,
                dark: obfuscateDarkIcon,
            },
            enable: () => !(fileManagement.currentPath in fileManagement.obfuscated),
            callback: obfuscate,
        },
        'cancel': {
            icon: {
                light: obfuscateCancelIcon,
                dark: obfuscateCancelDarkIcon,
            },
            enable: () => fileManagement.currentPath in fileManagement.obfuscated,
            callback: cancel,
        },
        'copy': {
            icon: {
                light: copyIcon,
                dark: copyDarkIcon,
            },
            enable: () => fileManagement.currentPath in fileManagement.obfuscated,
            callback: () => null,
        },
    });

    const copy = () => {
        navigator.clipboard.writeText(fileManagement.obfuscated[fileManagement.currentPath]).then(() => {
            isCopied = true;
            actions['copy'].icon.light = copiedIcon;
            actions['copy'].icon.dark = copiedDarkIcon;
            setTimeout(() => {
                isCopied = false;
                actions['copy'].icon.light = copyIcon;
                actions['copy'].icon.dark = copyDarkIcon;
            }, 2000);
        }).catch(err => console.error("Failed to copy:", err));
    };

    actions['copy'].callback = copy;
</script>

<div class="fixed top-9.5 right-2 opacity-50 hover:opacity-100 bg-black/5 hover:bg-slate-100 dark:bg-white/15 hover:shadow dark:hover:bg-white/20 rounded-lg p-1.5 flex justify-around gap-x-1">
    {#each Object.keys(actions) as action}
        {@const act = actions[action]}
        {#if (act.enable())}
            <button onclick={act.callback} class="enabled:cursor-pointer p-1 hover:enabled:bg-slate-300 rounded-lg"
                    aria-label={action.replaceAll('_', ' ')}>
                <img src={act.icon[theme.dark ? 'dark' : 'light']} alt={action.replaceAll('_', ' ')} />
            </button>
        {/if}
    {/each}
</div>

<style>
    img {
        width: 14px;
        height: 14px;
    }
</style>