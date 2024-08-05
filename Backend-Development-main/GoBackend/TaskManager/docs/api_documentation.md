# Task Management REST API

## Overview

The Task Management API is a RESTful service designed to facilitate the creation, retrieval, updating, and deletion of tasks. Developed using the Go programming language and the Gin framework, this API provides a robust and efficient solution for managing tasks. It supports basic CRUD operations and stores task data in an in-memory database.

* main.go: Entry point of the application.
* controllers/task_controller.go: Handles incoming HTTP requests and invokes the appropriate service methods.
* models/task.go: Defines the data structures used in the application.
* data/task_service.go: Contains business logic and data manipulation functions.
* router/router.go: Sets up the routes and initializes the Gin router.
* docs/api_documentation.md: Contains API documentation and other related documentation.
* go.mod: Go module file.

## API Endpoints

### GET /tasks
- Description: Retrieves a list of all tasks.
- Method: GET
- URL: http://localhost:8080/tasks

### GET /tasks/
- Description: Retrieves details of a specific task by ID.
- Method: GET
- URL: http://localhost:8080/tasks/{id}
- Replace {id} with the actual task ID.

### POST /tasks
- Description: Creates a new task.
- Method: POST
- URL: http://localhost:8080/tasks
- Request Body:
{
  "title": "New Task",
  "description": "Description of the new task",
  "due_date": "2024-08-10T00:00:00Z",
  "status": "Pending"
}

### PUT /tasks/
- Description: Updates an existing task by ID.
- Method: PUT
- URL: http://localhost:8080/tasks/{id}
- Replace {id} with the actual task ID.
- Request Body:
{
  "title": "Updated Task",
  "description": "Updated description",
  "due_date": "2024-08-15T00:00:00Z",
  "status": "In Progress"
}

### DELETE /tasks/
- Description: Deletes a task by ID.
- Method: DELETE
- URL: http://localhost:8080/tasks/{id}
- Replace {id} with the actual task ID.

## Postman Documentation
For detailed API documentation and to test the API endpoints using Postman, refer to the Postman documentation link below:
# Postman API Documentation
https://documenter.getpostman.com/view/37384694/2sA3rwLE7C

# Testing the API Using Postman
* Open Postman.
* Create a new collection named "Task Management API".
* Add requests to this collection for each endpoint as described in the API documentation.
* Send requests to test the API functionality and observe the responses.

# How to Run
* Clone the repository.
* git clone https://github.com/yourusername/TaskManager.git
* cd TaskManager
* Install dependencies.
* go mod tidy
* Run the application.
* go run main.go
* The server will start on http://localhost:8080.