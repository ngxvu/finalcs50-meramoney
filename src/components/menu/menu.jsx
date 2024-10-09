import React from 'react';
import PropTypes from 'prop-types';
import './styles.scss'; // Import the SCSS file for styling

const Menu = ({ onLogout }) => {
    return (
        <div className="menu">
            <button className="menu-button">☰</button>
            <div className="menu-content">
                <a href="/settings">User Settings</a>
                <button onClick={onLogout}>Logout</button>
            </div>
        </div>
    );
};

Menu.propTypes = {
    onLogout: PropTypes.func.isRequired,
};

export default Menu;