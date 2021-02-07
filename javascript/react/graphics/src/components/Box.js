import React, { useState, useEffect } from 'react'

function Box() {
  const [boxWidth, setBoxWidth] = useState(100)
  const [boxHeight, setBoxHeight] = useState(100)
  const [boxStyling, setBoxStyling] = useState({
    width: `${boxWidth}px`,
    height: `${boxHeight}px`,
    backgroundColor: "red"
  })

  useEffect(() => {
    setBoxStyling({
      ...boxStyling,
      width: `${boxWidth}px`,
      height: `${boxHeight}px`,
    })
  }, [boxWidth, boxHeight])

  const boxClick = (event) => {
    changeBoxSize(10, 10)
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