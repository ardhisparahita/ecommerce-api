# E-commerce API

[![Go CI](https://github.com/ardhisparahita/ecommerce-api/actions/workflows/ci.yml/badge.svg)](https://github.com/ardhisparahita/ecommerce-api/actions/workflows/ci.yml)

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00AB6C?style=for-the-badge)
![MySQL](https://img.shields.io/badge/MySQL-8-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=for-the-badge&logo=docker&logoColor=white)

RESTful API for an e-commerce application built with Golang, Fiber, GORM, and MySQL following the Clean Architecture approach.

> This README is a professional template tailored for your project. Replace repository links, screenshots, and versions where necessary.

## Features

### Authentication

- User Registration
- User Login
- JWT Authentication
- User Profile
- Update Profile
- Change Password
- Role-Based Authorization (Admin & Customer)

### Product Management

- Category Management
- Product Management
- Product Image Upload

### Shopping

- Address Management
- Shopping Cart
- Checkout
- Order Management
- Payment Status Management

### Developer Tools

- Swagger API Documentation
- Database Migration (golang-migrate)
- Docker Compose (MySQL)
- GitHub Actions (CI)

## Tech Stack

- Golang
- Fiber
- GORM
- MySQL
- JWT
- Swagger
- Docker Compose
- golang-migrate
- GitHub Actions

## Architecture Overview

```text
Client
   │
   ▼
Fiber Router
   │
Middleware
   │
Handlers
   │
Services
   │
Repositories
   │
MySQL
```

## Project Structure

```text
.
├── .github/
├── cmd/
├── docs/
├── internal/
├── migrations/
├── postman/
├── pkg/
├── uploads/
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Prerequisites

- Go 1.25+
- MySQL 8 (or Docker)
- Docker & Docker Compose
- golang-migrate

## Installation

```bash
git clone https://github.com/ardhisparahita/ecommerce-api.git
cd ecommerce-api
go mod tidy
```

## Environment Variables

Create a `.env` file with:

- APP_PORT
- DB_HOST
- DB_PORT
- DB_NAME
- DB_USER
- DB_PASSWORD
- DB_ROOT_PASS
- JWT_SECRET

## Database Migration

Start MySQL:

```bash
docker compose up -d mysql
```

Run migrations:

```bash
migrate -path migrations -database "<DATABASE_URL>" up
```

Rollback:

```bash
migrate -path migrations -database "<DATABASE_URL>" down 1
```

## Running the Application

1. Start MySQL

```bash
docker compose up -d mysql
```

2. Run migrations

```bash
migrate -path migrations -database "<DATABASE_URL>" up
```

3. Start backend

```bash
go run ./cmd
```

Application:

- http://localhost:3000

Swagger:

- http://localhost:3000/swagger/index.html

## Running with Docker

This project uses Docker Compose **only for the MySQL database**.

```bash
docker compose up -d mysql
docker compose ps
docker compose logs -f mysql
docker compose stop mysql
docker compose down
docker compose down -v
```

> Backend runs locally:
>
> ```bash
> go run ./cmd
> ```

## API Documentation

Swagger UI:

`http://localhost:3000/swagger/index.html`

```md
![Swagger UI](docs/images/swagger.png)
```

## Database Schema

```md
![Database ERD](docs/images/erd.svg)
```

## Postman Collection

- `postman/Ecommerce API.postman_collection.json`
- `postman/Ecommerce Local.postman_environment.json`

## API Endpoints

- Authentication
- Users
- Categories
- Products
- Addresses
- Carts
- Checkout
- Orders

(Use the endpoint tables you already prepared.)

## Authentication

Protected endpoints require:

```http
Authorization: Bearer <access_token>
```

## Continuous Integration

GitHub Actions automatically runs:

- Build
- Go Vet
- Unit Test

## License

This project was developed for learning purposes and as a backend development portfolio.

## Author

**Ardhis Parahita**

Backend Developer

- GitHub: https://github.com/ardhisparahita
