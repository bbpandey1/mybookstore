import React, { useState } from "react";
import { addBook } from "../services/api";
import { useNavigate } from "react-router-dom";

const AddBook = () => {
  const [form, setForm] = useState({
    title: "",
    author: "",
    quantity: "",
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const payload = {
      ...form,
      quantity: parseInt(form.quantity), // ✅ convert quantity to number
    };

    try {
      await addBook(payload);
      alert("✅ Book added!");
      navigate("/");
    } catch (err) {
      console.error("Error adding book", err);
      alert("❌ Failed to add book");
    }
  };

  return (
    <div>
      <h2>➕ Add Book</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Title: </label>
          <input
            name="title"
            placeholder="Book title"
            value={form.title}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Author: </label>
          <input
            name="author"
            placeholder="Author name"
            value={form.author}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Quantity: </label>
          <input
            name="quantity"
            type="number"
            placeholder="Number of books"
            value={form.quantity}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" style={{ marginTop: "10px" }}>
          ➕ Add Book
        </button>
      </form>
    </div>
  );
};

export default AddBook;
