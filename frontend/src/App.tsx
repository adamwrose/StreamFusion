import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Dashboard from './components/Dashboard/Dashboard'
import Overlay from './components/Overlay/Overlay'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/overlay"   element={<Overlay />} />
      </Routes>
    </BrowserRouter>
  )
}
