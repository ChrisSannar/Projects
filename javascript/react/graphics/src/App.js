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
    }
  ])

  useEffect(() => {
    console.log(boxes)
  }, [])

  // Set all the zIndicies to 0
  const setOnTop = (index) => {
    const newBoxes = [ ...boxes]
    newBoxes.map(box => {
      box.zIndex = 0
      return box
    })
    newBoxes[index].zIndex = 10
    console.log("NEW", newBoxes)
    setBoxes(newBoxes)
  }

  return (
    <div 
    className="App"
    onMouseMove={(e) => setMousePosition([e.clientX, e.clientY])}>
      <MouseContext.Provider value={mousePosition}>
        {boxes.map((box, index) => 
          <Box 
            // zIndex={box.zIndex}
            // width={box.width} 
            // height={box.height} 
            {...box}
            focus={() => setOnTop(index)} />)}
      </MouseContext.Provider>
    </div>
  );
}

export default App;
