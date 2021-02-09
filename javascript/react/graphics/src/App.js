import React, { useState, useEffect, createContext } from 'react'
import './App.css';
import Box from './components/Box'

export const MouseContext = createContext(null)

function App() {
  const [mousePosition, setMousePosition] = useState([0, 0])
  const [boxes, setBoxes] = useState([
    {
      width: "200",
      height: "150",
      zIndex: 0
    },
    {
      width: "400",
      height: "300",
      zIndex: 0
    },
    {
      width: "100",
      height: "100",
      zIndex: 0
    }
  ])

  // Set all the zIndicies to 0
  const setOnTop = (index) => {
    const newBoxes = [ ...boxes]

    // Find the largest zIndex and set it +1 to the focused box
    const total = newBoxes.reduce((acc, current) => acc > current.zIndex ? acc : current.zIndex, 0)
    newBoxes[index].zIndex = total + 1
    setBoxes(newBoxes)
  }

  return (
    <div 
    className="App"
    onMouseMove={(e) => setMousePosition([e.clientX, e.clientY])}>
      <MouseContext.Provider value={mousePosition}>
        {boxes.map((box, index) => 
          <Box
            {...box}
            focus={() => setOnTop(index)} />)}
      </MouseContext.Provider>
    </div>
  );
}

export default App;
