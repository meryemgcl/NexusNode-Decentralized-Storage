import { contextBridge, ipcRenderer } from 'electron'

// Expose a safe, typed API to the renderer process via window.nexusAPI
contextBridge.exposeInMainWorld('nexusAPI', {
  openFile: (): Promise<string | null> =>
    ipcRenderer.invoke('dialog:open-file'),

  splitFile: (
    filePath: string,
    key: string,
    dataShards: number,
    parityShards: number
  ): Promise<{ success: boolean; output: string }> =>
    ipcRenderer.invoke('nexus:split-file', filePath, key, dataShards, parityShards),

  assembleFile: (
    shards: string,
    metaPath: string,
    key: string,
    outputPath: string
  ): Promise<{ success: boolean; output: string }> =>
    ipcRenderer.invoke('nexus:assemble-file', shards, metaPath, key, outputPath),

  getNodeStatus: (): Promise<{
    totalNodes: number
    activeNodes: number
    offlineNodes: number
    totalShards: number
    recoveredShards: number
    networkHealth: number
    history: Array<{ time: string; active: number }>
  }> => ipcRenderer.invoke('nexus:get-node-status')
})
