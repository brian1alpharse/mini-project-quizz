import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import { AllUser, Login } from './Pages';
import './App.css';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false); // State untuk menentukan status login

  useEffect(() => {
    // Cek apakah token ada di local storage (artinya pengguna sudah login)
    const token = localStorage.getItem('token');
    if (token) {
      setIsLoggedIn(true);
    } else {
      setIsLoggedIn(false);
    }
  }, []);

  return (
    <Router>
      <div className="App">
        <h1>Data from API:</h1>
        <Routes>
          {/* Route untuk halaman login */}
          <Route path="/login" element={<Login />} />
          {/* Route untuk halaman allUser */}
          <Route path="/user" element={isLoggedIn ? <AllUser /> : <Navigate to="/login" />} />
          {/* Redirect dari root path ke halaman login */}
          <Route path="/" element={<Navigate to="/login" />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
