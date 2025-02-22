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


  async function readFilesContent() {
    if (selectedFiles.length === 0) {
      console.warn("No files selected!");
      return;
    }

    try {
      const newFilesContent = await ReadFilesContent(selectedFiles);
      filesContent = { ...filesContent, ...newFilesContent };
    } catch (error) {
      console.error("Error reading files:", error);
    }
  }

  async function obfuscateAll() {
    if (Object.keys(filesContent).length === 0) return;

    try {
      for (let filePath of Object.keys(filesContent)) {
        obfuscatedContent[filePath] = await ObfuscateJS(filesContent[filePath]);
      }
    } catch (error) {
      console.error("Error obfuscating files:", error);
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
</script>

<div class="p-4">
  <div class="flex flex-wrap gap-1">
    <button class="bg-blue-200 px-4 py-1 rounded-lg cursor-pointer" on:click={openFiles}>Open Files</button>
    <button class="bg-blue-200 px-4 py-1 rounded-lg cursor-pointer" on:click={openFolder}>Open Folder</button>
    <button class="bg-blue-200 px-4 py-1 rounded-lg cursor-pointer" on:click={readFilesContent} disabled={selectedFiles.length === 0}>Read Files Content</button>
    <button class="bg-red-200 px-4 py-1 rounded-lg cursor-pointer" on:click={removeAllFiles} disabled={selectedFiles.length === 0}>Remove All</button>
    <button class="bg-green-600 text-white px-4 py-1 rounded-lg cursor-pointer" on:click={obfuscateAll} disabled={Object.keys(filesContent).length === 0}>Obfuscate All</button>
  </div>

  {#if selectedFiles.length > 0}
    <div class="mt-4 p-2 bg-gray-200 rounded-lg">
      <h3 class="font-medium">Selected Files:</h3>
      <ul>
        {#each selectedFiles as file, index}
          <li class="flex justify-between items-center bg-white p-2 rounded-lg mt-1">
            <span>{file}</span>
            <button class="bg-red-600 text-sm text-white px-2 py-1 rounded-lg cursor-pointer" on:click={() => removeFile(index)}>Remove</button>
          </li>
        {/each}
      </ul>
    </div>
  {/if}

  {#if Object.keys(filesContent).length > 0}
    <div class="mt-4 p-2 bg-gray-200 rounded-lg">
      {#each Object.entries(filesContent) as [filePath, content]}
        <h3 class="font-medium">{filePath}</h3>
        <pre class="mb-4 whitespace-pre-wrap break-words">{content}</pre>
        {#if obfuscatedContent[filePath]}
          <h4 class="font-medium">Obfuscated:</h4>
          <pre class="mb-4 bg-gray-300 p-2 rounded whitespace-pre-wrap break-words">{obfuscatedContent[filePath]}</pre>
        {/if}
      {/each}
    </div>
  {/if}
</div>
