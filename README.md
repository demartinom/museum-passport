# 🖼️ Museum Passport

A high-performance web application that aggregates cultural heritage data from global institutions. Built with a **Go** microservice backend and a **Next.js 15** frontend, it features AI-driven historical context and a robust containerized architecture.

---

## Installation

The entire stack (Frontend, Backend, and Networking) is containerized for instant deployment and environment parity.

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/museum-passport.git
cd museum-passport
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory:

```env
# API Keys
OPENAI_KEY=your_openai_key_here
HARVARD_KEY=your_harvard_key_here

# Networking
NEXT_PUBLIC_API_URL=http://localhost:3001
DOCKER_ENV=true
```

### 3. Launch with Docker

```bash
docker compose up --build
```

| Service      | URL                              |
| ------------ | -------------------------------- |
| Frontend     | http://localhost:3000            |
| API Explorer | http://localhost:3001/api/search |

---

## 🏗️ Architecture & Tech Stack

### Backend

- **Language:** Go 1.25+ (Standard library + Chi Router)
- **Concurrency:** Uses Goroutines to fetch data from multiple museum APIs (The Met & Harvard) in parallel, reducing latency by ~60%.
- **Intelligence:** OpenAI GPT-4o-mini integration for generating historical `AISummaries`.
- **Infrastructure:** Dockerized with multi-stage builds to keep production images lightweight and secure.

### Frontend

- **Framework:** Next.js 15 (App Router)
- **Rendering:** Server-Side Rendering (SSR) for artwork pages to optimize SEO and initial load speed.
- **Styling:** Tailwind CSS + shadcn/ui for a minimalist, museum-grade aesthetic.

---

## 🔧 Dev & Deployment Logic

### Local Development (Manual)

If you prefer running the binaries directly on your host machine:

```bash
# Backend — defaults to port 3001
cd server && go run main.go

# Frontend — defaults to port 3000
cd frontend && npm run dev
```

### Production Networking

- **Fly.io (Backend):** The Go binary dynamically detects the `$PORT` assigned by the cloud environment, defaulting to `8080` in production while maintaining `3001` for local Docker development.
- **Vercel (Frontend):** Optimized for edge deployment. CORS is pre-configured to allow secure communication between your Vercel domain and the Fly.io API.

---

## Engineering Challenges

### 1. Data Normalization

**Challenge:** Museum APIs (Met vs. Harvard) return vastly different JSON structures for `Artist` and `Medium` fields.

**Solution:** Implemented a unified `Art` struct in Go. Each museum client includes a custom "Mapper" function that cleans and translates raw API responses into a consistent internal schema.

### 2. Environment Synchronization

**Challenge:** Differences in case-sensitivity between macOS (development) and Linux/Docker (production) caused build-time "Module Not Found" errors.

**Solution:** Enforced a strict casing convention for all React components and leveraged Docker to validate build integrity before deploying to production.

---

## 🛰️ API Endpoints

| Method | Endpoint                  | Description                                     |
| ------ | ------------------------- | ----------------------------------------------- |
| `GET`  | `/api/search?general={q}` | Global search across all integrated museums     |
| `GET`  | `/api/artwork/{id}`       | Fetches high-res metadata for a specific object |
| `GET`  | `/api/summary?id={id}`    | Triggers AI generation of historical context    |

---

## 📜 License & Acknowledgments

- **Data Providers:** [The Metropolitan Museum of Art Open Access API](https://metmuseum.github.io/) and [Harvard Art Museums API](https://github.com/harvardartmuseums/api-docs).
- **License:** MIT
