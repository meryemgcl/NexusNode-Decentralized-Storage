import { FileArchive, Info } from 'lucide-react'

// Placeholder for files list — real persistence will be added in Phase 4
const MOCK_FILES = [
  { name: 'report_q3_2026.pdf',    size: '4.2 MB', shards: 14, status: 'distributed', date: '2026-09-13' },
  { name: 'dataset_final.zip',     size: '128 MB', shards: 14, status: 'distributed', date: '2026-09-12' },
  { name: 'presentation_deck.pptx', size: '9.1 MB', shards: 14, status: 'recovering',  date: '2026-09-10' }
]

export default function Files() {
  return (
    <div className="p-6 space-y-6">
      <div>
        <h1 className="text-xl font-bold">Files</h1>
        <p className="text-nexus-muted text-sm mt-1">All files distributed across the network</p>
      </div>

      <div className="card flex items-start gap-3 border-nexus-accent/30">
        <Info size={16} className="text-nexus-accent mt-0.5 flex-shrink-0" />
        <p className="text-xs text-nexus-muted">
          File persistence and metadata indexing will be completed in Phase 4 with DHT integration.
          The entries below are sample data for UI demonstration.
        </p>
      </div>

      <div className="card p-0 overflow-hidden">
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b border-nexus-border text-nexus-muted text-xs uppercase tracking-wider">
              <th className="text-left px-5 py-3">File</th>
              <th className="text-left px-5 py-3">Size</th>
              <th className="text-left px-5 py-3">Shards</th>
              <th className="text-left px-5 py-3">Status</th>
              <th className="text-left px-5 py-3">Date</th>
            </tr>
          </thead>
          <tbody>
            {MOCK_FILES.map((f) => (
              <tr key={f.name} className="border-b border-nexus-border last:border-0 hover:bg-nexus-border/10 transition-colors">
                <td className="px-5 py-3 flex items-center gap-2">
                  <FileArchive size={15} className="text-nexus-muted flex-shrink-0" />
                  <span className="font-mono text-xs">{f.name}</span>
                </td>
                <td className="px-5 py-3 text-nexus-muted text-xs">{f.size}</td>
                <td className="px-5 py-3 text-xs">{f.shards}</td>
                <td className="px-5 py-3">
                  {f.status === 'distributed'
                    ? <span className="badge-green">distributed</span>
                    : <span className="badge-yellow">recovering</span>
                  }
                </td>
                <td className="px-5 py-3 text-nexus-muted text-xs font-mono">{f.date}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
