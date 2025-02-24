<script>
  import {
    OpenFiles,
    OpenFolder,
    ReadFilesContent,
    ObfuscateJS,
    ExportEncryptedFiles,
    OpenImportedFileLocation,
  } from "../wailsjs/go/main/App";

  import logo from "/src/assets/images/hyperion-blue.png";

  let selectedFiles = [];
  let filesContent = {};
  let obfuscatedContent = {};
  let showContent = {};
  let isMac = navigator.userAgent.includes("Mac");

  async function openFiles() {
    try {
      const newFiles = await OpenFiles();
      if (newFiles.length > 0) {
        selectedFiles = [...new Set([...selectedFiles, ...newFiles])];
      }
    } catch (error) {
      console.error("Error opening file dialog:", error);
    }
  }

  async function openFolder() {
    try {
      const folderFiles = await OpenFolder();
      if (folderFiles && folderFiles.length > 0) {
        selectedFiles = [...new Set([...selectedFiles, ...folderFiles])];
      }
    } catch (error) {
      console.error("Error opening folder dialog:", error);
    }
  }

  async function obfuscateAll() {
    if (selectedFiles.length === 0) {
      console.warn("No files selected!");
      return;
    }

    try {
      const newFilesContent = await ReadFilesContent(selectedFiles);
      filesContent = { ...filesContent, ...newFilesContent };

      for (let filePath of Object.keys(filesContent)) {
        obfuscatedContent[filePath] = await ObfuscateJS(filesContent[filePath]);
      }
    } catch (error) {
      console.error("Error processing files:", error);
    }
  }

  function removeFile(index) {
    const removedFile = selectedFiles[index];
    selectedFiles = selectedFiles.filter((_, i) => i !== index);

    filesContent = { ...filesContent };
    obfuscatedContent = { ...obfuscatedContent };
    showContent = { ...showContent };

    delete filesContent[removedFile];
    delete obfuscatedContent[removedFile];
    delete showContent[removedFile];
  }

  function removeAllFiles() {
    selectedFiles = [];
    filesContent = {};
    obfuscatedContent = {};
    showContent = {};
  }

  function getFileName(path) {
    return path.split(/[\\/]/).pop();
  }

  function toggleShowCode(filePath) {
    showContent[filePath] = !showContent[filePath];
  }

  async function copyOutput(content) {
    try {
      await navigator.clipboard.writeText(content);
    } catch (err) {
      console.error("Failed to copy:", err);
    }
  }

  async function exportFiles() {
    if (Object.keys(obfuscatedContent).length === 0) {
      console.error("No encrypted files to export!");
      return;
    }
    // @ts-ignore
    await ExportEncryptedFiles(obfuscatedContent);
  }

  async function openImportedFolder(filePath) {
    if (filePath) {
      await OpenImportedFileLocation(filePath);
    }
  }
</script>

<div class="flex flex-col h-full">
  <div style="--wails-draggable:drag; padding-top: {isMac ? '32px' : '8px'}" class="flex flex-wrap gap-1 px-4 pb-2">
    <button
      class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
      on:click={openFiles}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-file-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M14 3v4a1 1 0 0 0 1 1h4"
        /><path
          d="M17 21h-10a2 2 0 0 1 -2 -2v-14a2 2 0 0 1 2 -2h7l5 5v11a2 2 0 0 1 -2 2z"
        /><path d="M10 13l-1 2l1 2" /><path d="M14 13l1 2l-1 2" /></svg
      >
      Open Files
    </button>
    <button
      class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
      on:click={openFolder}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-folder-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M11 19h-6a2 2 0 0 1 -2 -2v-11a2 2 0 0 1 2 -2h4l3 3h7a2 2 0 0 1 2 2v4"
        /><path d="M20 21l2 -2l-2 -2" /><path d="M17 17l-2 2l2 2" /></svg
      >
      Open Folder
    </button>
    <button
      class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
      on:click={removeAllFiles}
      disabled={selectedFiles.length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-copy-minus"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          stroke="none"
          d="M0 0h24v24H0z"
        /><path
          d="M7 9.667a2.667 2.667 0 0 1 2.667 -2.667h8.666a2.667 2.667 0 0 1 2.667 2.667v8.666a2.667 2.667 0 0 1 -2.667 2.667h-8.666a2.667 2.667 0 0 1 -2.667 -2.667z"
        /><path
          d="M4.012 16.737a2 2 0 0 1 -1.012 -1.737v-10c0 -1.1 .9 -2 2 -2h10c.75 0 1.158 .385 1.5 1"
        /><path d="M11 14h6" /></svg
      >
      Remove All
    </button>
    <button
      class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
      on:click={obfuscateAll}
      disabled={selectedFiles.length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-shield-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M12 21a12 12 0 0 1 -8.5 -15a12 12 0 0 0 8.5 -3a12 12 0 0 0 8.5 3a12 12 0 0 1 -.078 7.024"
        /><path d="M20 21l2 -2l-2 -2" /><path d="M17 17l-2 2l2 2" /></svg
      >
      Obfuscate All
    </button>
    <button
      class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
      on:click={exportFiles}
      disabled={Object.keys(obfuscatedContent).length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-upload"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2 -2v-2"
        /><path d="M7 9l5 -5l5 5" /><path d="M12 4l0 12" /></svg
      >
      Export All
    </button>
  </div>
  {#if selectedFiles.length === 0}
  <div class="flex flex-grow w-full p-20 overflow-hidden">
    <img class="mx-auto my-auto object-scale-down max-h-full max-h-full grayscale opacity-25" src={logo} alt="app logo">
  </div>
  {/if}
  <div class="p-4 pt-0 flex flex-col gap-2 overflow-auto flex-grow {selectedFiles.length == 0?'hidden':''}">
    {#if selectedFiles.length > 0}
      <div class="flex flex-col gap-2">
        <ul class="flex flex-col gap-2">
          {#each selectedFiles as file, index}
            <li class="bg-white dark:bg-neutral-900 border border-black/15 dark:border-white/15 dark:text-white p-2 rounded-lg">
              <div class="flex items-center">
                <div class="py-1 pl-1 me-auto">
                  <span class="">{getFileName(file)}</span>
                  <div class="text-xs text-neutral-500">{file}</div>
                </div>
                <button
                  class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-1 rounded enabled:cursor-pointer flex gap-1 items-center"
                  on:click={() => openImportedFolder(file)}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-folder"
                    ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
                      d="M5 4h4l3 3h7a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2v-11a2 2 0 0 1 2 -2"
                    /></svg
                  >
                  Open
                </button>
                <!-- svelte-ignore a11y_consider_explicit_label -->
                <button
                  class="dark:text-white px-3 py-3 cursor-pointer flex gap-1 items-center"
                  on:click={() => removeFile(index)}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-x"
                    ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
                      d="M18 6l-12 12"
                    /><path d="M6 6l12 12" /></svg
                  >
                </button>
              </div>
            </li>
          {/each}
        </ul>
      </div>
    {/if}
  
    {#if Object.keys(filesContent).length > 0}
      <div class="flex flex-col gap-2">
        {#each Object.entries(filesContent) as [filePath, content]}
          <div class="bg-white dark:bg-neutral-900 border border-black/15 dark:border-white/15 dark:text-white rounded-lg">
            <div class="p-4 flex flex-col sm:flex-row gap-2 w-full md:items-center border-b border-black/15">
              <div class="">
                {getFileName(filePath)}
                <div class="text-xs text-neutral-500">{filePath}</div>
              </div>
              <div class="flex ms-auto gap-1">
                <button
                  class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 px-3 py-1 rounded cursor-pointer flex gap-1 items-center"
                  on:click={() => toggleShowCode(filePath)}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-eye"
                    ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
                      d="M10 12a2 2 0 1 0 4 0a2 2 0 0 0 -4 0"
                    /><path
                      d="M21 12c-2.4 4 -5.4 6 -9 6c-3.6 0 -6.6 -2 -9 -6c2.4 -4 5.4 -6 9 -6c3.6 0 6.6 2 9 6"
                    /></svg
                  >
                  {showContent[filePath] ? "Hide Code" : "Show Code"}
                </button>
                <button
                  class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 px-3 py-1 rounded cursor-pointer flex gap-1 items-center"
                  on:click={() => copyOutput(obfuscatedContent[filePath])}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-copy"
                    ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
                      d="M7 7m0 2.667a2.667 2.667 0 0 1 2.667 -2.667h8.666a2.667 2.667 0 0 1 2.667 2.667v8.666a2.667 2.667 0 0 1 -2.667 2.667h-8.666a2.667 2.667 0 0 1 -2.667 -2.667z"
                    /><path
                      d="M4.012 16.737a2.005 2.005 0 0 1 -1.012 -1.737v-10c0 -1.1 .9 -2 2 -2h10c.75 0 1.158 .385 1.5 1"
                    /></svg
                  >
                  Copy Output
                </button>
              </div>
            </div>
            {#if showContent[filePath]}
              <div class="p-4 border-b border-black/15">
                <pre class="whitespace-pre-wrap break-words">{content}</pre>
              </div>
            {/if}
            {#if obfuscatedContent[filePath]}
              <pre
                class="p-4 rounded-lg whitespace-pre-wrap break-words">{obfuscatedContent[
                  filePath
                ]}</pre>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
