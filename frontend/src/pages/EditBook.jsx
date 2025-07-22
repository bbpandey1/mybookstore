import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { getBooks, updateBook } from "../services/api";

const EditBook = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const [form, setForm] = useState({
    title: "",
    author: "",
    quantity: "",
    sold: "",
  });

  useEffect(() => {
    getBooks().then((res) => {
      const book = res.data.find((b) => b.id === parseInt(id));
      if (book) {
        setForm(book);
      } else {
        alert("Book not found");
        navigate("/");
      }
    });
  }, [id, navigate]);

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.name === "quantity" || e.target.name === "sold"
        ? parseInt(e.target.value)
        : e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await updateBook(id, form);
      alert("✅ Book updated!");
      navigate("/");
    } catch (err) {
      console.error("Error updating book", err);
      alert("❌ Failed to update book");
    }
  };

  return (
    <div>
      <h2>✏️ Edit Book</h2>
      <form onSubmit={handleSubmit}>
        <input name="title" value={form.title} onChange={handleChange} required />
        <input name="author" value={form.author} onChange={handleChange} required />
        <input name="quantity" type="number" value={form.quantity} onChange={handleChange} required />
        <input name="sold" type="number" value={form.sold} onChange={handleChange} required />
        <button type="submit" style={{ marginTop: "10px" }}>
          ✅ Update Book
        </button>
      </form>
    </div>
  );
};

export default EditBook;
