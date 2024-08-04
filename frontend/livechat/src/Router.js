// src/Router.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import App from './App';
import Register from './Register';
import Login from './Login';
import MainMenu from './MainMenu';
import Chat from './Chat';

function AppRouter() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<App />} />
                <Route path="/register" element={<Register />} />
                <Route path="/login" element={<Login />} />
                <Route path="/main-menu" element={<MainMenu />} />
                <Route path="/chat/:id" element={<Chat />} />
            </Routes>
        </Router>
    );
}

export default AppRouter;