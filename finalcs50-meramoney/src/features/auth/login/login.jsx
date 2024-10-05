import React, { useState } from 'react';
import PropTypes from 'prop-types';

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        // Call the onLogin prop with username and password
        onLogin({ username, password });
    };

    return (
        <div>
            <header style={{ textAlign: 'center', marginBottom: '20px' }}>
                <img src="/finalcs50-meramoney.png" alt="Logo" style={{ verticalAlign: 'middle', width: '36px', height: 'auto' }} />
                <span style={{ verticalAlign: 'middle', marginLeft: '10px' }}>meramoney</span>
            </header>
            <p style={{ textAlign: 'center' }}>
                Don't have an account yet? <a href="/signup">Sign up here</a>
            </p>
            <div>
                <h2>Login</h2>
                <form onSubmit={handleSubmit}>
                    <div>
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            placeholder="Username"
                            required
                        />
                    </div>
                    <div>
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="Password"
                            required
                        />
                    </div>
                    <button type="submit">Login</button>
                </form>
            </div>
            <footer style={{ textAlign: 'center', marginTop: '20px' }}>
                cs50-final-meramoney - made by Nguyen Xuan Vu
            </footer>
        </div>
    );
}

Login.propTypes = {
    onLogin: PropTypes.func.isRequired,
};

export default Login;