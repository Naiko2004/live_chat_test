import React, { useEffect, useState, useMemo } from 'react';
import './App.css';

function App() {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');
  const ws = useMemo(() => new WebSocket('ws://localhost:8080/ws'), []);

  useEffect(() => {
    ws.onopen = () => {
      console.log('WebSocket connection opened');
    };

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log('Received message:', message);
      setMessages((prevMessages) => {
        console.log('Previous messages:', prevMessages);
        return [...prevMessages, message.text];
      });
    };

    ws.onerror = (error) => {
      console.error('WebSocket error:', error);
    };

    ws.onclose = () => {
      console.log('WebSocket connection closed');
    };
  }, [ws]);

  const sendMessage = () => {
    if (ws.readyState === WebSocket.OPEN) {
      console.log('Sending message:', input);
      const message = {
        text: input,
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI3NjgyOTcsImlkIjoxfQ.zAtU9UBEdUJN8SQjwm4PYJ25g9mjkkeRDGWH8XKK2MQ',
        chat_id: 2 
      }
      ws.send(JSON.stringify(message));
      setInput('');
    } else {
      console.error('WebSocket is not open');
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>Live Chat</h1>
        <div className="chat-box">
          {messages.map((msg, index) => (
            <p key={index}>{msg}</p>
          ))}
        </div>
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button onClick={sendMessage}>Send</button>
      </header>
    </div>
  );
}

export default App;