# Backend Go Microservices - E-Commerce Monorepo

This is a monorepo for a microservices-based e-commerce backend application built with Go.

## Project Structure

```
backend-go/
├── go.work                 # Go workspace configuration
├── shared/                 # Shared packages across all services
│   ├── go.mod
│   └── pkg/
├── services/
│   ├── product-service/   # Product catalog service
│   ├── order-service/     # Order management service
│   └── user-service/      # User authentication service
└── protobuf/              # Protocol buffer definitions
```

## Setup Instructions

### Prerequisites

- Go 1.25.0 or higher
- PostgreSQL (for product-service)
- Docker (optional, for containerized deployment)

### Getting Started

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd backend-go
   ```

2. **Initialize Go Workspace**

   The project uses Go workspaces (introduced in Go 1.18) to manage multiple modules. The `go.work` file is already configured.

   ```bash
   # Verify workspace setup
   go work sync
   ```

3. **Install Dependencies**

   ```bash
   # Install shared module dependencies
   cd shared && go mod tidy

   # Install product-service dependencies
   cd ../services/product-service && go mod tidy

   # Install other services...
   ```

4. **Environment Configuration**

   Create a `.env` file in each service directory:

   ```bash
   cd services/product-service
   cp .env.example .env  # If example exists
   # Edit .env with your configuration
   ```

5. **Run Services**

   ```bash
   # Run product-service
   cd services/product-service
   go run cmd/main.go
   ```

## Using Shared Packages

The `shared` module contains common utilities used across all services. To use shared packages in a service:

```go
import "github.com/kaunghtethein/backend-go/shared/pkg"

// Use shared validator
func RegisterRoutes(e *echo.Echo) {
    e.POST("/endpoint", pkg.BindAndValidate(handler.Function))
}

// Use shared error definitions
if err != nil {
    if errors.Is(err, pkg.DuplicateEntry) {
        // Handle duplicate entry
    }
}

// Use shared string utilities
capitalized := pkg.CapitalizeFirstLetter("hello")
```

## Module Structure

Each service is a separate Go module with its own `go.mod` file. The shared module is referenced using:

```go
// In service go.mod
require github.com/kaunghtethein/backend-go/shared v0.0.0

replace github.com/kaunghtethein/backend-go/shared => ../../shared
```

## Development Workflow

1. **Adding a New Service**

   ```bash
   mkdir -p services/new-service
   cd services/new-service
   go mod init new-service
   ```

   Then add it to `go.work`:

   ```bash
   go work use ./services/new-service
   ```

2. **Adding Shared Utilities**

   - Add code to `shared/pkg/`
   - Run `go mod tidy` in the shared directory
   - The changes will be automatically available to all services via the workspace

3. **Building a Service**
   ```bash
   cd services/product-service
   go build -o bin/product-service ./cmd/main.go
   ```

## Testing

```bash
# Test a specific service
cd services/product-service
go test ./...

# Test shared package
cd shared
go test ./...
```

## Docker Support

Each service has its own Dockerfile for containerized deployment:

```bash
cd services/product-service
docker build -t product-service .
docker run -p 8080:8080 product-service
```

## Contributing

1. Create a feature branch
2. Make your changes
3. Run tests
4. Submit a pull request
