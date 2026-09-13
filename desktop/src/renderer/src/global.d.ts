// Global type declarations for the window.nexusAPI exposed via Electron contextBridge
import type { NodeStatus, SplitResult } from './types'

declare global {
  interface Window {
    nexusAPI: {
      openFile: () => Promise<string | null>
      splitFile: (filePath: string, key: string, dataShards: number, parityShards: number) => Promise<SplitResult>
      assembleFile: (shards: string, metaPath: string, key: string, outputPath: string) => Promise<SplitResult>
      getNodeStatus: () => Promise<NodeStatus>
    }
  }
}
