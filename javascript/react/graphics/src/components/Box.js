import React, { useState, useEffect, useLayoutEffect, useRef, useContext } from 'react'
import { MouseContext } from '../App.js'
import "./Box.css"

// How many pixels away from the border before we consider it for resizing 
const BORDER_MARGIN = 10

function Box(props) {
  const { width, height, zIndex, focus } = props

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
    cursor: "auto",
  })

  // When we click the mouse down, we need to set the position and set dragging to true
  const handleHeaderMouseDown = (event) => {
    // focus()
    const rect = box.current.getBoundingClientRect()
    setMouseBoxPosition([event.clientX - rect.x, event.clientY - rect.y])
    setDragging(true)
  }

  // When we click on a possible edge to drag...
  const handleBoxMouseDown = () => {
    focus()
    const edgey = getEdgeWithinMargin(mouseX, mouseY) 
    if (edgey) {
      setEdge(edgey)
      setResizing(true)
    } else {
      setEdge("")
      setResizing(false)
    }
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

  // resizes the box based on were the mouse is
  const resizeBoxToMouse = (x, y) => {
    const { 
      width, 
      height,
      bottom,
      left,
      right 
    } = box.current.getBoundingClientRect()

    // What edge are we resizing?
    switch (edge) {
      case "bottom":

        // Set the new height to itself minus the new difference (plus 1 to keep up the cursor)
        box.current.style.height = (height - (bottom - y) + 1) + "px";
        break;
      case "right":

        // Same idea here
        box.current.style.width = (width - (right - x) + 1) + "px";
        break;
      case "bottomright":

        // Run both values in the case of a corner
        box.current.style.height = (height - (bottom - y) + 1) + "px";
        box.current.style.width = (width - (right - x) + 1) + "px";
        break;
      case "left":

        // With the left, we have to shift it over as well
        box.current.style.width = (width + (left - x)) + "px";
        box.current.style.left = x + "px";
        break;
      case "bottomleft":
        box.current.style.height = (height - (bottom - y) + 1) + "px";
        box.current.style.width = (width + (left - x)) + "px";
        box.current.style.left = x + "px";
        break;
    }

  }

  // We when we re-render the application. Make sure the following are being done (if in that given state)
  useLayoutEffect(() => {

    // Move the box to the correct position (if we're dragging)
    if (dragging) {
      moveBoxToMouse(mouseX, mouseY)
      return // If we're dragging, we don't need to do anything else more
    }

    // Change the box size if we're resizing
    if (resizing) {
      resizeBoxToMouse(mouseX, mouseY)
    }

    // Set the cursor when we're just hovering over it
    const edge = getEdgeWithinMargin(mouseX, mouseY)
    if (edge) {
      setCursorToEdge(edge)
    } else {
      setBoxStyling({ ...boxStyling, cursor: "auto"})
    } 

  }, [mouseX, mouseY])

  return (
    <div 
      className="Box"
      ref={box}
      style={{ ...boxStyling, zIndex}}
      // onMouseMove={handleMouseMove}
      onMouseDown={handleBoxMouseDown}
      onMouseUp={() => {
        setEdge("")
        setResizing(false)
      }}
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