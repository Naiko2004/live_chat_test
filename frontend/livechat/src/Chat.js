// src/Chat.js
import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import Cookies from 'js-cookie';
import useWebSocket from './useWebSocket';

function Chat() {
    const { id } = useParams();
    const {messages, sendMessage} = useWebSocket(`ws://localhost:8080/ws/chat/${id}`);
    const [input, setInput] = useState('');
    
    const handleSendMessage = () => {
        const token = Cookies.get('token');
        const message = {
            token: token,
            text: input,
            chat_id: parseInt(id)
        };
        sendMessage(message);
        setInput('');
    }

    return (
        <div className="App">
            <header className="App-header">
                <h1>Chat {id}</h1>
                <div className="chat-box">
                    {messages.map((msg, index) => (
                        <p key={index}>{msg.text}</p>
                    ))}
                </div>
                <input
                    type="text"
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
                />
                <button onClick={handleSendMessage}>Send</button>
            </header>
        </div>
    );
}

export default Chat;