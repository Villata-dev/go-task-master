# go-task-master

`go-task-master` is a powerful and intuitive task management application designed to streamline your workflow. With support for both a command-line interface (CLI) and a RESTful API, it offers the flexibility to manage your tasks in the way that works best for you.

## Features

- **CLI & API Support**: Manage tasks via the command line or integrate with other applications through the API.
- **Easy to Use**: A simple and intuitive interface makes task management a breeze.
- **Docker Support**: Quickly get up and running with Docker.

## Getting Started

To get started with `go-task-master`, you can either run the application using Docker or build it from the source.

### Prerequisites

- Go (if building from source)
- Docker (if using Docker)

### Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/user/go-task-master.git
   cd go-task-master
   ```

2. **Build the application**:
   ```sh
   go build -o task-master .
   ```

## Docker

You can also run the application using Docker.

1. **Build the Docker image**:
   ```sh
   docker build -t go-task-master .
   ```

2. **Run the Docker container**:
   ```sh
   docker run -p 8080:8080 go-task-master
   ```

## CLI Usage

The CLI provides a simple way to manage your tasks.

- **Serve**: Start the Task Master server.
  ```sh
  ./task-master serve
  ```

- **List**: List all tasks.
  ```sh
  ./task-master list
  ```

- **Add**: Add a new task.
  ```sh
  ./task-master add "My new task"
  ```

- **Complete**: Mark a task as complete.
  ```sh
  ./task-master complete 1
  ```

- **Delete**: Delete a task.
  ```sh
  ./task-master delete 1
  ```

## API Documentation

The API provides a set of endpoints to manage tasks programmatically.

- **GET /tasks**: Retrieve all tasks.
  ```sh
  curl http://localhost:8080/tasks
  ```

- **POST /tasks**: Create a new task.
  ```sh
  curl -X POST -H "Content-Type: application/json" -d '{"title":"New Task"}' http://localhost:8080/tasks
  ```

- **PUT /tasks/{id}**: Update a task.
  ```sh
  curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated title"}' http://localhost:8080/tasks/1
  ```

- **DELETE /tasks/{id}**: Delete a task.
  ```sh
  curl -X DELETE http://localhost:8080/tasks/1
  ```
