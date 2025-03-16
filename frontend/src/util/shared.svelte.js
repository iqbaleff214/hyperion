export const theme = $state({
    dark: window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
});

export const fileManagement = $state({
    tree: {},
    files: {},
    folderPath: '',
    currentPath: '',
});