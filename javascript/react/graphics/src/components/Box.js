import React, { useState, useEffect } from 'react'

function Box() {
  
  // Default Box size
  const [boxWidth, setBoxWidth] = useState(100)
  const [boxHeight, setBoxHeight] = useState(100)
  
  // Default box position
  const [boxLeft, setBoxLeft] = useState(0)
  const [boxTop, setBoxTop] = useState(0)

  const [mouseX, setMouseX] = useState(0)
  const [mouseY, setMouseY] = useState(0)

  const [dragging, setDragging] = useState(false)

  // Box styling
  const [boxStyling, setBoxStyling] = useState({
    backgroundColor: "red",
    position: "absolute",
    cursor: "pointer",
    width: `${boxWidth}px`,
    height: `${boxHeight}px`,
    left: `${boxLeft}px`,
    top: `${boxTop}px`,
  })

  // When we change any of the box stylings, make sure to update them
  useEffect(() => {
    setBoxStyling({
      ...boxStyling,
      width: `${boxWidth}px`,
      height: `${boxHeight}px`,
      left: `${boxLeft}px`,
      top: `${boxTop}px`,
    })
    // eslint-disable-next-line
  }, [boxWidth, boxHeight, boxLeft, boxTop])

  // Change the boxes size by the amount of the delta input
  const changeBoxSize = (deltaWidth, deltaHeight) => {
    setBoxWidth(boxWidth + deltaWidth)
    setBoxHeight(boxHeight + deltaHeight)
  }

  // Changes the boxes position by the amount of the delta input
  const changeBoxPosition = (deltaLeft, deltaTop) => {
    setBoxLeft(boxLeft + deltaLeft)
    setBoxTop(boxTop + deltaTop)
  }

  // When
  const moveMouse = (event) => {

    // Set the mouse information including the change
    const xDelta = event.clientX - mouseX;
    const yDelta = event.clientY - mouseY;
    setMouseX(event.clientX)
    setMouseY(event.clientY)

    // If we're in a "dragging" moment then move the box position 
    if (dragging) {
      changeBoxPosition(xDelta, yDelta)
    }
  }

  return (
    <div 
      className="Box" 
      style={boxStyling} 
      onMouseMove={moveMouse}
      onMouseDown={() =>  setDragging(true)} 
      onMouseUp={() => setDragging(false)}>
      Imma box
    </div>
  )
}

export default Box;