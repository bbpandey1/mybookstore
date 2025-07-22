import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080",
});

export const getBooks = (page, limit, search = "", sort = "", order = "") =>
  API.get("/books", {
    params: { page, limit, search, sort, order },
  });

export const addBook = (data) => API.post("/books", data);
export const deleteBook = (id) => API.delete(`/books/${id}`);
export const updateBook = (id, data) => API.put(`/books/${id}`, data);
export const getBookById = (id) => API.get(`/books/${id}`);
export const getBestSellerToday = () => API.get("/bestsellers/day");
export const getBestSellerWeek = () => API.get("/bestsellers/week");
export const getBestSellerYear = () => API.get("/bestsellers/year");