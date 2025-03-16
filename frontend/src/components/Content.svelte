<script lang="ts">
    let width = $state<Number>(15);
    let isResizing = $state(false);
    let container: HTMLDivElement | null = null;

    type Prop = {
        isExplorerOpen: boolean,
        sideContent: () => any,
        mainContent: () => any,
    };

    let {
        isExplorerOpen,
        sideContent, mainContent,
    }: Prop = $props();


    function startResize(event: MouseEvent) {
        isResizing = true;
        document.addEventListener("mousemove", resize);
        document.addEventListener("mouseup", stopResize);
    }

    function resize(event: MouseEvent) {
        if (!isResizing || !container) return;

        const containerWidth = container.clientWidth;
        const offsetLeft = container.getBoundingClientRect().left;

        width = Math.min(Math.max(((event.clientX - offsetLeft) / containerWidth) * 100, 10), 90);
    }

    function stopResize() {
        isResizing = false;
        document.removeEventListener("mousemove", resize);
        document.removeEventListener("mouseup", stopResize);
    }
</script>

<div class="flex flex-1" bind:this={container}>
    {#if (isExplorerOpen)}
        <div class="text-sm overflow-x-hidden h-full flex flex-col" style="width: {width}%">
            {@render sideContent?.()}
        </div>

        <div role="presentation" class="relative" onmousedown={startResize}>
            <div class="absolute h-full w-[1px] hover:w-[2px] cursor-ew-resize {isResizing ? 'bg-blue-500' : 'bg-black/15 dark:bg-white/15'} hover:bg-blue-500 transition"></div>
        </div>
    {/if}

    <div class="flex-1 h-full">
        {@render mainContent?.()}
    </div>
</div>