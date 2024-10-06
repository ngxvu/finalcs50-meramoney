import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import './styles.scss';
import logo from '../../assests/images/finalcs50-meramoney.png';

function Dashboard({ totalBalance, totalPeriodExpenses, totalPeriodIncome }) {
    const [selectedRange, setSelectedRange] = useState([new Date(), new Date()]);

    const handleDateChange = (range) => {
        setSelectedRange(range);
        // Add logic to fetch and update the overview data based on the selected date range
    };

    return (
        <>
            <header className="banner">
                <img src={logo} alt="Logo" />
                <span>Meramoney</span>
            </header>
            <div className="dashboard-container">
                <h2>Overview</h2>
                <div className="date-picker-container">
                    <Calendar
                        selectRange
                        onChange={handleDateChange}
                        value={selectedRange}
                    />
                </div>
                <div className="overview">
                    <div className="overview-item">
                        <h3>Total Balance</h3>
                        <p>${totalBalance}</p>
                    </div>
                    <div className="overview-item">
                        <h3>Total Period Expenses</h3>
                        <p>${totalPeriodExpenses}</p>
                    </div>
                    <div className="overview-item">
                        <h3>Total Period Income</h3>
                        <p>${totalPeriodIncome}</p>
                    </div>
                </div>
                <footer>
                    Cs50FinalMeramoney - by Nguyen Xuan Vu
                </footer>
            </div>
        </>
    );
}

Dashboard.propTypes = {
    totalBalance: PropTypes.number.isRequired,
    totalPeriodExpenses: PropTypes.number.isRequired,
    totalPeriodIncome: PropTypes.number.isRequired,
};

export default Dashboard;