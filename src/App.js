import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import Login from './features/login/login';
import SignUp from './features/signup/sign_up';


function App() {
    const handleLogin = (credentials) => {
        console.log('Login credentials:', credentials);
    };

    const handleSignUp = (credentials) => {
      console.log('Sign-up credentials:', credentials);
    };

    return (
      <Router>
          <div className="App">
              <Routes>
                  <Route path="/login" element={<Login onLogin={handleLogin} />} />
                  <Route path="/signup" element={<SignUp onSignUp={handleSignUp} />} />
                  <Route path="/" element={<Login onLogin={handleLogin} />} />
              </Routes>
          </div>
      </Router>
    );
}

export default App;
