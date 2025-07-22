import React from "react";

import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Home from "./pages/Home";
import AddBook from "./pages/AddBook"; // 👈 New
import EditBook from "./pages/EditBook"; 
import BestSellerDay from "./pages/BestSellerDay";
import BestSellerWeek from "./pages/BestSellerWeek";
import BestSellerYear from "./pages/BestSellerYear";


function App() {
  return (
    <Router>
      <nav style={{ marginBottom: "1rem" }}>
        <Link to="/">📚 Home</Link> |{" "}
        <Link to="/add">➕ Add Book</Link> |{" "}
        <Link to="/bestseller-day">📅 Best Seller (Today)</Link>
        <Link to="/bestseller-week">📆 Best Seller (Week)</Link>
        <Link to="/bestseller-year">📈 Best Seller (Year)</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/add" element={<AddBook />} />
        <Route path="/edit/:id" element={<EditBook />} />
        <Route path="/bestseller-day" element={<BestSellerDay />} />
        <Route path="/bestseller-week" element={<BestSellerWeek />} />
        <Route path="/bestseller-year" element={<BestSellerYear />} />
      </Routes>
    </Router>
  );
}

export default App;