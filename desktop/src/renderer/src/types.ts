export interface NodeStatus {
  totalNodes: number
  activeNodes: number
  offlineNodes: number
  totalShards: number
  recoveredShards: number
  networkHealth: number
  history: Array<{ time: string; active: number }>
}

export interface SplitResult {
  success: boolean
  output: string
}
