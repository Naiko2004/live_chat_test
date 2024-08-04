// src/App.js
import React from 'react';
import { useNavigate } from 'react-router-dom';
import './App.css';

function App() {
  const navigate = useNavigate();

  return (
    <div className="App">
      <header className="App-header">
        <h1>Live Chat Test</h1>
        <div className='button-container'>
          <button onClick={() => navigate('/register')}>Register</button>
          <button onClick={() => navigate('/login')}>Login</button>
          <button onClick={() => navigate('/main-menu')}>Main Menu</button>
        </div>
      </header>
    </div>
  );
}

export default App;