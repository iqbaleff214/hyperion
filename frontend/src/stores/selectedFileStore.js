import { writable } from "svelte/store";

// Define the selected file store
export const selectedFile = writable(null);
export const selectedFiles = writable([]);
