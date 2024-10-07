import React from 'react';
import PropTypes from 'prop-types';
import './styles.scss';

const Loading = ({ message }) => {
    return (
        <div className="loading-container">
            <div className="loading-spinner"></div>
            {message && <p className="loading-message">{message}</p>}
        </div>
    );
};

Loading.propTypes = {
    message: PropTypes.string,
};

export default Loading;