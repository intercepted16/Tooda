import './style.css'
import App from './App.svelte'
import { GetWindowsTheme } from '../wailsjs/go/main/App.js'

const app = new App({
  target: document.getElementById('app')
})

GetWindowsTheme().then((theme) => {
  if (theme === 'dark') {
    document.documentElement.classList.add('dark')
  }
})

export default app
