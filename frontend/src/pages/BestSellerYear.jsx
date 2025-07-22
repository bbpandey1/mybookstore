import React, { useEffect, useState } from "react";
import { getBestSellerYear } from "../services/api";

const BestSellerYear = () => {
  const [book, setBook] = useState(null);

  useEffect(() => {
    getBestSellerYear()
      .then((res) => {
        if (res.data.length > 0) setBook(res.data[0]);
      })
      .catch((err) => {
        console.error("Error fetching best seller (year)", err);
      });
  }, []);

  return (
    <div>
      <h2>📈 Best Seller of the Year</h2>
      {book ? (
        <div>
          <p><strong>Title:</strong> {book.title}</p>
          <p><strong>Author:</strong> {book.author}</p>
          <p><strong>Sold:</strong> {book.sold}</p>
        </div>
      ) : (
        <p>No best seller found for this year.</p>
      )}
    </div>
  );
};

export default BestSellerYear;
