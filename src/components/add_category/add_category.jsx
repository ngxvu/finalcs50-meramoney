import React, { useState } from 'react';

function AddCategory({ onAddCategory }) {
    const [category, setCategory] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onAddCategory(category);
        setCategory('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <h3>Add Category</h3>
            <div>
                <label>Category Name:</label>
                <input
                    type="text"
                    value={category}
                    onChange={(e) => setCategory(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Add Category</button>
        </form>
    );
}

export default AddCategory;