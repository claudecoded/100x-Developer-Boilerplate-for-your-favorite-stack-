# 🚀 The 100x Developer Boilerplate (Go + PostgreSQL + Docker)

[![License: Apache 2.0](https://shields.io)](https://opensource.org)
[![Go Version](https://shields.io)](https://golang.org)
[![Docker](https://shields.io)](https://docker.com)
[![Architecture: Clean](https://shields.io)](#)

Stop wasting your first 3 days of a new project wireframing directories, setting up database connections, and managing container environments. 

This is the **ultimate, enterprise-grade production boilerplate** built in **Go (Golang)**. It implements a strict **Clean Architecture**, native dependency injection, resilient database connection pooling, auto-migrations, and smooth graceful shutdown orchestrations. Spin up a rock-solid microservice in exactly one command.

---

## ⚡ Quick Start (Up & Running in 10 Seconds)

Ensure you have [Docker](https://docker.com) installed, then run:

```bash
docker-compose up --build
```

The system will automatically pull PostgreSQL, initialize the isolated volume, run database schema migrations, perform health-check retries, and bind the microservice web server to port `8080`.

---

## 📐 Software Architecture

This boilerplate strictly follows Uncle Bob's **Clean Architecture (Domain-Driven Design concepts)**. The business rules are decoupled from frameworks, databases, and transport layers.

```mermaid
graph TD
    A[🌍 Transport Layer: HTTP / REST REST] --> B[💼 Domain: User Usecase Interface]
    B --> C[🗄️ Infrastructure: User Repository Interface]
    C --> D[(📊 PostgreSQL Database)]

    subgraph Core Domain (Zero External Dependencies)
        E[👤 User Entity & Structural Model]
    end

    B -.-> E
    C -.-> E
    
    style A fill:#1f232a,stroke:#38bdf8,stroke-width:1px,color:#fff
    style B fill:#1f232a,stroke:#34d399,stroke-width:2px,color:#fff
    style C fill:#1f232a,stroke:#34d399,stroke-width:1px,color:#fff
    style D fill:#1f232a,stroke:#fbbf24,stroke-width:1px,color:#fff
    style E fill:#1f232a,stroke:#a855f7,stroke-width:2px,color:#fff
```

### Directory Breakdown
* `main.go`: Global app entry point, orchestrating environment bootstrap, dependency injection, and signal handling loops.
* `internal/domain/`: Enterprise pure business models and storage interface contracts. Zero dependencies.
* `internal/usecase/`: Orchestration flow of application-specific business logic rules.
* `internal/handler/`: HTTP layer mapping parameters, validation routines, and writing responses.
* `internal/repository/`: Low-level data access implementations (PostgreSQL drivers/SQL scripts).
* `pkg/database/`: Shared utilities featuring resilient DB health checking wrapper loops.

---

## 💎 100x Superpowers Included

* **Resilient DB Retry Loop:** If the database container takes a few seconds to spin up, the API will not crash. It loops over a thread-safe exponential backoff health check ping until the pipe connects.
* **Production-Grade Connection Pooling:** Pre-configured maximum idle/open connection lifespans preventing memory leaks and socket starvation issues during massive horizontal scaling.
* **Graceful Shutdown Engine:** Listens to system kill interrupts (`SIGINT`, `SIGTERM`). Safely stops processing incoming traffic, drains existing connection pipelines cleanly, and closes handles without database data corruption.
* **Multi-Stage Docker Optimization:** The image compilation leverages multi-stage layers, outputting a scratch alpine runtime file measuring less than 20MB for blistering deployment intervals.

---

## 🛣️ API Endpoint Documentation

### 1. Register a New User
* **Endpoint:** `POST /api/v1/users`
* **Payload:**
```json
{
  "name": "Alex Developer",
  "email": "alex@100xdev.io"
}
```
* **Response (`214 Created`):**
```json
{
  "id": "c16faec0-3b02-4fc4-bc8a-f584bb7cbdb5",
  "name": "Alex Developer",
  "email": "alex@100xdev.io",
  "created_at": "2026-03-28T16:50:00Z",
  "updated_at": "2026-03-28T16:50:00Z"
}
```

### 2. Fetch All Registered Users
* **Endpoint:** `GET /api/v1/users`
* **Response (`200 OK`):** An array listing registered structural users sorted by creation intervals.

### 3. Get Specific User Profile
* **Endpoint:** `GET /api/v1/users/{id}`

---

## 🛠️ Testing Local Development (Without Containers)

If you wish to run the binary natively on your machine, simply map your environment strings:

```bash
export PORT=8080
export DATABASE_URL="postgres://postgres:postgres123@localhost:5432/hundredx_db?sslmode=disable"

go run main.go
```

---

💡 Why does this code provide "Superpowers"? 

Instant Production Readiness: The developer runs `docker-compose up` and has a fully functional environment with a database, resilient connectivity, and REST routes. 

No Wasted Time on Architecture: Strict isolation within `internal/domain` ensures the database can be swapped (e.g., from PostgreSQL to MongoDB or DynamoDB) simply by changing the repository file, without breaking any business logic. 

Native Security and Robustness: Graceful shutdown management in `main.go` prevents database transaction corruption when the server or container restarts.

## 🤝 Contributing

We welcome structural additions to enhance the architecture scope! Want to add Prometheus metrics middlewares, Redis caching wrappers, or JWT auth interceptors?

1. Fork this Repository
2. Create your Feature Branch (`git checkout -b feature/CoolUpgrade`)
3. Commit changes securely (`git commit -m 'Add JWT Layer Middleware'`)
4. Push upstream (`git push origin feature/CoolUpgrade`)
5. File a clean Pull Request

## 📝 License

Distributed under the Apache License. See `LICENSE` for more structural information.
