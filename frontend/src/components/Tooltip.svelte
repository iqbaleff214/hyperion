<script lang="ts">
    type Prop = {
        text?: string,
        children: () => any,
    }

    let {
        children,
        text,
    }: Prop = $props();
</script>

<div class="tooltip">
    {@render children?.()}
    {#if (text)}
        <span class="text text-sm">{text}</span>
    {/if}
</div>

<style>
    .tooltip {
        position: relative;
    }

    .tooltip .text {
        visibility: hidden;
        width: 120px;
        background-color: #555;
        color: #fff;
        text-align: center;
        border-radius: 6px;
        padding: 5px 0;
        position: absolute;
        z-index: 1;
        top: 0px;
        left: 100px;
        margin-left: -60px;
        opacity: 0;
        transition: opacity 0.3s;
    }

    .tooltip .text::after {
        content: "";
        position: absolute;
        top: 50%;
        right: 100%;
        margin-top: -5px;
        border-width: 5px;
        border-style: solid;
        border-color: transparent #555 transparent transparent;
    }

    .tooltip:hover .text {
        visibility: visible;
        opacity: 1;
    }
</style>