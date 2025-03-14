
import { 
  openFiles, 
  openFolder, 
  exportFiles, 
  removeAllFiles, 
  unloadSelectedFile,
} from "./appActions.js";

export const menuConfig = [
    {
        name: "File",
        children: [
          { name: "Open", shortcut: "Ctrl+O", action: async () => await openFiles() }, // Ensure async execution
          { name: "Open Folder", shortcut: "Ctrl+Shift+O", action: async () => await openFolder()},
          { name: "Save", shortcut: "Ctrl+S", action: () => console.log("Save") },
          { name: "Save As", shortcut: "Ctrl+Shift+S", action: async () => await exportFiles() },
          { name: "Close", shortcut: "Ctrl+W", action: () => unloadSelectedFile() },
          { name: "Close Folder", shortcut: "Ctrl+Shift+W", action: () => removeAllFiles() }
        ]
      },
    {
      name: "Run",
      children: [
        { name: "Obfuscate", action: () => console.log("Obfuscate") },
        { name: "Obfuscate All", action: () => console.log("Obfuscate All") },
        { name: "Configuration", action: () => console.log("Configuration") },
        { name: "Single Line", action: () => console.log("Single Line") },
        { name: "String Literal", action: () => console.log("String Literal") },
        { name: "Loop Statement", action: () => console.log("Loop Statement") },
        { name: "If Statement", action: () => console.log("If Statement") },
        { name: "Constant Name", action: () => console.log("Constant Name") },
        { name: "Variable Name", action: () => console.log("Variable Name") },
        { name: "Function Name", action: () => console.log("Function Name") },
        { name: "Remove Comments", action: () => console.log("Remove Comments") }
      ]
    },
    {
      name: "Help",
      children: [
        { name: "How to Contribute", action: () => console.log("How to Contribute") },
        { name: "Report Issue", action: () => console.log("Report Issue") }
      ]
    }
  ];
  