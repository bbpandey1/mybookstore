import React, { useEffect, useState } from "react";
import { getBestSellerWeek } from "../services/api";

const BestSellerWeek = () => {
  const [book, setBook] = useState(null);

  useEffect(() => {
    getBestSellerWeek()
      .then((res) => {
        if (res.data.length > 0) setBook(res.data[0]);
      })
      .catch((err) => {
        console.error("Error fetching best seller (week)", err);
      });
  }, []);

  return (
    <div>
      <h2>📆 Best Seller of the Week</h2>
      {book ? (
        <div>
          <p><strong>Title:</strong> {book.title}</p>
          <p><strong>Author:</strong> {book.author}</p>
          <p><strong>Sold:</strong> {book.sold}</p>
        </div>
      ) : (
        <p>No best seller found for this week.</p>
      )}
    </div>
  );
};

export default BestSellerWeek;
