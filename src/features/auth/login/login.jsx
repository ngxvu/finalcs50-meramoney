import React, { useState } from 'react';
import PropTypes from 'prop-types';
import './styles.scss'; // Import the SCSS file

function Login({ onLogin }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        // Call the onLogin prop with username and password
        onLogin({ username, password });
    };

    return (
        <div className="login-container">
            <header>
                <img src="/finalcs50-meramoney.png" alt="Logo" />
                <span>meramoney</span>
            </header>
            <p>
                Don't have an account yet? <a href="/signup">Sign up here</a>
            </p>
            <div className="form-container">
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
            <footer>
                cs50-final-meramoney - made by Nguyen Xuan Vu
            </footer>
        </div>
    );
}

Login.propTypes = {
    onLogin: PropTypes.func.isRequired,
};

export default Login;