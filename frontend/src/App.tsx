import { Routes, Route } from 'react-router-dom'
import Home from './components/pages/Home/Home'
import About from './components/pages/About/About'
import { Nominations } from './components/pages/Nominations/Nominations'

function App() {
  return (
    <div className="min-h-screen">
      <main>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/nominations" element={<Nominations />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </main>
    </div>
  )
}

export default App
