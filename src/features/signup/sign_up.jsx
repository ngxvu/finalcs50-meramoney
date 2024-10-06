import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import PropTypes from 'prop-types';
import './styles.scss';
import logo from '../../assests/images/finalcs50-meramoney.png';

function SignUp({ onSignUp }) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const history = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        setError('');
        setSuccess('');

        if (password !== confirmPassword) {
            setError('Passwords do not match.');
            return;
        }

        try {
            const response = await fetch('http://localhost:8989/signup', {
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
            onSignUp(data);
            setSuccess('Sign-up successful! Redirecting to login page...');
            setTimeout(() => {
                history.push('/login');
            }, 2000); // Redirect after 2 seconds
        } catch (error) {
            setError('Sign-up failed. Please try again.');
        }
    };

    return (
        <>
            <header className="banner">
                <img src={logo} alt="Logo" />
                <span>Meramoney</span>
            </header>
            <div className="signup-container">
                <p>
                    Already have an account? <a href="/login">Login here</a>
                </p>
                <h2>Sign Up</h2>
                {error && <p className="error">{error}</p>} {/* Display error message */}
                {success && <p className="success">{success}</p>} {/* Display success message */}
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
                    <label htmlFor="confirmPassword">Confirm Password:</label>
                    <input
                        type="password"
                        id="confirmPassword"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                        placeholder="Confirm Password"
                        required
                    />
                    <button type="submit">Sign Up</button>
                </form>
                <footer>
                   Cs50FinalMeramoney - by Nguyen Xuan Vu
                </footer>
            </div>
        </>
    );
}

SignUp.propTypes = {
    onSignUp: PropTypes.func.isRequired,
};

export default SignUp;