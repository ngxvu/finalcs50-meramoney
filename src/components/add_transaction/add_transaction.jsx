import React, { useState } from 'react';

function AddTransaction({ onAddTransaction }) {
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');
    const [type, setType] = useState('income');

    const handleSubmit = (e) => {
        e.preventDefault();
        onAddTransaction({ description, amount: parseFloat(amount), type });
        setDescription('');
        setAmount('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <h3>Add Transaction</h3>
            <div>
                <label>Description:</label>
                <input
                    type="text"
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Amount:</label>
                <input
                    type="number"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Type:</label>
                <select value={type} onChange={(e) => setType(e.target.value)}>
                    <option value="income">Income</option>
                    <option value="expense">Expense</option>
                </select>
            </div>
            <button type="submit">Add Transaction</button>
        </form>
    );
}

export default AddTransaction;