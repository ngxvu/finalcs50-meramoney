import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css';
import Menu from '../../components/menu/menu';
import { Link } from 'react-router-dom'; // Import Link from react-router-dom
import './styles.scss';
import logo from '../../assests/images/finalcs50-meramoney.png';

function Dashboard({ totalBalance, totalPeriodExpenses, totalPeriodIncome }) {
    const [selectedRange, setSelectedRange] = useState([new Date(), new Date()]);
    const [filterType, setFilterType] = useState('all');
    const [transactions, setTransactions] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');


    const handleDateChange = (range) => {
        setSelectedRange(range);
        //
    };

    const handleFilterChange = (event) => {
        setFilterType(event.target.value);
    };

    const handleLogout = () => {
        // Add your logout logic here
        console.log('User logged out');
    };

    const handleAddTransaction = (transaction) => {
        setTransactions([...transactions, transaction]);
    };

    useEffect(() => {
        const fetchTransactions = async () => {
            setLoading(true);
            setError('');
            try {
                const startDate = selectedRange[0].toISOString();
                const endDate = selectedRange[1].toISOString();
                const response = await fetch(`http://localhost:8989/transaction?page=1&pageSize=10&type=${filterType}&start=${startDate}&end=${endDate}`);
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                const data = await response.json();
                setTransactions(data.transactions);
            } catch (error) {
                setError('Failed to fetch transactions. Please try again.');
            } finally {
                setLoading(false);
            }
        };

        fetchTransactions();
    }, [filterType, selectedRange]);

    return (
        <>
            <div className="banner-container">
                <header className="banner">
                    <div className="logo-container">
                        <img src={logo} alt="Logo" />
                        <span className="logo-text">Meramoney</span>
                    </div>
                    <Menu onLogout={handleLogout} /> {/* Add the Menu component */}
                </header>
            </div>
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
                <div className="transactions-filter">
                    <label htmlFor="filter">Filter by:</label>
                    <select id="filter" value={filterType} onChange={handleFilterChange}>
                        <option value="all">All</option>
                        <option value="income">Income</option>
                        <option value="expense">Expense</option>
                    </select>
                </div>
                <div className="transactions-container">
                    <div className="transactions">
                        <h3>Transactions</h3>
                        {loading && <p>Loading transactions...</p>}
                        {error && <p className="error">{error}</p>}
                        <ul>
                            {transactions.map((transaction, index) => (
                                <li key={index} className={`transaction ${transaction.type}`}>
                                    <span>{transaction.date}</span>
                                    <span>{transaction.description}</span>
                                    <span>${transaction.amount}</span>
                                </li>
                            ))}
                        </ul>
                    </div>
                    <div className="add-buttons">
                        <Link to="/add-transaction" className="button">Add Transaction</Link>
                        <Link to="/add-category" className="button">Add Category</Link>
                    </div>
                </div>
            </div>
            <footer>
                    Cs50FinalMeramoney - by Nguyen Xuan Vu
            </footer>
        </>
    );
}

Dashboard.propTypes = {
    totalBalance: PropTypes.number.isRequired,
    totalPeriodExpenses: PropTypes.number.isRequired,
    totalPeriodIncome: PropTypes.number.isRequired,
};

export default Dashboard;