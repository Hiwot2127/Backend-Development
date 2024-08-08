# Task Management REST API

## Overview

The Task Management API is a RESTful service built with Go, the Gin framework, and MongoDB. This API allows users to manage tasks and includes authentication and authorization using JSON Web Tokens (JWT). It supports user roles such as admin and regular user, with different access levels to various endpoints.

## Features

- User registration and login with JWT authentication
- Role-based access control (admin and regular user)
- CRUD operations for tasks
- Middleware for protected routes
- Admin-specific functionalities
- Secure password hashing
- JWT token generation and validation
- Role-based access control for admin-specific endpoints

## API Endpoints

### User Management

- **POST /register**
  - Description: Register a new user
  - URL: http://localhost:8080/register
  - Request Body:
    ```json
    {
        "username": "your_username",
        "password": "your_password",
        "role": "admin" // or "user"
    }
    ```

- **POST /login**
  - Description: Login a user
  - URL: http://localhost:8080/login
  - Request Body:
    ```json
    {
        "username": "your_username",
        "password": "your_password"
    }
    ```
  - Response:
    ```json
    {
        "token": "your_jwt_token"
    }
    ```

### Task Management (Protected Routes)

- **GET /api/tasks**
  - Description: Retrieves a list of all tasks
  - URL: http://localhost:8080/api/tasks

- **GET /api/tasks/:id**
  - Description: Retrieves details of a specific task by ID
  - URL: http://localhost:8080/api/tasks/{id}
  - Replace `{id}` with the actual task ID.

- **POST /api/tasks**
  - Description: Creates a new task
  - URL: http://localhost:8080/api/tasks
  - Request Body:
    ```json
    {
        "title": "New Task",
        "description": "Description of the new task",
        "due_date": "2024-08-10T00:00:00Z",
        "status": "Pending"
    }
    ```

- **PUT /api/tasks/:id**
  - Description: Updates an existing task by ID
  - URL: http://localhost:8080/api/tasks/{id}
  - Replace `{id}` with the actual task ID.
  - Request Body:
    ```json
    {
        "title": "Updated Task",
        "description": "Updated description",
        "due_date": "2024-08-15T00:00:00Z",
        "status": "In Progress"
    }
    ```

- **DELETE /api/tasks/:id**
  - Description: Deletes a task by ID
  - URL: http://localhost:8080/api/tasks/{id}
  - Replace `{id}` with the actual task ID.

### Admin-Specific Endpoints (Protected Routes)

- **GET /api/admin/users**
  - Description: Get all users (Admin only)
  - URL: http://localhost:8080/api/admin/users

## Postman Documentation

For detailed API documentation and to test the API endpoints using Postman, refer to the Postman documentation link below:

[Postman API Documentation](https://documenter.getpostman.com/view/37384694/2sA3s1nXD2)

## Setup

### MongoDB Setup

1. **Install MongoDB**: If you haven't already, install MongoDB on your local machine or use a cloud service like MongoDB Atlas.
2. **Start MongoDB**: Ensure MongoDB is running.

### Go Project Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/Hiwot2127/backend-development.git
    cd backend-development
    ```

2. **Install Dependencies**:
    ```sh
    go mod tidy
    ```

3. **Run the Application**:
    ```sh
    go run main.go
    ```

### Configuration

The application connects to MongoDB using the connection URI. Update the connection URI in `main.go` as needed:
```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.Background(), clientOptions)

