## Project Name
An Event-Driven E-Commerce Backend built with Golang, Kafka, PostgreSQL, and Redis

---

## Table of Contents

- [Project Name](#project-name)
- [Table of Contents](#table-of-contents)
- [About Project](#about-project)
- [Requirement](#requirement)
- [Folder Structure](#folder-structure)
- [Getting Started](#getting-started)
- [Author](#author)

---

## About Project
Project ini meng-implementasikan microservice-based architecture untuk e-commerce platform menggunakan event-driven communication.
Digunakan untuk belajar dan demonstrates scalable service design dengan Kafka dan PostgreSQL serta menggunakan native Golang.

**Key Features:**
- Event-driven architecture (Kafka)
- Clean Architecture with domain-driven design
- PostgreSQL persistent using `pgx`
- Redis caching layer
- JWT-based authentication
- REST + gRPC communication between services

## Requirement
- Golang (net/http, gRPC) : Internal and extenal API
- Apache Kafka : Event-driven communication
- PostgreSQL : Persistent data storage
- Redis : Cache layer
- JWT : Secure Authentication mechanism

## Folder Structure
```
/learn-microservice-kafka/
│
├── docker-compose.yml        # Kafka, Zookeeper, Postgres, Redis setup
│
├── shared/
│   ├── kafka/
│   │   ├── producer.go       # shared Kafka producer helper
│   │   ├── consumer.go       # shared Kafka consumer helper
│   │   └── config.go
│   ├── db/
│   │   └── postgres.go       # shared DB connection (pgx)
│   ├── redis/
│   │   └── client.go         # optional caching helpers
│   ├── logger/
│   │   └── logger.go         # zap or log standard wrapper
│   └── models/
│       └── events.go         # standardized event models
│
├── user-service/
│   ├── main.go
│   ├── internal/
│   │   ├── http/
│   │   │   └── handlers.go
│   │   ├── repository/
│   │   └── service/
│   └── go.mod
│
├── product-service/
│   ├── main.go
│   ├── internal/
│   │   ├── http/
│   │   ├── repository/
│   │   └── service/
│
├── order-service/
│   ├── main.go
│   ├── internal/
│   │   ├── http/
│   │   │   └── handler.go
│   │   ├── repository/
│   │   │   ├── order_repository.go
│   │   │   └── outbox_repository.go
│   │   ├── service/
│   │   │   ├── order_service.go
│   │   │   ├── outbox_publisher.go
│   │   │   └── consumer.go
│   │   └── model/
│   │       ├── order.go
│   │       └── outbox.go
│
├── payment-service/
│   ├── main.go
│   └── internal/
│       ├── consumer/
│       └── repository/
│
└── notification-service/
    ├── main.go
    └── internal/
        └── consumer/
```

## Getting Started
Follow these steps to set up and run the project locally

**📦 Installation**
Clone repository:
```bash
git clone https://github.com/kharisma-wardhana/learn-microservice-kafka.git
```
Install dependencies:
```bash
go mod tidy
```

**🐳 Run with Docker**
```bash
docker-compose up -d
```

**▶️ Run Locally**
```bash
go run main.go
```

## Author
Kharisma Wardhana