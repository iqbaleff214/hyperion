<script>
  // @ts-nocheck

  import {
    OpenFiles,
    OpenFolder,
    ReadFilesContent,
    ObfuscateJS,
    ExportEncryptedFiles,
    OpenImportedFileLocation,
  } from "../wailsjs/go/main/App";

  import logo from "/src/assets/images/hyperion-blue.png";
  import { writable } from "svelte/store";
  import { obfuscationConfig } from "./configStore";

  let selectedFiles = [];
  let filesContent = {};
  let obfuscatedContent = {};
  let showContent = {};
  let isMac = navigator.userAgent.includes("Mac");
  let isFullscreen = false;

  async function checkFullscreen() {
    isFullscreen = await window.runtime.WindowIsFullscreen();
  }

  window.addEventListener("resize", checkFullscreen);

  setTimeout(() => {
    checkFullscreen();
  }, 500);

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

  function toggleConfig() {
    document.getElementById("config-container").classList.toggle("hidden");
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
        obfuscatedContent[filePath] = await ObfuscateJS(
          filesContent[filePath],
          { ...$obfuscationConfig },
        );
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

<div id="config-container" class="fixed p-2 h-full w-full z-1 flex hidden">
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <div on:click={toggleConfig} class="fixed h-full w-full top-0 left-0"></div>
  <div
    class="relative bg-neutral-100 dark:bg-neutral-800 max-w-full h-full border border-black/15 dark:border-white/15 rounded-lg shadow-xl"
  >
    <div
      class="flex flex-col gap-1 p-2 overflow-auto mt-[40px] h-[calc(100%-40px)]"
    >
      <div class="absolute top-0 left-0 w-full flex items-center">
        <div class="ms-4 text-sm dark:text-white">Obfuscator Config</div>
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <!-- svelte-ignore a11y_click_events_have_key_events -->
        <div
          on:click={toggleConfig}
          class="ms-auto dark:text-white p-3 cursor-pointer"
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
        </div>
      </div>
      <label
        class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-2 rounded-lg cursor-pointer flex gap-1.5 items-center"
      >
        <input
          type="checkbox"
          bind:checked={$obfuscationConfig.removeWhiteSpace}
        />
        Remove Whitespace
      </label>

      <label
        class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-2 rounded-lg cursor-pointer flex gap-1.5 items-center"
      >
        <input
          type="checkbox"
          bind:checked={$obfuscationConfig.removeComments}
        />
        Remove Comments
      </label>

      <label
        class="bg-white dark:bg-neutral-900 text-sm border border-black/15 dark:border-white/15 text-black dark:text-white disabled:text-black/50 dark:disabled:text-white/50 px-3 py-2 rounded-lg cursor-pointer flex gap-1.5 items-center"
      >
        <input
          type="checkbox"
          bind:checked={$obfuscationConfig.changeVariable}
        />
        Change Variable Names
      </label>
    </div>
  </div>
</div>
<div
  class="flex flex-col h-full transition-colors duration-300 {isFullscreen
    ? 'bg-white dark:bg-neutral-900'
    : isMac
      ? 'bg-white/50 dark:bg-neutral-900/50'
      : 'border-t border-black/15 dark:border-white/15'}"
>
<div style="--wails-draggable:drag;" class="{isMac ? 'h-[32px]' : 'hidden'} w-full border-b border-black/15 dark:border-white/15"></div>
<div class="flex">

  <div
    style="--wails-draggable:drag;"
    class="flex flex-col border-r border-black/15 dark:border-white/15 bg-black/10"
  >
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={toggleConfig}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-settings"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065z"
        /><path d="M9 12a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" /></svg
      >
    </button>
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={openFiles}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-file-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M14 3v4a1 1 0 0 0 1 1h4"
        /><path
          d="M17 21h-10a2 2 0 0 1 -2 -2v-14a2 2 0 0 1 2 -2h7l5 5v11a2 2 0 0 1 -2 2z"
        /><path d="M10 13l-1 2l1 2" /><path d="M14 13l1 2l-1 2" /></svg
      >
    </button>
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={openFolder}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-folder-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M11 19h-6a2 2 0 0 1 -2 -2v-11a2 2 0 0 1 2 -2h4l3 3h7a2 2 0 0 1 2 2v4"
        /><path d="M20 21l2 -2l-2 -2" /><path d="M17 17l-2 2l2 2" /></svg
      >
    </button>
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={removeAllFiles}
      disabled={selectedFiles.length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
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
    </button>
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={obfuscateAll}
      disabled={selectedFiles.length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-shield-code"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M12 21a12 12 0 0 1 -8.5 -15a12 12 0 0 0 8.5 -3a12 12 0 0 0 8.5 3a12 12 0 0 1 -.078 7.024"
        /><path d="M20 21l2 -2l-2 -2" /><path d="M17 17l-2 2l2 2" /></svg
      >
    </button>
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="text-black dark:text-white disabled:opacity-50 p-2 enabled:cursor-pointer flex gap-1 items-center"
      on:click={exportFiles}
      disabled={Object.keys(obfuscatedContent).length === 0}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="1.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="icon icon-tabler icons-tabler-outline icon-tabler-upload"
        ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
          d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2 -2v-2"
        /><path d="M7 9l5 -5l5 5" /><path d="M12 4l0 12" /></svg
      >
    </button>
  </div>
  {#if selectedFiles.length === 0}
    <div class="flex flex-grow w-full p-20 overflow-hidden">
      <img
        class="mx-auto my-auto object-scale-down max-h-full max-h-full grayscale opacity-25 pointer-events-none"
        src={logo}
        alt="app logo"
      />
    </div>
  {/if}
  <div
    class="flex flex overflow-hidden grow {selectedFiles.length == 0
      ? 'hidden'
      : ''}"
  >
    {#if selectedFiles.length > 0}
      <ul
        class="flex flex-col min-w-52 shrink-0 bg-black/5 dark:bg-white/5 border-r border-black/15 dark:border-white/15"
      >
        <li class="p-1 px-2 text-sm dark:text-white">
          Imported Files
        </li>
        {#each selectedFiles as file, index}
          <li class="dark:text-white rounded-lg">
            <div class="bg-red-200/0 flex items-center px-2">
              <div class="py-1 me-auto">
                <span class="text-sm">{getFileName(file)}</span>
                <div class="text-xs text-neutral-400 break-words hidden">
                  {file}
                </div>
              </div>
              <!-- svelte-ignore a11y_consider_explicit_label -->
              <button
                class="dark:text-white p-1 cursor-pointer flex gap-1 items-center"
                on:click={() => openImportedFolder(file)}
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
                  class="icon icon-tabler icons-tabler-outline icon-tabler-folder"
                  ><path stroke="none" d="M0 0h24v24H0z" fill="none" /><path
                    d="M5 4h4l3 3h7a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2v-11a2 2 0 0 1 2 -2"
                  /></svg
                >
              </button>
              <!-- svelte-ignore a11y_consider_explicit_label -->
              <button
                class="dark:text-white p-1 cursor-pointer flex gap-1 items-center"
                on:click={() => removeFile(index)}
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
    {/if}

    {#if Object.keys(filesContent).length > 0}
      <div class="flex flex-col gap-2 overflow-x-hidden overflow-y-auto p-2 grow">
        {#each Object.entries(filesContent) as [filePath, content]}
          <div
            class="bg-white dark:bg-neutral-700/50 border border-black/15 dark:border-white/15 dark:text-white rounded-lg"
          >
            <div
              class="p-4 flex flex-col md:flex-row gap-2 w-full md:items-center border-b border-black/15"
            >
              <div class="shrink">
                {getFileName(filePath)}
                <div
                  class="text-xs text-neutral-500 whitespace-pre-wrap break-all"
                >
                  {filePath}
                </div>
              </div>
              <div class="flex ms-auto gap-1 shrink-0 items-start">
                <button
                  class="btn"
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
                  class="btn"
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
</div>
