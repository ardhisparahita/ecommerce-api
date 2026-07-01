# E-commerce API

[![Go CI](https://github.com/ardhisparahita/ecommerce-api/actions/workflows/ci.yml/badge.svg)](https://github.com/ardhisparahita/ecommerce-api/actions/workflows/ci.yml)

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00AB6C?style=for-the-badge)
![MySQL](https://img.shields.io/badge/MySQL-8-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=for-the-badge&logo=docker&logoColor=white)

RESTful API for an e-commerce application built with Golang, Fiber, GORM, and MySQL following the Clean Architecture approach.

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

### Option 1 — Local Development

#### 1. Start MySQL

```bash
docker compose up -d mysql
```

#### 2. Run Database Migrations

```bash
migrate -path migrations -database "<DATABASE_URL>" up
```

#### 3. Start the Backend

```bash
go run ./cmd
```

The API server will be available at:

```
http://localhost:3000
```

Swagger UI:

```
http://localhost:3000/swagger/index.html
```

## Running with Docker

You can also run the backend together with the MySQL database using Docker Compose.

### Start MySQL and Backend

```bash
docker compose up --build
```

Or run in detached mode.

```bash
docker compose up -d --build
```

### Start Only the Backend

If the MySQL container is already running, you can start only the backend.

```bash
docker compose up backend
```

Or run in detached mode.

```bash
docker compose up -d backend
```

### View Running Containers

```bash
docker compose ps
```

### View Backend Logs

```bash
docker compose logs -f backend
```

### View Database Logs

```bash
docker compose logs -f mysql
```

### Stop Backend

```bash
docker compose stop backend
```

### Stop MySQL

```bash
docker compose stop mysql
```

### Remove Containers

```bash
docker compose down
```

### Remove Containers and Database Volume

```bash
docker compose down -v
```

> **Note**
>
> Before starting the backend container for the first time, make sure the database migrations have been applied using **golang-migrate**.

## API Documentation

Swagger UI:

http://localhost:3000/swagger/index.html

![Swagger UI](docs/images/swagger.png)

## Database Schema

![Database ERD](docs/images/erd.svg)

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
