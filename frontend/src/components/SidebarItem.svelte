<script lang="ts">
    import {theme} from '../util/shared.svelte';
    import Tooltip from "./Tooltip.svelte";

    type Prop = {
        onclick?: (event: MouseEvent) => void,
        tooltip?: string,
        active?: boolean,
        lightIcon: () => any,
        darkIcon: () => any,
    }

    let hovered = $state(false);

    let {
        onclick = (_event: MouseEvent) => null,
        lightIcon, darkIcon, active = false,
        tooltip = '',
    }: Prop = $props();
</script>

<Tooltip text={tooltip}>
    <button {onclick}
            onmouseenter={() => hovered = true}
            onmouseleave={() => hovered = false}
            class="{active ? 'bg-black/11 dark:bg-white/20' : ''} cursor-pointer rounded p-1 hover:bg-blue-500"
            aria-label="menu">
        {#if (theme.dark || hovered)}
            {@render darkIcon?.()}
        {:else}
            {@render lightIcon?.()}
        {/if}
    </button>
</Tooltip>
