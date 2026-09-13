import { NavLink } from 'react-router-dom'
import { LayoutDashboard, Upload, FolderOpen, Network } from 'lucide-react'

const links = [
  { to: '/',        label: 'Dashboard', icon: LayoutDashboard },
  { to: '/upload',  label: 'Upload',    icon: Upload          },
  { to: '/files',   label: 'Files',     icon: FolderOpen      },
  { to: '/network', label: 'Network',   icon: Network         }
]

export default function Sidebar() {
  return (
    <aside className="w-56 flex-shrink-0 bg-nexus-surface border-r border-nexus-border flex flex-col">
      {/* Logo */}
      <div className="px-5 py-5 border-b border-nexus-border">
        <div className="flex items-center gap-2">
          <div className="w-7 h-7 rounded-lg bg-nexus-accent flex items-center justify-center">
            <Network size={16} className="text-nexus-bg" />
          </div>
          <span className="font-bold text-sm tracking-wide">NexusNode</span>
        </div>
        <p className="text-nexus-muted text-xs mt-1">Decentralized Storage</p>
      </div>

      {/* Navigation */}
      <nav className="flex-1 px-3 py-4 space-y-1">
        {links.map(({ to, label, icon: Icon }) => (
          <NavLink
            key={to}
            to={to}
            end={to === '/'}
            className={({ isActive }) =>
              `flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors duration-150 ${
                isActive
                  ? 'bg-nexus-accent/10 text-nexus-accent font-medium'
                  : 'text-nexus-muted hover:text-gray-100 hover:bg-nexus-border/30'
              }`
            }
          >
            <Icon size={16} />
            {label}
          </NavLink>
        ))}
      </nav>

      {/* Footer */}
      <div className="px-5 py-4 border-t border-nexus-border">
        <p className="text-nexus-muted text-xs">v0.1.0 — Phase 3</p>
      </div>
    </aside>
  )
}
