import {
    selectedFile,
    selectedFiles,
    fileTree,
} from "./stores/selectedFileStore.js";

import {
    OpenFiles,
    OpenFolder,
    ReadFilesContent,
    ObfuscateJS,
    ExportEncryptedFiles,
    OpenImportedFileLocation,
} from "../wailsjs/go/main/App";

import { get } from "svelte/store";

export async function openFiles() {
    try {
        const newFiles = await OpenFiles();
        if (newFiles.length > 0) {
            selectedFiles.update((files) => [
                ...new Set([...(files || []), ...newFiles]),
            ]);
            fileTree.set(buildFileTree(get(selectedFiles)));
        }
    } catch (error) {
        console.error("Error opening file dialog:", error);
    }
}

export async function openFolder() {
    try {
      const folderFiles = await OpenFolder();
      if (folderFiles && folderFiles.length > 0) {
        selectedFiles.update((files) => {
          const updatedFiles = [...new Set([...(files || []), ...folderFiles])];
          fileTree.set(buildFileTree(updatedFiles));
          return updatedFiles;
        });
      }
    } catch (error) {
      console.error("Error opening folder dialog:", error);
    }
  }

export function buildFileTree(files) {
    const tree = {};

    files.forEach((file) => {
      const normalizedFile = file.replace(/\\/g, "/");
      const parts = normalizedFile.split("/");

      let current = tree;
      parts.forEach((part, index) => {
        if (!current[part]) {
          current[part] =
            index === parts.length - 1 ? { path: normalizedFile } : {};
        }
        current = current[part];
      });
    });

    let keys = Object.keys(tree);
    let trimmedTree = tree;

    while (keys.length === 1) {
      trimmedTree = trimmedTree[keys[0]];
      keys = Object.keys(trimmedTree);
    }

    if (typeof trimmedTree !== "object" || Array.isArray(trimmedTree)) {
      const path = trimmedTree;
      const fileName = trimmedTree.split("/").at(-1);
      console.log(fileName);
      trimmedTree = { [fileName]: { path: path } };
    }

    console.log(trimmedTree);
    return trimmedTree;
  }