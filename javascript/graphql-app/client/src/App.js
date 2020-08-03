import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';

import axios from 'axios';

function App() {

  // useEffect(() => {
  //   let query = `
  //     query { hello }
  //   `
  //   axios.get('graphql?query=' + query).then(resp => console.log(resp.data));
  // }, []);

  let [message, setMessage] = useState("");
  let [result, setResult] = useState("");

  const updateMessage = (e) => {
    setMessage(e.target.value);
  }

  const sendRequest = () => {

    let query = `query { message(value: "${message}") }`

    axios.get(`graphql?query=` + query)
      .then(resp => {
        let val = resp.data.data.message
        setResult(JSON.stringify(val));
      })
      .catch(err => {
        console.error(err);
      })
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <input onChange={updateMessage} value={message} />
        <button onClick={sendRequest}>Send</button>
        <div>{result}</div>
      </header>
    </div>
  );
}

export default App;
