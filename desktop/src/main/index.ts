import { app, BrowserWindow, ipcMain, dialog } from 'electron'
import { join } from 'path'
import { spawn } from 'child_process'
import { existsSync } from 'fs'

// Resolve the Go CLI binary path (built nexusnode binary sits at repo root)
const CLI_PATH = join(app.getAppPath(), '..', '..', 'nexusnode.exe')

function createWindow(): void {
  const win = new BrowserWindow({
    width: 1200,
    height: 800,
    minWidth: 800,
    minHeight: 600,
    backgroundColor: '#0d1117',
    titleBarStyle: 'hiddenInset',
    webPreferences: {
      preload: join(__dirname, '../preload/index.js'),
      sandbox: false,
      contextIsolation: true,
      nodeIntegration: false
    },
    title: 'NexusNode'
  })

  if (process.env['ELECTRON_RENDERER_URL']) {
    win.loadURL(process.env['ELECTRON_RENDERER_URL'])
    win.webContents.openDevTools()
  } else {
    win.loadFile(join(__dirname, '../renderer/index.html'))
  }
}

app.whenReady().then(() => {
  createWindow()
  app.on('activate', () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow()
  })
})

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit()
})

// ─── IPC: Open file picker ───────────────────────────────────────────────────
ipcMain.handle('dialog:open-file', async () => {
  const result = await dialog.showOpenDialog({ properties: ['openFile'] })
  return result.canceled ? null : result.filePaths[0]
})

// ─── IPC: Split file → calls Go CLI ─────────────────────────────────────────
ipcMain.handle(
  'nexus:split-file',
  (_event, filePath: string, key: string, dataShards: number, parityShards: number) => {
    return new Promise<{ success: boolean; output: string }>((resolve) => {
      if (!existsSync(CLI_PATH)) {
        resolve({
          success: false,
          output: `Go CLI not found at ${CLI_PATH}. Run: go build -o nexusnode cmd/nexusnode/main.go`
        })
        return
      }

      const args = [
        '-chunk', filePath,
        '-key', key,
        '-data', String(dataShards),
        '-parity', String(parityShards)
      ]

      let output = ''
      const proc = spawn(CLI_PATH, args)
      proc.stdout.on('data', (d) => (output += d.toString()))
      proc.stderr.on('data', (d) => (output += d.toString()))
      proc.on('close', (code) => resolve({ success: code === 0, output }))
    })
  }
)

// ─── IPC: Assemble file → calls Go CLI ──────────────────────────────────────
ipcMain.handle(
  'nexus:assemble-file',
  (_event, shards: string, metaPath: string, key: string, outputPath: string) => {
    return new Promise<{ success: boolean; output: string }>((resolve) => {
      if (!existsSync(CLI_PATH)) {
        resolve({ success: false, output: `Go CLI not found at ${CLI_PATH}` })
        return
      }

      const args = [
        '-assemble', shards,
        '-meta', metaPath,
        '-key', key,
        '-out', outputPath
      ]

      let output = ''
      const proc = spawn(CLI_PATH, args)
      proc.stdout.on('data', (d) => (output += d.toString()))
      proc.stderr.on('data', (d) => (output += d.toString()))
      proc.on('close', (code) => resolve({ success: code === 0, output }))
    })
  }
)

// ─── IPC: Mock node status (real DHT integration in Phase 4) ─────────────────
ipcMain.handle('nexus:get-node-status', () => {
  // Returns mock data until real P2P node status streaming is wired in Faz 4
  return {
    totalNodes: 14,
    activeNodes: 11,
    offlineNodes: 3,
    totalShards: 84,
    recoveredShards: 12,
    networkHealth: Math.round((11 / 14) * 100),
    history: Array.from({ length: 10 }, (_, i) => ({
      time: `-${(10 - i) * 30}s`,
      active: 9 + Math.floor(Math.random() * 4)
    }))
  }
})
