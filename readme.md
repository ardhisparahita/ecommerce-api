Isi README lengkap

# Ecommerce API

[![Go CI](https://github.com/ardhisparahitaa/ecommerce-api/actions/workflows/ci.yml/badge.svg)](https://github.com/ardhisparahitaa/ecommerce-api/actions/workflows/ci.yml)

![Go](<[https://img.shields.io/badge/Go-1.25-00ADD8?logo=go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)>)

![Fiber](<[https://img.shields.io/badge/Fiber-v2-00AB6C](https://img.shields.io/badge/Fiber-v2-00AB6C)>)

![MySQL](<[https://img.shields.io/badge/MySQL-8-4479A1?logo=mysql](https://img.shields.io/badge/MySQL-8-4479A1?logo=mysql)>)

![Docker](<[https://img.shields.io/badge/Docker-Enabled-2496ED?logo=docker](https://img.shields.io/badge/Docker-Enabled-2496ED?logo=docker)>)

RESTful API for an e-commerce application built with Golang, Fiber, GORM, and MySQL following the Clean Architecture approach.

### Features

Authentication

- User Registration

- User Login

- JWT Authentication

- User Profile

- Update Profile

- Change Password

- Role-Based Authorization (Admin & Customer)

Product Management

- Category Management

- Product Management

- Product Image Upload

Shopping

- Address Management

- Shopping Cart

- Checkout

- Order Management

- Payment Status Management

Developer Tools

- Swagger Documentation

- Database Migration

- Docker & Docker Compose

- GitHub Actions (CI)

### Tech Stack

- Golang

- Fiber

- GORM

- MySQL

- JWT

- Swagger

- Docker

- Docker Compose

- golang-migrate

- GitHub Actions

### Architecture Overview

Client

Fiber Router

Middleware

Handler

Service

Repository

MySQL

### Project Structure

. ├── .github │ └── workflows ├── cmd ├── docs │ └── images │ ├── swagger.png │ └── erd.png ├── internal │ ├── domain │ ├── dto │ ├── handler │ ├── middleware │ ├── repository │ ├── routes │ └── service ├── migrations ├── postman │ ├── Ecommerce API.postman_collection.json │ └── Ecommerce Local.postman_environment.json ├── pkg │ ├── config │ ├── database │ └── utils ├── uploads ├── Dockerfile ├── docker-compose.yml ├── go.mod ├── go.sum └── README.md

### Prerequisites

- Go 1.25+

- MySQL 8

- Docker & Docker Compose (optional)

- golang-migrate

### Installation

Clone repository

Install dependencies

### Environment Variables

Create a `.env` file in the project root and configure the following variables.

| Variable     | Description         |
| ------------ | ------------------- |
| APP_PORT     | Application Port    |
| DB_HOST      | Database Host       |
| DB_PORT      | Database Port       |
| DB_NAME      | Database Name       |
| DB_USER      | Database Username   |
| DB_PASSWORD  | Database Password   |
| DB_ROOT_PASS | MySQL Root Password |
| JWT_SECRET   | JWT Secret          |

Note: The `.env` file is intentionally excluded from this repository because it contains sensitive configuration.

### Database Migration

Run migration

Rollback migration

### Running the Application

The application will be available at `http://localhost:3000`

### Running with Docker

Build and start the application

Run in detached mode

Stop containers

Remove containers and volumes

### API Documentation

Swagger UI is available at `http://localhost:3000/swagger/index.html`

![Swagger UI](docs/images/swagger.png)

### Database Schema (ERD)

![Database ERD](docs/images/erd.svg)

### Postman Collection

Import the collection from:

Collection

JSON

Ecommerce API.postman_collection.json

All API endpoints for testing

Ecommerce Local.postman_environment.json

Local environment variables

### API Endpoints

### Authentication

Public

| Method | Endpoint              | Description         |
| ------ | --------------------- | ------------------- |
| POST   | /api/v1/auth/register | Register a new user |
| POST   | /api/v1/auth/login    | Login               |

### User

JWT

| Method | Endpoint                      | Description         |
| ------ | ----------------------------- | ------------------- |
| GET    | /api/v1/users/profile         | Get user profile    |
| PUT    | /api/v1/users/profile         | Update user profile |
| PATCH  | /api/v1/users/change-password | Change password     |

### Categories

Admin Write

| Method | Endpoint                | Description             |
| ------ | ----------------------- | ----------------------- |
| GET    | /api/v1/categories      | Get all categories      |
| GET    | /api/v1/categories/{id} | Get category by ID      |
| POST   | /api/v1/categories      | Create category (Admin) |
| PUT    | /api/v1/categories/{id} | Update category (Admin) |

### Products

Admin Manage

| Method | Endpoint                    | Description                  |
| ------ | --------------------------- | ---------------------------- |
| GET    | /api/v1/products            | Get all products             |
| GET    | /api/v1/products/{id}       | Get product details          |
| POST   | /api/v1/products            | Create product (Admin)       |
| PUT    | /api/v1/products/{id}       | Update product (Admin)       |
| DELETE | /api/v1/products/{id}       | Delete product (Admin)       |
| POST   | /api/v1/products/{id}/image | Upload product image (Admin) |

### Addresses

JWT

| Method | Endpoint               | Description       |
| ------ | ---------------------- | ----------------- |
| GET    | /api/v1/addresses      | Get all addresses |
| GET    | /api/v1/addresses/{id} | Get address by ID |
| POST   | /api/v1/addresses      | Create address    |
| PUT    | /api/v1/addresses/{id} | Update address    |
| DELETE | /api/v1/addresses/{id} | Delete address    |

### Shopping Cart

JWT

| Method | Endpoint           | Description         |
| ------ | ------------------ | ------------------- |
| GET    | /api/v1/carts      | Get shopping cart   |
| POST   | /api/v1/carts      | Add product to cart |
| PUT    | /api/v1/carts/{id} | Update cart item    |
| DELETE | /api/v1/carts/{id} | Remove cart item    |

### Checkout

JWT

| Method | Endpoint          | Description   |
| ------ | ----------------- | ------------- |
| POST   | /api/v1/checkouts | Checkout cart |

### Orders

Admin Status

| Method | Endpoint                     | Description                   |
| ------ | ---------------------------- | ----------------------------- |
| GET    | /api/v1/orders               | Get all orders                |
| GET    | /api/v1/orders/{id}          | Get order details             |
| PATCH  | /api/v1/orders/{id}/pay      | Mark order as paid (Admin)    |
| PATCH  | /api/v1/orders/{id}/fail     | Mark payment as failed        |
| PATCH  | /api/v1/orders/{id}/cancel   | Cancel order                  |
| PATCH  | /api/v1/orders/{id}/ship     | Mark order as shipped (Admin) |
| PATCH  | /api/v1/orders/{id}/complete | Mark order as completed       |

### Authentication

Protected endpoints require a JWT Bearer Token.

### Admin Only

- Create Category

- Update Category

- Create Product

- Update Product

- Delete Product

- Upload Product Image

- Mark Order as Paid

- Mark Order as Shipped

### Continuous Integration

GitHub Actions automatically performs the following checks on every push and pull request:

- Build Project

- Go Vet

- Unit Test

### License

This project was developed for learning purposes and as a backend development portfolio.

### Author

Ardhis Parahita

Backend Developer

GitHub: https://github.com/ardhisparahitaa
