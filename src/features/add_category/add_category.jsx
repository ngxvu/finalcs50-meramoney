import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.scss';

function AddCategory() {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();
        setError('');
        setSuccess('');
        setLoading(true);

        try {
            const response = await fetch('http://143.198.193.9:8989/category', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5ndXllbnh1YW52dSIsImV4cCI6MTcyODQ4NDc4N30.HtqT-VOkZx63QSBYfeZG41HhNgQ61zgdOL9epTN3TY0',
                },
                body: JSON.stringify({
                    name: name,
                    description: description,
                }),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            setSuccess('Category added successfully!');
            setTimeout(() => {
                setLoading(false);
                navigate('/dashboard');
            }, 1000);
        } catch (error) {
            setError('Failed to add category. Please try again.');
            setLoading(false);
        }
    };

    return (
        <div className="add-category-container">
            <h2>Add Category</h2>
            {error && <p className="error">{error}</p>}
            {success && <p className="success">{success}</p>}
            <form onSubmit={handleSubmit}>
                <label htmlFor="name">Category Name:</label>
                <input
                    type="text"
                    id="name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Category Name"
                    required
                />
                <label htmlFor="description">Description:</label>
                <input
                    type="text"
                    id="description"
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    placeholder="Description"
                    required
                />
                <button type="submit" disabled={loading}>
                    {loading ? 'Adding...' : 'Add Category'}
                </button>
            </form>
        </div>
    );
}

export default AddCategory;