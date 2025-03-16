<script lang="ts">
    import {theme} from '../util/shared.svelte';
    import jsIcon from './icons/js.svg';
    import jsDarkIcon from './icons/js-dark.svg';
    import phpIcon from './icons/php.svg';
    import phpDarkIcon from './icons/php-dark.svg';
    import pythonIcon from './icons/python.svg';
    import pythonDarkIcon from './icons/python-dark.svg';
    import rubyIcon from './icons/ruby.svg';
    import rubyDarkIcon from './icons/ruby-dark.svg';

    type ThemeType = 'dark' | 'light';

    type ExtensionIcons = {
        [key: string]: {
            [key in ThemeType]: string;
        };
    };

    let { path }: { path: string } = $props();
    const extension: string = path.split('.').pop().toLowerCase();

    const icons: ExtensionIcons = {
        'js': {
            dark: jsDarkIcon,
            light: jsIcon,
        },
        'php': {
            dark: phpDarkIcon,
            light: phpIcon,
        },
        'rb': {
            dark: rubyDarkIcon,
            light: rubyIcon,
        },
        'py': {
            dark: pythonDarkIcon,
            light: pythonIcon,
        },
    };
</script>

{#if (icons[extension])}
    <img src={icons[extension][theme.dark ? 'dark' : 'light']} alt={`${extension} icon`} class="icon">
{/if}

<style>
    img {
        width: 12px;
        height: 12px;
    }
</style>