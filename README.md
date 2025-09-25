# Mini Books Library

> A small full-stack Next.js + TypeScript frontend and a Golang backend for managing a tiny book library (CRUD) — plus an URL cleanup / redirection service for assignment test purpose.

---

## Table of contents
- [Overview](#overview)
- [Project Structure](#project-layout)
  - [Frontend](#frontend)
  - [Backend](#backend)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Make CLI](#make-cli)
- [Dockerization](#Dockerization)
- [Look! It Runs!!](#look--it-runs--)
- [API — endpoints & examples](#api---endpoints--examples)
  - [Books endpoints](#books-endpoints)
  - [URL cleanup & redirection endpoint](#url-cleanup--redirection-endpoint)
- [Postman collection (import-ready snippet)](#postman-collection-import-ready-snippet)
- [cURL examples](#curl-examples)
- [Data schema & ERD (books)](#data-schema--erd-books)
- [Swagger / API specification](#swagger--api-specification)
- [Running tests (BE & FE)](#running-tests-be--fe)
- [Design decisions & trade-offs](#design-decisions--trade-offs)
- [Future improvements](#future-improvements)
- [Contributing](#contributing)
- [Assumptions & TODO / what I couldn't auto-detect](#assumptions--todo--what-i-couldnt-auto-detect)

---
## Overview

This repository is a monorepo containing:

- `backend/` — Golang REST API implementing:
  - Books CRUD (`GET /books`, `POST /books`, `GET /books/:id`, `PUT /books/:id`, `DELETE /books/:id`)
  - URL cleanup & redirection POST endpoint (accepts `{ "url": "...", "operation": "redirection|canonical|all" }` and returns `{ "processed_url": "..." }`) per the assignment.
- `frontend/` — Next.js + TypeScript app that provides a dashboard and forms (add/edit/view/delete books).
- Top-level: `Makefile`, `docker-compose.yml` for local orchestration.


> The project assignment describes both the Books CRUD API and the URL-cleanup/redirection behavior and output format — I used that spec to craft the README examples. 

---

## Prerequisites

- Node.js (v18+ recommended) & npm (or yarn/pnpm) — for frontend.
- Go (1.25+ recommended) — for backend.
- make — for the Makefile convenience targets (optional).
- Docker & docker-compose — recommended for DB and full-stack local dev.
- (Optional) `swag` CLI for Swagger docs: `go install github.com/swaggo/swag/cmd/swag@latest`

---

## Installation

1. Install dependencies
- Frontend
```
npm install
```
- Backend
```
go mod tidy
```
2. Create environment file `.env.local` for FE and set key/value below:
```
NEXT_PUBLIC_API_URL=http://localhost:9000
```

---



## Make CLI

| Command                               | Description                                                         |
|---------------------------------------|---------------------------------------------------------------------|
| `make help`                           | Show available commands on Makefile.            |
| `make migrate-up`                     | Apply database migrations (e.g., add new tables/columns).            |
| `make migrate-down`                   | Revert database migrations (e.g., remove tables/columns).            |
| `make docker-build service=<name>`    | Build the specified service (e.g., `backend`, `postgres`).          |
| `make docker-up service=<name>`       | Start the specified service (e.g., `frontend`, `backend`, `postgres`) |
| `make docker-down`                    | Stop and remove all running containers and networks.                 |
| `make swag-up`                    | Generate Swagger API documentation.                 |
| `make seed`                    | Seed books data into the database.                 |

![Makefile Commands](docs/makefile/makefile-help.png)

---

## Dockerization
> Step by step installation using docker
1. Run Database Postgres until Up
```
make docker-up service=postgres

or

docker-compose up postgres
```

2. Build Backend 
```
make docker-build service=backend

or

docker-compose build --no-cache backend

```
3. Run Backend 
```
make docker-up service=backend

or

docker-compose up backend

```
4. Seed Database 
```
make seed

or

cd backend
go run seeds/main.go

```
5. Run Frontend (Next.js/Typescript)
```
make docker-up service=frontend

or 

docker-compose up frontend
```
![Docker Image](docs/docker/running-docker-image.png)

---

## Look! It Runs!!

1. Open http://localhost:3000 in your browser
> After all services on docker image running and also seeding done. It's time to check FE from browser.

![Frontend Running](docs/frontend/book-dashboard.png)

2. Open http://localhost:9000/swagger/index.html in your browser
> We can play with data using swagger

![Swagger API](docs/swagger/swagger-1.png)
![Swagger Entities](docs/swagger/swagger-2.png)
---

