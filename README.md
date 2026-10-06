# 🚀 GoTask API

A production-style **RESTful Task Management API** built with **Go, Gin, PostgreSQL, GORM, JWT authentication, bcrypt, and Docker**.

GoTask API provides secure user authentication and a complete task management system with user-specific authorization. The application is fully containerized using Docker Compose for consistent local development and deployment.

---

## ✨ Features

- 🔐 User registration and login
- 🔑 JWT-based authentication
- 🔒 Protected API routes
- 🔐 Password hashing with bcrypt
- 📝 Create tasks
- 📋 Get all tasks for the authenticated user
- 🔎 Get a single task by ID
- ✏️ Update tasks
- 🗑️ Delete tasks
- 👤 User-specific task authorization
- 🐘 PostgreSQL database
- 🗃️ GORM ORM
- 🌱 Environment-based configuration
- 🐳 Dockerized Go application
- 🐳 Docker Compose for API + PostgreSQL
- 🧪 Postman API testing
- 🏗️ RESTful API architecture

---

## 🛠️ Tech Stack

| Technology | Purpose |
|---|---|
| **Go** | Backend programming language |
| **Gin** | HTTP web framework |
| **GORM** | ORM and database operations |
| **PostgreSQL** | Relational database |
| **JWT** | Authentication and authorization |
| **bcrypt** | Password hashing |
| **Docker** | Application containerization |
| **Docker Compose** | Multi-container orchestration |
| **Postman** | API testing |
| **Git & GitHub** | Version control |

---

## 🏗️ Architecture

```text
                    ┌─────────────────┐
                    │   Client /      │
                    │    Postman      │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   Gin Router    │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ JWT Middleware  │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │    Handlers     │
                    │ Auth + Tasks    │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │      GORM       │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   PostgreSQL    │
                    └─────────────────┘
```

---

## 📁 Project Structure

```text
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
```

> ⚠️ `.env` contains sensitive configuration and should never be committed to GitHub.

---

# 🔐 Authentication

GoTask API uses **JWT-based authentication**.

### Authentication Flow

```text
User
 │
 │ Register
 ▼
POST /api/register
 │
 │ bcrypt password hashing
 ▼
PostgreSQL
```

```text
User
 │
 │ Login
 ▼
POST /api/login
 │
 │ Verify password
 ▼
Generate JWT
 │
 ▼
Return JWT
```

```text
Client
 │
 │ Authorization: Bearer <JWT>
 ▼
JWT Middleware
 │
 │ Valid token
 ▼
Protected Task Routes
 │
 ▼
PostgreSQL
```

---

# 📡 API Endpoints

## Authentication

| Method | Endpoint | Description | Authentication |
|---|---|---|---|
| `POST` | `/api/register` | Register a new user | ❌ |
| `POST` | `/api/login` | Login and receive JWT | ❌ |

## Task Management

| Method | Endpoint | Description | Authentication |
|---|---|---|---|
| `POST` | `/api/tasks` | Create a task | ✅ |
| `GET` | `/api/tasks` | Get all user tasks | ✅ |
| `GET` | `/api/tasks/:id` | Get a specific task | ✅ |
| `PUT` | `/api/tasks/:id` | Update a task | ✅ |
| `DELETE` | `/api/tasks/:id` | Delete a task | ✅ |

---

# 🧪 API Examples

## 1. Register User

### Request

```http
POST http://localhost:8080/api/register
Content-Type: application/json
```

### Body

```json
{
  "name": "Isha",
  "email": "isha@gotask.com",
  "password": "password123"
}
```

### Response

```json
{
  "message": "User registered successfully",
  "user": {
    "id": 1,
    "name": "Isha",
    "email": "isha@gotask.com"
  }
}
```

---

## 2. Login

### Request

```http
POST http://localhost:8080/api/login
Content-Type: application/json
```

### Body

```json
{
  "email": "isha@gotask.com",
  "password": "password123"
}
```

### Response

```json
{
  "message": "Login successful",
  "token": "eyJ..."
}
```

---

## 3. Create Task

### Request

```http
POST http://localhost:8080/api/tasks
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

### Body

```json
{
  "title": "Learn Go Backend",
  "description": "Complete GoTask API project"
}
```

### Response

```json
{
  "message": "Task created successfully",
  "task": {
    "ID": 1,
    "title": "Learn Go Backend",
    "description": "Complete GoTask API project",
    "completed": false,
    "user_id": 1
  }
}
```

---

## 4. Get All Tasks

### Request

```http
GET http://localhost:8080/api/tasks
Authorization: Bearer <JWT_TOKEN>
```

### Response

```json
{
  "tasks": [
    {
      "ID": 1,
      "title": "Learn Go Backend",
      "description": "Complete GoTask API project",
      "completed": false,
      "user_id": 1
    }
  ]
}
```

---

## 5. Get Single Task

### Request

```http
GET http://localhost:8080/api/tasks/1
Authorization: Bearer <JWT_TOKEN>
```

### Response

```json
{
  "task": {
    "ID": 1,
    "title": "Learn Go Backend",
    "description": "Complete GoTask API project",
    "completed": false,
    "user_id": 1
  }
}
```

---

## 6. Update Task

### Request

```http
PUT http://localhost:8080/api/tasks/1
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

### Body

```json
{
  "title": "Learn Go and Docker",
  "description": "Complete Go backend and Docker setup",
  "completed": true
}
```

### Response

```json
{
  "message": "Task updated successfully",
  "task": {
    "ID": 1,
    "title": "Learn Go and Docker",
    "description": "Complete Go backend and Docker setup",
    "completed": true,
    "user_id": 1
  }
}
```

---

## 7. Delete Task

### Request

```http
DELETE http://localhost:8080/api/tasks/1
Authorization: Bearer <JWT_TOKEN>
```

### Response

```json
{
  "message": "Task deleted successfully"
}
```

---

# 🔑 Using JWT Authentication in Postman

For protected endpoints:

1. Open the request in Postman.
2. Select **Authorization**.
3. Select **Bearer Token**.
4. Paste the JWT received from `/api/login`.
5. Click **Send**.

Alternatively, add the header manually:

```text
Authorization: Bearer <your-jwt-token>
```

---

# 🐳 Docker Setup

The application uses Docker Compose to run:

```text
┌──────────────────────────────┐
│       Docker Compose         │
│                              │
│  ┌──────────────┐            │
│  │  GoTask API  │            │
│  │    :8080     │            │
│  └──────┬───────┘            │
│         │                    │
│         ▼                    │
│  ┌──────────────┐            │
│  │  PostgreSQL  │            │
│  │    :5432     │            │
│  └──────────────┘            │
│                              │
└──────────────────────────────┘
```

## Build the Application

```bash
docker compose build
```

## Start Containers

```bash
docker compose up -d
```

## Check Container Status

```bash
docker compose ps
```

Expected services:

```text
gotask-api
gotask-postgres
```

The PostgreSQL container should show:

```text
healthy
```

## View API Logs

```bash
docker compose logs api
```

## Follow API Logs

```bash
docker compose logs -f api
```

## View Database Logs

```bash
docker compose logs db
```

## Stop Containers

```bash
docker compose down
```

---

# 🌱 Environment Configuration

Create a `.env` file in the project root.

```env
DB_HOST=localhost
DB_USER=gotask_user
DB_PASSWORD=gotask_password
DB_NAME=gotask_db
DB_PORT=5432
JWT_SECRET=your_secure_jwt_secret
```

For Docker Compose, the API container uses the PostgreSQL service name:

```text
DB_HOST=db
```

The `.env` file should **never be committed to GitHub**.

Use `.env.example` to document the required environment variables.

---

# 🗄️ Database

The project uses **PostgreSQL** with **GORM**.

### Models

#### User

```text
User
├── ID
├── Name
├── Email
└── Password
```

#### Task

```text
Task
├── ID
├── Title
├── Description
├── Completed
└── UserID
```

Each task is associated with a specific user.

This ensures users can only access their own tasks.

---

# 🔒 Security

The API implements several security practices:

- Passwords are hashed using **bcrypt**
- Authentication uses **JWT**
- Task endpoints are protected by middleware
- Tasks are filtered by authenticated `user_id`
- Database credentials are stored using environment variables
- JWT secret is stored using environment variables
- `.env` is excluded from Git

---

# 🧪 Testing

The API was tested using **Postman** for:

- User registration
- User login
- JWT authentication
- Task creation
- Get all tasks
- Get single task
- Task update
- Task deletion

---

# 🚀 Running Locally

## Prerequisites

Make sure you have:

- Go installed
- Docker Desktop installed
- Git installed
- Postman installed

## Clone Repository

```bash
git clone https://github.com/Isha4002/GoTask-API.git
cd GoTask-API
```

## Install Dependencies

```bash
go mod download
```

## Configure Environment

Create `.env` based on `.env.example`.

## Run with Docker

```bash
docker compose up -d
```

## Verify

Open:

```text
http://localhost:8080/
```

Expected response:

```json
{
  "message": "GoTask API is running!"
}
```

---

# 📈 Future Improvements

Potential future improvements:

- [ ] Refresh token authentication
- [ ] Role-based authorization
- [ ] Task priorities
- [ ] Task deadlines
- [ ] Pagination
- [ ] Search and filtering
- [ ] Unit testing
- [ ] Integration testing
- [ ] Swagger/OpenAPI documentation
- [ ] Structured logging
- [ ] Rate limiting
- [ ] CI/CD pipeline
- [ ] Cloud deployment
- [ ] Graceful server shutdown

---

# 👩‍💻 Author

**Isha Pal**

B.Tech Information Technology  
JSS Academy of Technical Education, Noida

GitHub:  
https://github.com/Isha4002

---

## ⭐ If you found this project useful

Feel free to star the repository!
