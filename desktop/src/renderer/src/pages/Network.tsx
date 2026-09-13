import { useEffect, useState } from 'react'
import { RadialBarChart, RadialBar, Legend, ResponsiveContainer, Tooltip } from 'recharts'
import { Wifi, WifiOff } from 'lucide-react'
import type { NodeStatus } from '../types'

export default function Network() {
  const [status, setStatus] = useState<NodeStatus | null>(null)

  useEffect(() => {
    window.nexusAPI.getNodeStatus().then(setStatus)
  }, [])

  if (!status) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="animate-spin w-8 h-8 border-2 border-nexus-accent border-t-transparent rounded-full" />
      </div>
    )
  }

  const chartData = [
    { name: 'Offline', value: status.offlineNodes, fill: '#f85149' },
    { name: 'Active',  value: status.activeNodes,  fill: '#3fb950' }
  ]

  return (
    <div className="p-6 space-y-6">
      <div>
        <h1 className="text-xl font-bold">Network</h1>
        <p className="text-nexus-muted text-sm mt-1">P2P node topology and shard distribution</p>
      </div>

      <div className="grid grid-cols-2 gap-6">
        {/* Radial chart */}
        <div className="card flex flex-col items-center">
          <h2 className="font-semibold text-sm mb-4 self-start">Node Distribution</h2>
          <ResponsiveContainer width="100%" height={240}>
            <RadialBarChart
              cx="50%" cy="50%"
              innerRadius="30%" outerRadius="90%"
              data={chartData}
            >
              <RadialBar dataKey="value" cornerRadius={6} label={false} />
              <Legend iconSize={10} layout="horizontal" verticalAlign="bottom" />
              <Tooltip
                contentStyle={{ background: '#161b22', border: '1px solid #30363d', borderRadius: 8, fontSize: 12 }}
              />
            </RadialBarChart>
          </ResponsiveContainer>
          <p className="text-nexus-muted text-xs mt-2 text-center">
            {status.activeNodes} active / {status.totalNodes} total nodes
          </p>
        </div>

        {/* Node list */}
        <div className="card overflow-y-auto max-h-80">
          <h2 className="font-semibold text-sm mb-3">All Nodes</h2>
          <div className="space-y-2">
            {Array.from({ length: status.totalNodes }, (_, i) => {
              const online = i < status.activeNodes
              return (
                <div key={i} className="flex items-center gap-3 text-xs">
                  {online
                    ? <Wifi size={13} className="text-nexus-green flex-shrink-0" />
                    : <WifiOff size={13} className="text-nexus-red flex-shrink-0" />
                  }
                  <span className="font-mono flex-1">node-{String(i + 1).padStart(2, '0')}</span>
                  <span className={online ? 'badge-green' : 'badge-red'}>
                    {online ? 'online' : 'offline'}
                  </span>
                </div>
              )
            })}
          </div>
        </div>
      </div>

      {/* Shard integrity */}
      <div className="card">
        <h2 className="font-semibold text-sm mb-4">Shard Integrity</h2>
        <div className="grid grid-cols-3 gap-4 text-center">
          <div>
            <p className="text-2xl font-bold text-nexus-accent">{status.totalShards}</p>
            <p className="text-nexus-muted text-xs mt-1">Total Shards</p>
          </div>
          <div>
            <p className="text-2xl font-bold text-nexus-green">{status.totalShards - status.recoveredShards}</p>
            <p className="text-nexus-muted text-xs mt-1">Intact Shards</p>
          </div>
          <div>
            <p className="text-2xl font-bold text-nexus-yellow">{status.recoveredShards}</p>
            <p className="text-nexus-muted text-xs mt-1">Recovered via Parity</p>
          </div>
        </div>
      </div>
    </div>
  )
}
