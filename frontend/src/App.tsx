import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import HelloWorld from './components/HelloWorld'

function App() {
  return (
    <Router>
      <div id="app" className="App">
        <header className="App-header">
          <HelloWorld />
        </header>
      </div>
    </Router>
  )
}

export default App
