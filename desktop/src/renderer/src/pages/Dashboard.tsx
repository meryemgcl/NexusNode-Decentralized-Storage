import { useEffect, useMemo, useState } from 'react'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts'
import { Server, Database, ShieldCheck, AlertTriangle } from 'lucide-react'
import type { NodeStatus } from '../types'

function StatCard({
  icon: Icon,
  label,
  value,
  sub,
  color
}: {
  icon: React.ElementType
  label: string
  value: string | number
  sub?: string
  color: string
}) {
  return (
    <div className="card flex items-start gap-4">
      <div className={`p-2.5 rounded-lg ${color}`}>
        <Icon size={20} />
      </div>
      <div>
        <p className="text-nexus-muted text-xs uppercase tracking-wider">{label}</p>
        <p className="text-2xl font-bold mt-0.5">{value}</p>
        {sub && <p className="text-nexus-muted text-xs mt-1">{sub}</p>}
      </div>
    </div>
  )
}

export default function Dashboard() {
  const [status, setStatus] = useState<NodeStatus | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const load = async () => {
      setLoading(true)
      const s = await window.nexusAPI.getNodeStatus()
      setStatus(s)
      setLoading(false)
    }
    load()
    const interval = setInterval(load, 15_000)
    return () => clearInterval(interval)
  }, [])

  if (loading || !status) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="animate-spin w-8 h-8 border-2 border-nexus-accent border-t-transparent rounded-full" />
      </div>
    )
  }

  const healthColor =
    status.networkHealth >= 80
      ? 'text-nexus-green'
      : status.networkHealth >= 50
      ? 'text-nexus-yellow'
      : 'text-nexus-red'

  // Stabilize random storage display values across re-renders.
  // useMemo ensures these values are computed once per status change, not per render.
  const nodeSizes = useMemo(
    () => Array.from({ length: status.totalNodes }, () => Math.floor(Math.random() * 512 + 64)),
    [status.totalNodes]
  )


  return (
    <div className="p-6 space-y-6">
      {/* Header */}
      <div>
        <h1 className="text-xl font-bold">Dashboard</h1>
        <p className="text-nexus-muted text-sm mt-1">Real-time overview of your NexusNode network</p>
      </div>

      {/* Stat Cards */}
      <div className="grid grid-cols-2 xl:grid-cols-4 gap-4">
        <StatCard
          icon={Server}
          label="Active Nodes"
          value={`${status.activeNodes} / ${status.totalNodes}`}
          sub={`${status.offlineNodes} offline`}
          color="bg-nexus-accent/10 text-nexus-accent"
        />
        <StatCard
          icon={Database}
          label="Total Shards"
          value={status.totalShards}
          sub="distributed across nodes"
          color="bg-nexus-green/10 text-nexus-green"
        />
        <StatCard
          icon={ShieldCheck}
          label="Recovered Shards"
          value={status.recoveredShards}
          sub="via Reed-Solomon parity"
          color="bg-nexus-yellow/10 text-nexus-yellow"
        />
        <StatCard
          icon={AlertTriangle}
          label="Network Health"
          value={`${status.networkHealth}%`}
          sub={status.networkHealth >= 70 ? 'Healthy' : 'Degraded'}
          color={status.networkHealth >= 70 ? 'bg-nexus-green/10 text-nexus-green' : 'bg-nexus-red/10 text-nexus-red'}
        />
      </div>

      {/* Network Health Chart */}
      <div className="card">
        <div className="flex items-center justify-between mb-4">
          <h2 className="font-semibold text-sm">Active Nodes Over Time</h2>
          <span className={`text-xs font-mono font-bold ${healthColor}`}>{status.networkHealth}% healthy</span>
        </div>
        <ResponsiveContainer width="100%" height={200}>
          <LineChart data={status.history}>
            <CartesianGrid strokeDasharray="3 3" stroke="#30363d" />
            <XAxis dataKey="time" tick={{ fill: '#8b949e', fontSize: 11 }} axisLine={false} tickLine={false} />
            <YAxis tick={{ fill: '#8b949e', fontSize: 11 }} axisLine={false} tickLine={false} domain={[0, status.totalNodes]} />
            <Tooltip
              contentStyle={{ background: '#161b22', border: '1px solid #30363d', borderRadius: 8, fontSize: 12 }}
              labelStyle={{ color: '#8b949e' }}
              itemStyle={{ color: '#58a6ff' }}
            />
            <Line
              type="monotone"
              dataKey="active"
              stroke="#58a6ff"
              strokeWidth={2}
              dot={{ r: 3, fill: '#58a6ff', strokeWidth: 0 }}
              activeDot={{ r: 5 }}
            />
          </LineChart>
        </ResponsiveContainer>
      </div>

      {/* Node List */}
      <div className="card">
        <h2 className="font-semibold text-sm mb-4">Node Status</h2>
        <div className="space-y-2">
          {Array.from({ length: status.totalNodes }, (_, i) => {
            const isOnline = i < status.activeNodes
            return (
              <div key={i} className="flex items-center justify-between py-2 border-b border-nexus-border last:border-0">
                <div className="flex items-center gap-3">
                  <div className={`w-2 h-2 rounded-full ${isOnline ? 'bg-nexus-green' : 'bg-nexus-red'}`} />
                  <span className="text-sm font-mono">node-{String(i + 1).padStart(2, '0')}</span>
                </div>
                <div className="flex items-center gap-4">
                  <span className="text-nexus-muted text-xs">{isOnline ? `${nodeSizes[i]} MB` : '—'}</span>
                  <span className={isOnline ? 'badge-green' : 'badge-red'}>
                    {isOnline ? 'online' : 'offline'}
                  </span>
                </div>
              </div>
            )
          })}
        </div>
      </div>
    </div>
  )
}
