<script>
  import { OpenFiles, OpenFolder, ReadFilesContent, ObfuscateJS } from '../wailsjs/go/main/App';

  let selectedFiles = [];
  let filesContent = {};
  let obfuscatedContent = {};

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
    delete filesContent[removedFile];
    delete obfuscatedContent[removedFile];
  }

  function removeAllFiles() {
    selectedFiles = [];
    filesContent = {};
    obfuscatedContent = {};
  }

  function getFileName(path) {
    return path.split(/[\\/]/).pop();
  }
</script>

<div class="p-4 flex flex-col gap-2">
  <div class="flex flex-wrap gap-1">
    <button class="bg-blue-200 px-4 py-1 rounded-lg cursor-pointer" on:click={openFiles}>Open Files</button>
    <button class="bg-blue-200 px-4 py-1 rounded-lg cursor-pointer" on:click={openFolder}>Open Folder</button>
    <button class="bg-red-200 px-4 py-1 rounded-lg cursor-pointer" on:click={removeAllFiles} disabled={selectedFiles.length === 0}>Remove All</button>
    <button class="bg-green-600 text-white px-4 py-1 rounded-lg cursor-pointer" on:click={obfuscateAll} disabled={selectedFiles.length === 0}>Obfuscate All</button>
  </div>

  {#if selectedFiles.length > 0}
    <div class="flex flex-col gap-2">
      <h3 class="font-medium">Selected Files:</h3>
      <ul class="flex flex-col gap-2">
        {#each selectedFiles as file, index}
          <li class="bg-white p-2 rounded-xl">
            <div class="flex justify-between items-start">
              <div class="py-1 pl-1">
                <span class="font-semibold">{getFileName(file)}</span>
                <div class="text-xs text-gray-500">{file}</div>
              </div>
              <button class="bg-red-600 text-sm text-white px-2 py-1 rounded-lg cursor-pointer" on:click={() => removeFile(index)}>Remove</button>
            </div>
          </li>
        {/each}
      </ul>
    </div>
  {/if}

  {#if Object.keys(filesContent).length > 0}
    <div class="flex flex-col gap-2">
      {#each Object.entries(filesContent) as [filePath, content]}
        <div class="font-medium">{getFileName(filePath)}</div>
        <pre class="bg-white p-4 rounded-xl whitespace-pre-wrap break-words">{content}</pre>
        {#if obfuscatedContent[filePath]}
          <div class="font-medium">Obfuscated:</div>
          <pre class="bg-white p-4 rounded-xl whitespace-pre-wrap break-words">{obfuscatedContent[filePath]}</pre>
        {/if}
      {/each}
    </div>
  {/if}
</div>
