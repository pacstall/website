// All of the Node.js APIs are available in the preload process.
// It has the same sandbox as a Chrome extension.

const { contextBridge, ipcRenderer } = require('electron')

window.addEventListener('DOMContentLoaded', () => {
  const replaceText = (selector, text) => {
    const element = document.getElementById(selector)
    if (element) element.innerText = text
  }

  for (const type of ['chrome', 'node', 'electron']) {
    replaceText(`${type}-version`, process.versions[type])
  }

  contextBridge.exposeInMainWorld('pacChan', {
    getInstalledPackages: () => ipcRenderer.invoke('getInstalledPackages'),
    getPackageInstalledVersion: (pkgName) => ipcRenderer.invoke('getPackageInstalledVersion', pkgName),
    sendNotification: (title, message) => ipcRenderer.invoke('sendNotification', title, message)
  });
})
