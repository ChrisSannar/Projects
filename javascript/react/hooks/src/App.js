import React, { useState, useEffect } from 'react';
import './App.css';

const Temp = () => {
  return <p>temp</p>;
}

function App() {
  // passing in an function only renders it the first time
  const [countState, setCountState] = useState(() => expensiveComputation());
  // Multiple states can be used for different attribtutes
  const [otherState, setOtherState] = useState("a");

  // Mounting
  useEffect(() => {
    console.log('Effect Mounting');
  }, []);

  // 'otherState' change
  useEffect(() => {

    console.log('Other State Effect');

    // When after the effect goes
    return () => {
      console.log('Effect Cleanup');
    }
  }, [otherState])

  return (
    <div className="App">

      {/* pass in a function for 'setCount' that has the state as a paramater */}
      <button onClick={() => setCountState(countState => ({
        ...countState,  // Need to set populate in all the other fields
        count: countState.count + 1
      }))}>+</button>
      <p>{countState.count}</p>
      <button onClick={() => setOtherState(state => state + 'a')}>tick</button>
      <p>{otherState}</p>
      <Temp />
    </div>
  );
}

// pretend this takes time to render
function expensiveComputation() {
  // ...
  return { count: 10 };
}

export default App;
