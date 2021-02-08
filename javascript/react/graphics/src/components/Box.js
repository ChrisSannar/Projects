import React, { useState, useEffect, useRef } from 'react'
import "./Box.css"

function Box() {

  // Use Ref to keep a shallow render
  const box = useRef(null);

  const [mouseX, setMouseX] = useState(0)
  const [mouseY, setMouseY] = useState(0)

  const [dragging, setDragging] = useState(false)

  // When we move the mouse over our box
  const moveMouse = (event) => {

    // Set the mouse information including the change in direction (movenmentX/Y isn't accurate enough)
    const xDelta = event.clientX - mouseX;
    const yDelta = event.clientY - mouseY;
    setMouseX(event.clientX)
    setMouseY(event.clientY)

    // If we're in a "dragging" moment then move the box position 
    if (dragging) {
      const top = extractPxNum(box.current.style.top)
      const left = extractPxNum(box.current.style.left)
      box.current.style.top = top + yDelta + "px"
      box.current.style.left = left + xDelta + "px"
    }
  }

  // Given a css property with a px value, returns the number from the px value
  const extractPxNum = (value) => {
    return Number(value.replace("px", ""))
  }

  return (
    <div 
      className="Box"
      ref={box}
      onMouseMove={moveMouse}
      onMouseDown={() =>  setDragging(true)} 
      onMouseUp={() => setDragging(false)}
      onMouseOut={() => setDragging(false)}>
      Imma box
    </div>
  )
}

export default Box;