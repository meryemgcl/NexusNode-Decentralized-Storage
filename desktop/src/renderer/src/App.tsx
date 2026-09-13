import { HashRouter, Routes, Route } from 'react-router-dom'
import Sidebar from './components/Sidebar'
import Dashboard from './pages/Dashboard'
import Upload from './pages/Upload'
import Files from './pages/Files'
import Network from './pages/Network'

export default function App() {
  return (
    <HashRouter>
      <div className="flex h-full">
        <Sidebar />
        <main className="flex-1 overflow-y-auto">
          <Routes>
            <Route path="/"        element={<Dashboard />} />
            <Route path="/upload"  element={<Upload />} />
            <Route path="/files"   element={<Files />} />
            <Route path="/network" element={<Network />} />
          </Routes>
        </main>
      </div>
    </HashRouter>
  )
}
