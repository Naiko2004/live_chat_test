import React, { useState } from 'react';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router-dom';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        const response = await fetch('http://localhost:8080/login', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            const data = await response.json();
            Cookies.set('token', data.token);
            console.log('Login successful');
        } else {
            console.error('Failed to register');
        }
    }


  return (
    <div className='App'>
        <header className='App-header'>
            <button className='button-uwu' onClick={() => navigate('/')}>FUCK GO BACK</button>
            <h2>Login Page</h2>
            <form onSubmit={handleSubmit}>
                <label htmlFor='username'>Username: </label>
                <input 
                    type='text'
                    id='username'
                    name='username' 
                    value={username}
                    onChange={(event) => setUsername(event.target.value)}
                    />
                <br />
                <label htmlFor='password'>Password: </label>
                <input 
                    type='password' 
                    id='password' 
                    name='password'
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                    />
                <br />
                <br />
                <button type='submit' className='button-uwu'>Login</button>
            </form>
        </header>
    </div>
  );
}

export default Login;