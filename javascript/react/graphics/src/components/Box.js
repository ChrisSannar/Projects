import React, { useState, useEffect, useRef, useContext } from 'react'
import { MouseContext } from '../App.js'
import "./Box.css"

function Box(props) {

  const { width, height } = props

  // Use Ref to keep a shallow render
  const box = useRef(null);

  // Get the global mouse position
  const [mouseX, mouseY] = useContext(MouseContext)

  // Keeping track of the mouse position and state
  const [mouseBoxPosition, setMouseBoxPosition] = useState([0, 0])
  const [dragging, setDragging] = useState(false)

  const [boxStyling] = useState({
    width: (width ?? "0") + "px",
    height: (height ?? "") + "px",
  })

  // Moves the current box to this position
  const moveBoxToMouse = (x, y) => {
    const [mouseBoxX, mouseBoxY] = mouseBoxPosition;
    const left = x - mouseBoxX
    const top = y - mouseBoxY
    box.current.style.left = left + "px"
    box.current.style.top = top + "px"
  }

  // Just make sure whenever we re-render the application that we move the box the the necessary position
  useEffect(() => {
    if (dragging) {
      moveBoxToMouse(mouseX, mouseY)
    }
  })

  return (
    <div 
      className="Box"
      ref={box}
      style={boxStyling}>
        <div 
        className="headerBar"
        onMouseDown={(event) => {
          const rect = box.current.getBoundingClientRect()
          setMouseBoxPosition([event.clientX - rect.x, event.clientY - rect.y])
          setDragging(true)
        }} 
        onMouseUp={() => setDragging(false)}
        onMouseOut={() => {
          console.log("OUT", dragging)
          if (dragging) {
            moveBoxToMouse(mouseX, mouseY)
          } else {
            setDragging(false)
          }
        }}></div>
    </div>
  )
}

export default Box;