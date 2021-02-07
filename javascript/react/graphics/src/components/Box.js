import React, { useState, useEffect } from 'react'

function Box() {
  const [boxWidth, setBoxWidth] = useState(100)
  const [boxHeight, setBoxHeight] = useState(100)
  
  const [boxLeft, setBoxLeft] = useState(0)
  const [boxTop, setBoxTop] = useState(0)

  // const [dragging, setDragging] = useState(false)
  const [boxStyling, setBoxStyling] = useState({
    backgroundColor: "red",
    width: `${boxWidth}px`,
    height: `${boxHeight}px`,
    position: "absolute",
    left: `10px`,
    top: `10px`,
    // left: `${boxLeft}px`,
    // top: `${boxTop}px`,
  })

  useEffect(() => {
    setBoxStyling({
      ...boxStyling,
      width: `${boxWidth}px`,
      height: `${boxHeight}px`,
      left: `${boxLeft}px`,
      top: `${boxTop}px`,
    })
  }, [boxWidth, boxHeight, boxLeft, boxTop])

  const boxClick = (event) => {
    changeBoxSize(10, 10)
    setBoxLeft(boxLeft + 10)
    setBoxTop(boxTop + 10)
  }

  const changeBoxSize = (deltaWidth, deltaHeight) => {
    setBoxWidth(boxWidth + deltaWidth)
    setBoxHeight(boxHeight + deltaHeight)
  }

  return (
    <div className="Box" style={boxStyling} onClick={boxClick}>
      Imma box
    </div>
  )
}

export default Box;