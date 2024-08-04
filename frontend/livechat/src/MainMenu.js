import React, { useEffect, useState } from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';


function MainMenu() {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [chats, setChats] = useState([]);
    const [allChats, setAllChats] = useState([]);

    useEffect(() => {
        // Fetch the username from the backend
        const fetchUsername = async () => {
            const token = Cookies.get('token');
            const response = await fetch('http://localhost:8080/me', {
                headers: {
                    'Authorization': `Bearer ${token}`,
                },
            });
            if (response.ok) {
                const data = await response.json();
                setUsername(data.username);
            } else {
                console.error('Failed to fetch username');
            }
        };

        // Fetch the chat names from the backend
        const fetchChats = async () => {
            const token = Cookies.get('token');
            const response = await fetch('http://localhost:8080/user/chats', {
                headers: {
                    'Authorization': `Bearer ${token}`,
                },
            });
            if (response.ok) {
                const chatData = await response.json();
                console.log(chatData);
                if(chatData === null) {
                    setChats([]);
                    return;
                }
                setChats(chatData);
            } else {
                console.error('Failed to fetch chats');
            }
        };

        const fetchAllChats = async () => {
            const response = await fetch('http://localhost:8080/chats/all');
            if (response.ok) {
                const data = await response.json();
                console.log(data);
                if(data === null) {
                    setAllChats([]);
                    return;
                }
                setAllChats(data);
            } else {
                console.error('Failed to fetch all chats');
            }
        }



        fetchUsername();
        fetchChats();
        fetchAllChats();
    }, []);

    return (
        <div className="App">
            <header className="App-header">
                <button className='button-uwu' onClick={() => navigate('/')}>FUCK GO BACK</button>
                <h1>Welcome, {username}</h1>
                <h2>Your Chats</h2>
                <ul className='lists-uwu'>
                {chats && chats.length > 0 ? (
                        chats.map((chat) => (
                            <li key={chat.id} onClick={() => navigate(`/chat/${chat.id}`)}>{chat.name}</li>
                        ))
                    ) : (
                        <li>No chats available</li>
                    )}
                </ul>
                <h2>Other Chats</h2>
                <ul className='lists-uwu'>
                {allChats && allChats.length > 0 ? (
                        allChats.map((chat) => (
                            <li key={chat.id}>{chat.name}</li>
                        ))
                    ) : (
                        <li>There are no chats :c</li>
                    )}
                </ul>
            </header>
        </div>
    );
}

export default MainMenu;