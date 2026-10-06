# GoTask API

A production-style RESTful Task Management API built with Go, Gin, PostgreSQL, GORM, JWT authentication, bcrypt, and Docker.

## Features

- User registration and login
- Secure password hashing using bcrypt
- JWT-based authentication
- Protected task APIs
- Create tasks
- Get all tasks
- Get a single task by ID
- Update tasks
- Delete tasks
- User-specific task access
- PostgreSQL database
- GORM ORM
- Environment-based configuration
- Dockerized application
- Docker Compose
- RESTful API architecture
- Postman API testing

## Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- JWT
- bcrypt
- Docker
- Docker Compose
- Postman
- Git
- GitHub

## Project Structure

GoTask-API/
│
├── config/
│   └── database.go
│
├── handlers/
│   ├── auth.go
│   └── task.go
│
├── middleware/
│   └── auth.go
│
├── models/
│   ├── user.go
│   └── task.go
│
├── routes/
│   └── routes.go
│
├── .env
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/register | Register a new user |
| POST | /api/login | Login and receive JWT |

### Tasks

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/tasks | Create a task |
| GET | /api/tasks | Get all tasks |
| GET | /api/tasks/:id | Get one task |
| PUT | /api/tasks/:id | Update a task |
| DELETE | /api/tasks/:id | Delete a task |

## Authentication

Protected task endpoints require a JWT token.

Add the token to the request:

Authorization: Bearer <your-jwt-token>

In Postman:

1. Open the request.
2. Go to Authorization.
3. Select Bearer Token.
4. Paste the JWT token.
5. Send the request.

## Register User

### Request

POST /api/register

### Body

```json
{
  "name": "Isha",
  "email": "isha@gotask.com",
  "password": "password123"
}