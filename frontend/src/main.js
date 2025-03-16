import { mount } from 'svelte'
import Hyperion from "./Hyperion.svelte";
import './style.css'

const app = mount(Hyperion, {
  target: document.getElementById('app'),
})

export default app