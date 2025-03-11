import { writable } from "svelte/store";

const storedConfig = JSON.parse(localStorage.getItem("obfuscationConfig")) || {
  removeWhiteSpace: true,
  removeComments: true,
  changeVariable: true,
};

export const obfuscationConfig = writable(storedConfig);

obfuscationConfig.subscribe((value) => {
  localStorage.setItem("obfuscationConfig", JSON.stringify(value));
});
