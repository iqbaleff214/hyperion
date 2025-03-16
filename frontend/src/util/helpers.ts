export function buildTree(paths: string[]): Record<string, any> {
    const tree: Record<string, any> = {};

    for (const path of paths) {
        const parts = path.split("/");
        let current = tree;

        for (let i = 0; i < parts.length; i++) {
            const part = parts[i];

            if (!current[part]) {
                current[part] = i === parts.length - 1 ? path : {};
            }

            current = current[part];
        }
    }

    return tree;
}

export function removeCommonPrefix(paths: string[]): { commonPrefix: string, paths: string[] } {
    if (paths.length === 0)
        return { commonPrefix: "", paths: [] };

    if (paths.length === 1) {
        const path = paths[0];

        const lastSlashIndex = path.lastIndexOf("/");
        if (lastSlashIndex === -1)
            return { commonPrefix: "", paths: [path] };

        const commonPrefix = path.substring(0, lastSlashIndex);
        const filename = path.substring(lastSlashIndex + 1);

        return { commonPrefix, paths: [filename] };
    }

    const splitPaths = paths.map(path => path.split("/"));

    let prefixLength = splitPaths[0].length;
    for (let i = 0; i < prefixLength; i++) {
        const segment = splitPaths[0][i];
        if (!splitPaths.every(path => path[i] === segment)) {
            prefixLength = i;
            break;
        }
    }

    const commonPrefix = splitPaths[0].slice(0, prefixLength).join("/");
    const newPaths = splitPaths.map(path => path.slice(prefixLength).join("/")).filter(p => p !== "");

    return { commonPrefix, paths: newPaths };
}