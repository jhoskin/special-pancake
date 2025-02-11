# Todo Service

A simple and efficient Todo management service implemented in Go with gRPC.

## Project Structure

```
.
├── internal/
│   ├── infrastructure/
│   │   ├── db/          # Database implementation using BoltDB
│   │   └── server/      # HTTP server and gRPC handlers
│   └── features/
│       ├── listtodos/   # List todos handler
│       ├── createtodo/  # Create todo handler
│       ├── updatetodo/  # Update todo handler
│       └── deletetodo/  # Delete todo handler
├── models/              # Data models
├── proto/              # Protocol Buffers definitions
└── main.go            # Application entry point
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

### Installing Go on macOS

1. Install Homebrew if you haven't already:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Install Go using Homebrew:

```bash
brew install go
```

3. Verify the installation:

```bash
go version
```

4. Set up your Go workspace (add to ~/.zshrc or ~/.bash_profile):

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

5. Reload your shell configuration:

```bash
source ~/.zshrc  # or source ~/.bash_profile
```

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

1. Clone the repository:

```bash
git clone <repository-url>
cd todo-service
```

2. Install dependencies:

```bash
go mod download
```

3. Run the service:

```bash
go run main.go
```

## Development

### Building the Project

Build the binary with:

```bash
go build -o todo-service
```

### Generating Protocol Buffers

To regenerate the Protocol Buffers code after making changes to the `.proto` files:

```bash
buf generate
```

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

### Data Model

Todo structure:

```

```
