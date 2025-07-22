import React, { useEffect, useState } from "react";
import { getBestSellerToday } from "../services/api";

const BestSellerDay = () => {
  const [book, setBook] = useState(null);

  useEffect(() => {
    getBestSellerToday()
      .then((res) => {
        if (res.data.length > 0) setBook(res.data[0]);
      })
      .catch((err) => {
        console.error("Error fetching best seller", err);
      });
  }, []);

  return (
    <div>
      <h2>📅 Best Seller of the Day</h2>
      {book ? (
        <div>
          <p><strong>Title:</strong> {book.title}</p>
          <p><strong>Author:</strong> {book.author}</p>
          <p><strong>Sold:</strong> {book.sold}</p>
        </div>
      ) : (
        <p>No best seller found for today.</p>
      )}
    </div>
  );
};

export default BestSellerDay;
