import React from "react";
import { useNavigate } from "react-router-dom";
import { deleteBook } from "../services/api";

const BookTable = ({ books, refresh }) => {
  const navigate = useNavigate();

  const handleDelete = async (id) => {
    if (window.confirm("Are you sure you want to delete this book?")) {
      try {
        await deleteBook(id);
        alert("❌ Book deleted");
        refresh(); // ⬅️ Tell parent to reload books
      } catch (err) {
        console.error("Error deleting book", err);
        alert("⚠️ Failed to delete book");
      }
    }
  };

  if (!books.length) return <p>No books found.</p>;

  return (
    <table border="1" cellPadding="8">
      <thead>
        <tr>
          <th>ID</th><th>Title</th><th>Author</th><th>Quantity</th><th>Sold</th><th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {books.map((book) => (
          <tr key={book.id}>
            <td>{book.id}</td><td>{book.title}</td><td>{book.author}</td>
            <td>{book.quantity}</td><td>{book.sold}</td>
            <td>
              <button onClick={() => navigate(`/edit/${book.id}`)}>✏️ Edit</button>{" "}
              <button onClick={() => handleDelete(book.id)}>❌ Delete</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default BookTable;
