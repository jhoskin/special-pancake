# Todo Service

A simple and efficient Todo management service implemented in Go with gRPC.

## Project Structure

```
.
├── internal/
│   ├── domain/         # Domain models and interfaces
│   ├── infrastructure/
│   │   ├── config/     # Application configuration
│   │   ├── db/        # Database implementation using BoltDB
│   │   └── server/    # HTTP server and gRPC handlers
│   └── features/
│       ├── listtodos/  # List todos handler
│       ├── createtodo/ # Create todo handler
│       ├── updatetodo/ # Update todo handler
│       └── deletetodo/ # Delete todo handler
├── proto/             # Protocol Buffers definitions
│   ├── gen/          # Generated protobuf code
│   └── todo/         # Todo service proto definitions
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
└── main.go           # Application entry point
```

## Features

- Complete Todo CRUD operations via gRPC
- Persistent storage using BoltDB
- Clean architecture design
- Protocol Buffers for efficient data serialization

## Prerequisites

- Go 1.21 or higher
- Protocol Buffers compiler (protoc)
- Buf CLI tool for protocol buffer management

## Getting Started

### Configuration

The service can be configured using environment variables:

| Variable | Description               | Default       |
| -------- | ------------------------- | ------------- |
| DB_PATH  | Path to the database file | data/todos.db |
| PORT     | Port to listen on         | 8080          |

Example:

```bash
DB_PATH=/var/lib/todo-service/data/todos.db PORT=3000 ./todo-service
```

### Building for Production

For production deployments, you can create optimized builds for different architectures:

```bash
# Build for Linux AMD64 (most common)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o todo-service-linux-amd64

# Build for Linux ARM64 (e.g., AWS Graviton)
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o todo-service-linux-arm64

# Build for macOS Intel
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o todo-service-darwin-amd64

# Build for macOS Apple Silicon
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o todo-service-darwin-arm64

# Build for Windows AMD64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o todo-service-windows-amd64.exe
```

Build flags explained:

- `CGO_ENABLED=0`: Creates a statically linked binary
- `-ldflags="-w -s"`: Reduces binary size by removing debug information
- `GOOS`: Sets the target operating system
- `GOARCH`: Sets the target architecture

### API Reference

The service exposes the following HTTP endpoints:

#### TodoService

##### ListTodos

```bash
curl \
  -X POST \
  -H "Content-Type: application/json" \
  http://localhost:8080/todo.v1.TodoService/ListTodos
```

##### CreateTodo

```bash
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Buy groceries",
    "description": "Milk, bread, eggs",
    "completed": false
  }' \
  http://localhost:8080/todo.v1.TodoService/CreateTodo
```

##### UpdateTodo

```bash
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, bread, eggs, cheese",
    "completed": true
  }' \
  http://localhost:8080/todo.v1.TodoService/UpdateTodo
```

##### DeleteTodo

```bash
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1
  }' \
  http://localhost:8080/todo.v1.TodoService/DeleteTodo
```
