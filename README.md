# Todo Service

A simple and efficient Todo management service implemented in Go with gRPC.

## Project Structure

```
.
├── internal/
│   ├── common/
│   │   └── db/          # Database implementation using BoltDB
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

The service exposes the following gRPC endpoints:

#### TodoService

- `ListTodos`: Retrieve all todos
- `CreateTodo`: Create a new todo
- `UpdateTodo`: Update an existing todo
- `DeleteTodo`: Delete a todo by ID

### Data Model

Todo structure:

```go
type Todo struct {
    ID          uint
    Title       string
    Description string
    Completed   bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Database

The service uses BoltDB for data persistence. The database file will be automatically created in the configured location when the service starts.

## Testing

Run the tests with:

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
