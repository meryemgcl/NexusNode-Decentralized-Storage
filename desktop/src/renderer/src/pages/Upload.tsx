import { useState, useCallback } from 'react'
import { UploadCloud, Key, Layers, CheckCircle2, XCircle, Loader } from 'lucide-react'

type Status = 'idle' | 'loading' | 'success' | 'error'

export default function Upload() {
  const [filePath, setFilePath] = useState<string | null>(null)
  const [key, setKey] = useState('')
  const [dataShards, setDataShards] = useState(10)
  const [parityShards, setParityShards] = useState(4)
  const [status, setStatus] = useState<Status>('idle')
  const [output, setOutput] = useState('')

  const handlePickFile = async () => {
    const path = await window.nexusAPI.openFile()
    if (path) setFilePath(path)
  }

  const handleDrop = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault()
    const file = e.dataTransfer.files[0]
    if (file) setFilePath((file as unknown as { path: string }).path)
  }, [])

  const handleSplit = async () => {
    if (!filePath) return
    if (key.length !== 32) {
      setStatus('error')
      setOutput('Encryption key must be exactly 32 characters.')
      return
    }
    setStatus('loading')
    setOutput('')
    const result = await window.nexusAPI.splitFile(filePath, key, dataShards, parityShards)
    setStatus(result.success ? 'success' : 'error')
    setOutput(result.output)
  }

  return (
    <div className="p-6 space-y-6 max-w-2xl">
      <div>
        <h1 className="text-xl font-bold">Upload File</h1>
        <p className="text-nexus-muted text-sm mt-1">Split and encrypt a file for distributed P2P storage</p>
      </div>

      {/* Drop Zone */}
      <div
        onDrop={handleDrop}
        onDragOver={(e) => e.preventDefault()}
        onClick={handlePickFile}
        className="card border-dashed border-nexus-accent/40 cursor-pointer hover:border-nexus-accent transition-colors flex flex-col items-center justify-center gap-3 py-12"
      >
        <UploadCloud size={36} className="text-nexus-accent" />
        {filePath ? (
          <p className="text-sm font-mono text-nexus-accent break-all text-center">{filePath}</p>
        ) : (
          <>
            <p className="text-sm font-medium">Drop a file here or click to browse</p>
            <p className="text-nexus-muted text-xs">Any file type supported</p>
          </>
        )}
      </div>

      {/* Encryption Key */}
      <div className="card space-y-3">
        <div className="flex items-center gap-2 mb-1">
          <Key size={15} className="text-nexus-muted" />
          <label className="text-sm font-medium">AES-256 Encryption Key</label>
        </div>
        <input
          type="password"
          className="input font-mono"
          placeholder="Enter exactly 32 characters"
          maxLength={32}
          value={key}
          onChange={(e) => setKey(e.target.value)}
        />
        <div className="flex justify-between text-xs text-nexus-muted">
          <span>Key strength: {key.length === 32 ? <span className="text-nexus-green">Ready ✓</span> : `${key.length}/32`}</span>
          <span>AES-256 GCM</span>
        </div>
      </div>

      {/* Shard Settings */}
      <div className="card space-y-4">
        <div className="flex items-center gap-2 mb-1">
          <Layers size={15} className="text-nexus-muted" />
          <span className="text-sm font-medium">Shard Configuration</span>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="text-xs text-nexus-muted block mb-1">Data Shards: <strong className="text-gray-100">{dataShards}</strong></label>
            <input type="range" min={2} max={20} value={dataShards}
              onChange={(e) => setDataShards(Number(e.target.value))}
              className="w-full accent-nexus-accent" />
          </div>
          <div>
            <label className="text-xs text-nexus-muted block mb-1">Parity Shards: <strong className="text-gray-100">{parityShards}</strong></label>
            <input type="range" min={1} max={10} value={parityShards}
              onChange={(e) => setParityShards(Number(e.target.value))}
              className="w-full accent-nexus-accent" />
          </div>
        </div>
        <div className="bg-nexus-bg rounded-lg px-4 py-3 text-xs text-nexus-muted font-mono space-y-1">
          <p>Total shards: <span className="text-gray-100">{dataShards + parityShards}</span></p>
          <p>Fault tolerance: <span className="text-nexus-green">up to {parityShards} node(s) can be offline</span></p>
          <p>Min for recovery: <span className="text-gray-100">{dataShards} shards needed</span></p>
        </div>
      </div>

      {/* Action */}
      <button
        className="btn-primary w-full flex items-center justify-center gap-2"
        onClick={handleSplit}
        disabled={!filePath || key.length !== 32 || status === 'loading'}
      >
        {status === 'loading' ? (
          <><Loader size={16} className="animate-spin" /> Splitting & Encrypting...</>
        ) : (
          <><UploadCloud size={16} /> Split & Distribute</>
        )}
      </button>

      {/* Output */}
      {status !== 'idle' && (
        <div className={`card text-xs font-mono whitespace-pre-wrap ${
          status === 'success' ? 'border-nexus-green/40' : status === 'error' ? 'border-nexus-red/40' : ''
        }`}>
          <div className="flex items-center gap-2 mb-2 text-sm font-sans font-medium">
            {status === 'success' && <><CheckCircle2 size={16} className="text-nexus-green" /> Success</>}
            {status === 'error'   && <><XCircle     size={16} className="text-nexus-red"   /> Error</>}
            {status === 'loading' && <><Loader      size={16} className="animate-spin"     /> Processing...</>}
          </div>
          <p className="text-nexus-muted">{output}</p>
        </div>
      )}
    </div>
  )
}
