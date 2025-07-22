import React, { useEffect, useState } from "react";
import { getBooks } from "../services/api";
import BookTable from "../components/BookTable";

const Home = () => {
  const [books, setBooks] = useState([]);
  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [totalCount, setTotalCount] = useState(0);
  const [searchTerm, setSearchTerm] = useState("");
  const [sort, setSort] = useState("id");         // ✅ NEW
  const [order, setOrder] = useState("desc");      // ✅ NEW

  const loadBooks = () => {
    getBooks(page, limit, searchTerm, sort, order)
      .then((res) => {
        setBooks(res.data.books);
        setTotalCount(res.data.total);
      })
      .catch((err) => console.error("Error fetching books", err));
  };

  useEffect(() => {
    loadBooks();
  }, [page, limit, searchTerm, sort, order]);

  const totalPages = Math.ceil(totalCount / limit);

  return (
    <div>
      <h1>📚 Book Inventory</h1>

      <div style={{ marginBottom: "10px" }}>
        <input
          type="text"
          placeholder="🔍 Search title or author..."
          value={searchTerm}
          onChange={(e) => {
            setPage(1); // reset to first page
            setSearchTerm(e.target.value);
          }}
          style={{ padding: "8px", width: "250px", marginRight: "10px" }}
        />

        Sort by{" "}
        <select value={sort} onChange={(e) => setSort(e.target.value)}>
          <option value="id">ID</option>
          <option value="title">Title</option>
          <option value="author">Author</option>
          <option value="quantity">Quantity</option>
          <option value="sold">Sold</option>
        </select>{" "}

        <select value={order} onChange={(e) => setOrder(e.target.value)}>
          <option value="asc">⬆️ Asc</option>
          <option value="desc">⬇️ Desc</option>
        </select>
      </div>

      <BookTable books={books} refresh={loadBooks} />

      <div style={{ marginTop: "20px" }}>
        <button onClick={() => setPage((p) => Math.max(p - 1, 1))} disabled={page === 1}>
          ◀️ Prev
        </button>

        {" "}Page{" "}
        <select value={page} onChange={(e) => setPage(parseInt(e.target.value))}>
          {Array.from({ length: totalPages }, (_, i) => (
            <option key={i + 1} value={i + 1}>{i + 1}</option>
          ))}
        </select>{" "}
        of {totalPages} ({totalCount} books)

        <button onClick={() => setPage((p) => Math.min(p + 1, totalPages))} disabled={page === totalPages}>
          Next ▶️
        </button>

        {" "} | Show{" "}
        <select value={limit} onChange={(e) => {
          setPage(1);
          setLimit(parseInt(e.target.value));
        }}>
          <option value={10}>10</option>
          <option value={20}>20</option>
          <option value={50}>50</option>
        </select>{" "}
        per page
      </div>
    </div>
  );
};

export default Home;
