# 📚 mybookstore

A full-stack **Bookstore Application** built with:

- **Frontend**: React + Vite
- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Infrastructure**: Docker + Docker Compose
- **CI/CD**: GitHub Actions + GitHub Container Registry (GHCR)

---

## ✨ Features

- ➕ Add Book
- ✏️ Update Book
- ❌ Delete Book
- 📃 View Books with Pagination, Search & Sorting
- 📊 Bestsellers by Day, Week, Year
- 🐳 Fully Dockerized Setup
- 🔄 GitHub Actions workflows for `dev`, `qa`, `staging`, and `prod` environments
- 🚀 Versioned image publishing to GHCR (using Git tag + commit SHA)

---

## 🗂️ Project Structure

mybookstore/
├── book-service/ # Go backend API
│ ├── main.go
│ ├── routes/
│ ├── db/
│ ├── .env.template
│ └── Dockerfile
├── frontend/ # React frontend (Vite)
│ ├── src/
│ ├── public/
│ ├── .env.template
│ └── Dockerfile
├── scripts/ # Seed script
│ └── seed_books.go
├── environments/ # Environment-specific env files
│ ├── dev.env
│ ├── qa.env
│ ├── staging.env
│ └── prod.env
├── docker-compose.yml
└── .github/workflows/ # CI/CD workflows
├── dev.yml
├── qa.yml
├── staging.yml
└── prod.yml

---

## 🏁 Getting Started (Locally)

### 1. Clone the repository

```bash
git clone https://github.com/bbpandey1/mybookstore.git
cd mybookstore

2. Create .env files
🔐 Backend (book-service/.env)

DB_HOST=localhost         # or 'db' when using Docker Compose
DB_PORT=5432
DB_USER=yourDbUserName
DB_PASSWORD=yourPassword
DB_NAME=bookstore
SSL_MODE=disable
PORT=8080

🌐 Frontend (frontend/.env)

VITE_BACKEND_URL=http://localhost:8080


3. Run with Docker Compose

docker compose --profile dev up --build

Visit:

Frontend: http://localhost:5173

Backend: http://localhost:8080

4.🌱 Seed the Database (optional)
bash
Copy
Edit
go run scripts/seed_books.go
This loads 50 pre-defined books with randomized data.

5.🚀 CI/CD with GitHub Actions + GHCR
Each environment (dev, qa, staging, prod) has its own workflow in .github/workflows/.

✅ Dual Tagging Format
ghcr.io/<OWNER>/mybookstore-backend:dev

ghcr.io/<OWNER>/mybookstore-backend:dev-<commit_sha>

ghcr.io/<OWNER>/mybookstore-frontend:dev

ghcr.io/<OWNER>/mybookstore-frontend:dev-<commit_sha>

6. 🧪 Development Scripts

# Frontend
cd frontend
npm install
npm run dev

# Backend
cd book-service
go run main.go

7.📦 Environment Profiles

# Dev environment
docker compose --profile dev up --build

# QA
docker compose --profile qa up --build

# Staging
docker compose --profile staging up --build

# Production
docker compose --profile prod up --build


📄 License

8. MIT License

MIT License

Copyright (c) 2025 Bharat Bhushan Pandey

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights  
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell      
copies of the Software, and to permit persons to whom the Software is         
furnished to do so, subject to the following conditions:                      

The above copyright notice and this permission notice shall be included in    
all copies or substantial portions of the Software.                           

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR    
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,      
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE   
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER        
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN     
THE SOFTWARE.