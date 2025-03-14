import { writable } from "svelte/store";

// Define the selected file store
export const selectedFile = writable(null);
export const selectedFiles = writable([]);
export const fileTree = writable(null);
export const filesContent = writable({});
export const obfuscatedContent = writable({});
export const showContent = writable({});