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
      // sandbox: true is the secure default — contextBridge still works correctly.
      // Never set sandbox: false unless you have a documented, unavoidable reason.
      sandbox: true,
      contextIsolation: true,
      nodeIntegration: false
    },
    title: 'NexusNode'
  })

  if (process.env['ELECTRON_RENDERER_URL']) {
    win.loadURL(process.env['ELECTRON_RENDERER_URL'])
    // Only open DevTools in development
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

// ─── Helper: Spawn Go CLI with the encryption key delivered via stdin ─────────
// SECURITY: The AES key is written to the process's stdin pipe, NOT passed as a
// command-line argument. This prevents the key from appearing in `ps aux`,
// Windows Task Manager process lists, or shell history.
function spawnCLI(
  args: string[],
  key: string
): Promise<{ success: boolean; output: string }> {
  return new Promise((resolve) => {
    if (!existsSync(CLI_PATH)) {
      resolve({
        success: false,
        output: [
          `Go CLI binary not found at: ${CLI_PATH}`,
          'Build it first:',
          '  go build -o nexusnode.exe cmd/nexusnode/main.go'
        ].join('\n')
      })
      return
    }

    let output = ''
    const proc = spawn(CLI_PATH, args, { stdio: ['pipe', 'pipe', 'pipe'] })

    proc.stdout.on('data', (d: Buffer) => (output += d.toString()))
    proc.stderr.on('data', (d: Buffer) => (output += d.toString()))
    proc.on('close', (code: number | null) => resolve({ success: code === 0, output }))
    proc.on('error', (err: Error) => resolve({ success: false, output: err.message }))

    // Write the key to stdin then close it — the Go CLI reads key from stdin.
    proc.stdin.write(key + '\n')
    proc.stdin.end()
  })
}

// ─── IPC: Split file → calls Go CLI ─────────────────────────────────────────
ipcMain.handle(
  'nexus:split-file',
  (_event, filePath: string, key: string, dataShards: number, parityShards: number) => {
    const args = [
      '-chunk', filePath,
      '-data', String(dataShards),
      '-parity', String(parityShards)
      // NOTE: -key is intentionally omitted; key is delivered via stdin
    ]
    return spawnCLI(args, key)
  }
)

// ─── IPC: Assemble file → calls Go CLI ──────────────────────────────────────
ipcMain.handle(
  'nexus:assemble-file',
  (_event, shards: string, metaPath: string, key: string, outputPath: string) => {
    const args = [
      '-assemble', shards,
      '-meta', metaPath,
      '-out', outputPath
      // NOTE: -key is intentionally omitted; key is delivered via stdin
    ]
    return spawnCLI(args, key)
  }
)

// ─── IPC: Mock node status (real DHT integration in Phase 4) ─────────────────
ipcMain.handle('nexus:get-node-status', () => {
  // Returns mock data until real P2P node status streaming is wired in Phase 4
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
