import React, { useState, createContext } from 'react'
import './App.css';
import Box from './components/Box'

export const MouseContext = createContext(null)

function App() {
  const [mousePosition, setMousePosition] = useState([0, 0])

  return (
    <div className="App" onMouseMove={(e) => setMousePosition([e.clientX, e.clientY])}>
      <MouseContext.Provider value={mousePosition}>
        <Box width="200" height="150" />
        <Box width="400" height="300" />
      </MouseContext.Provider>
    </div>
  );
}

export default App;
