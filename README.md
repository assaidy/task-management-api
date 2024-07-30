# Task Management API

This repository contains a RESTful API for managing tasks. It provides endpoints for creating, retrieving, updating, completing, and deleting tasks.

### API Endpoints

The following API endpoints are available:

- **`GET /v1/tasks`**: Retrieve all tasks
  - **Example URL**: `http://localhost:8080/v1/tasks`
  - **Expected Responses**:
    - **Success**: 200 OK, JSON array of tasks
    - **Failure**: 500 Internal Server Error

- **`POST /v1/tasks`**: Create a new task
  - **Example URL**: `http://localhost:8080/v1/tasks?content=New%20Task`
  - **Expected Responses**:
    - **Success**: 201 Created, JSON of the created task
    - **Failure**: 400 Bad Request if content is missing, 500 Internal Server Error

- **`GET /v1/task`**: Retrieve a specific task by ID
  - **Example URL**: `http://localhost:8080/v1/tasks/1`
  - **Expected Responses**:
    - **Success**: 200 OK, JSON of the task
    - **Failure**: 400 Bad Request if ID is missing or invalid, 404 Not Found if task is not found

- **`PUT /v1/toggleComplete`**: Toggle the completion status of a task
  - **Example URL**: `http://localhost:8080/v1/toggleComplete/1`
  - **Expected Responses**:
    - **Success**: 200 OK, JSON of the updated task
    - **Failure**: 400 Bad Request if ID is missing or invalid, 404 Not Found if task is not found

- **`DELETE /v1/deleteTask`**: Delete a task by ID
  - **Example URL**: `http://localhost:8080/v1/deleteTask/1`
  - **Expected Responses**:
    - **Success**: 204 No Content
    - **Failure**: 400 Bad Request if ID is missing or invalid, 404 Not Found if task is not found

- **`PUT /v1/updateTask`**: Update a task's content by ID
  - **Example URL**: `http://localhost:8080/v1/updateTask/1&content=Updated%20Task`
  - **Expected Responses**:
    - **Success**: 200 OK, JSON of the updated task
    - **Failure**: 400 Bad Request if ID or content is missing or invalid, 404 Not Found if task is not found

- **`DELETE /v1/deleteAllTasks`**: Delete all tasks
  - **Example URL**: `http://localhost:8080/v1/deleteAllTasks`
  - **Expected Responses**: 204 No Content
### Installation

1. **Clone the Repository**

```sh
   git clone <repository_url>
   cd <repository_directory>
```
2. **run api server**
```sh
    make run
```

### Todo
- [ ] add logging
- [X] add some tests
- [ ] add real database

