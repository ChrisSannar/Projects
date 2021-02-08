import React, { useState, useEffect, useLayoutEffect, useRef, useContext } from 'react'
import { MouseContext } from '../App.js'
import "./Box.css"

// How many pixels away from the border before we consider it for resizing 
const BORDER_MARGIN = 10

function Box(props) {
  const { width, height } = props

  // Use Ref to keep a shallow render
  const box = useRef(null);

  // Get the global mouse position
  const [mouseX, mouseY] = useContext(MouseContext)

  // Keeping track of the mouse position and state
  const [mouseBoxPosition, setMouseBoxPosition] = useState([0, 0])
  const [dragging, setDragging] = useState(false)
  const [resizing, setResizing] = useState(false)
  const [edge, setEdge] = useState("")
  const [boxStyling, setBoxStyling] = useState({
    width: (width ?? "0") + "px",
    height: (height ?? "") + "px",
    cursor: "auto"
  })

  // When we click the mouse down, we need to set the position and set dragging to true
  const handleHeaderMouseDown = (event) => {
    const rect = box.current.getBoundingClientRect()
    setMouseBoxPosition([event.clientX - rect.x, event.clientY - rect.y])
    setDragging(true)
  }

  // Make sure when we go out but are still dragging, we move the mouse with it
  const handleMouseOut = () => {
    if (dragging) {
      moveBoxToMouse(mouseX, mouseY)
    } else {
      setDragging(false)
    }
  }

  // Moves the current box to this position (relative to where the mouse is)
  const moveBoxToMouse = (x, y) => {
    const [mouseBoxX, mouseBoxY] = mouseBoxPosition;
    const left = x - mouseBoxX
    const top = y - mouseBoxY
    box.current.style.left = left + "px"
    box.current.style.top = top + "px"
  }

  // Checks if the current x/y position is near the box (within the BORDER_MARGIN)
  const nearBox = (x, y) => {
    const rect = box.current.getBoundingClientRect()
    return x > (rect.left - BORDER_MARGIN) &&
    x < (rect.right + BORDER_MARGIN) &&
    y > (rect.top - BORDER_MARGIN) &&
    y < (rect.bottom + BORDER_MARGIN)
  }

  // Determines if the edge is close enough to consider "editing"
  const getEdgeWithinMargin = (x, y) => {
    let result = ""
    if (nearBox(x, y)) {
      const rect = box.current.getBoundingClientRect()
      // if (Math.abs(y - rect.top) < BORDER_MARGIN) {
      //   result += "top"
      // }
      if (Math.abs(y - rect.bottom) < BORDER_MARGIN) {
        result += "bottom"
      }
      if (Math.abs(x - rect.left) < BORDER_MARGIN) {
        result += "left"
      }
      if (Math.abs(x - rect.right) < BORDER_MARGIN) {
        result += "right"
      }
    }
    return result;
  }

  // Sets the cursor and the current resize configuration
  const setCursorToEdge = (edge) => {
    switch (edge) {
      case "top":
      case "bottom":
        setBoxStyling({ ...boxStyling, cursor: "ns-resize"})
        break;
      case "left":
      case "right":
        setBoxStyling({ ...boxStyling, cursor: "ew-resize"})
        break;
      case "bottomleft":
      case "topright":
        setBoxStyling({ ...boxStyling, cursor: "nesw-resize"})
        break;
      case "topleft":
      case "bottomright":
        setBoxStyling({ ...boxStyling, cursor: "nwse-resize"})
        break;
      default:
        setBoxStyling({ ...boxStyling, cursor: "auto"})
        break;
    }
  }

  // const [mouseBoxX, mouseBoxY] = mouseBoxPosition;
  //   const left = x - mouseBoxX
  //   const top = y - mouseBoxY
  //   box.current.style.left = left + "px"
  //   box.current.style.top = top + "px"

  // resizes the box based on were the mouse is
  const resizeBoxToMouse = (edge, x, y) => {
    const { 
      width, 
      height,
      top,
      bottom,
      left,
      right 
    } = box.current.getBoundingClientRect()
    console.log(edge, height, bottom, y)

  }

  // Just make sure whenever we re-render the application that we move the box to the necessary position
  useLayoutEffect(() => {

    // Move the box if we're dragging
    if (dragging) {
      moveBoxToMouse(mouseX, mouseY)
    } else {
      const edge = getEdgeWithinMargin(mouseX, mouseY)
      setEdge(edge)
      if (edge) {
        setCursorToEdge(edge)
      } else {
        setBoxStyling({ ...boxStyling, cursor: "auto"})
      }
    }

    if (resizing && edge) {
      resizeBoxToMouse(edge, mouseX, mouseY)
    }
  }, [mouseX, mouseY])

  return (
    <div 
      className="Box"
      ref={box}
      style={boxStyling}
      // onMouseMove={handleMouseMove}
      onMouseDown={() => setResizing(!dragging)}
      onMouseUp={() => setResizing(false)}
      >
        <div 
        className="headerBar"
        style={{ cursor: dragging ? "grabbing" : "grab"}}
        onMouseDown={handleHeaderMouseDown} 
        onMouseUp={() => setDragging(false)}
        onMouseOut={handleMouseOut}></div>
    </div>
  )
}

export default Box;