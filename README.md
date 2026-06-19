# Keto Granola

Monorepo for Keto Granola e-commerce platform.

## App architecture
![high-level system architecture](apparchitecture.png)

Cloudflare acts as a full proxy caching static assets and SSR HTML, forwarding cache misses to the backend server. The server renders all routes via `html/template`, with React islands hydrating interactive components 
client-side. 

## Structure

- `/frontend` — Vite + React islands. See [frontend/README.md](./frontend/README.md)
- `/backend` — Go + Echo. See [backend/README.md](./backend/README.md)