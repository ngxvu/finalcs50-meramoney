import React, { useState } from 'react';
import PropTypes from 'prop-types';
import './styles.scss'; // Import the SCSS file
import logo from '../../assests/images/finalcs50-meramoney.png';

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');


    const handleSubmit = async (event) => {
        event.preventDefault();
        setError('');
        setSuccess('');

        try {
            const response = await fetch('http://localhost:8989/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    user_name: username,
                    password: password,
                }),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            onLogin(data);
            setSuccess('Login successful!'); // Set success message
            
        } catch (error) {
            setError('Login failed. Please check your username and password.');
        }
    };

    return (
        <>
        <header className="banner">
            <img src={logo} alt="Logo" />
            <span>Meramoney</span>
        </header>

        <div className="login-container">
        <h2>Login to Meramoney</h2>
            <p>
                Don't have an account yet? <a href="/signup">Sign up here!</a>
            </p>
                <form onSubmit={handleSubmit}>
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            placeholder="Username"
                            required
                        />
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="Password"
                            required
                        />
                    <button type="submit">Login</button>
                </form>
                
                <footer>
                    Cs50FinalMeramoney - by Nguyen Xuan Vu
                </footer>
        </div>
        </>
    );
}

Login.propTypes = {
    onLogin: PropTypes.func.isRequired,
};

export default Login;