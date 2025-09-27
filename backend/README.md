# Easy Ballot Backend API

A Go-based REST API server for the Easy Ballot application.

## Features

- HTTP server with Gorilla Mux router
- CORS support for cross-origin requests
- Request logging middleware
- Health check endpoint
- Standardized API response format
- Environment-based configuration

## Prerequisites

- Go 1.21 or higher
- Git

## Setup

1. **Install dependencies:**

   ```bash
   go mod tidy
   ```

2. **Run the server:**

   ```bash
   go run main.go
   ```

   Or build and run:

   ```bash
   go build -o server main.go
   ./server
   ```

3. **Set environment variables (optional):**
   ```bash
   export PORT=8080  # Default port is 8080
   ```

## API Endpoints

### Health Check

- **GET** `/health`
- Returns server health status and timestamp

### API Info

- **GET** `/api`
- Returns API information and available endpoints

## Project Structure

```
backend/
├── main.go          # Main application file
├── go.mod           # Go module file
├── go.sum           # Go module checksums (generated)
└── README.md        # This file
```

## Development

### Adding New Routes

1. Create a new handler function:

   ```go
   func newHandler(w http.ResponseWriter, r *http.Request) {
       // Handler logic here
   }
   ```

2. Register the route in `setupRoutes()`:
   ```go
   r.HandleFunc("/new-endpoint", newHandler).Methods("GET")
   ```

### Environment Variables

- `PORT`: Server port (default: 8080)

### Dependencies

- `github.com/gorilla/mux`: HTTP router and URL matcher
- `github.com/rs/cors`: CORS handler

## Production Considerations

1. **CORS Configuration**: Update the CORS settings in `main.go` to restrict origins for production
2. **Environment Variables**: Use proper environment variable management
3. **Logging**: Consider using structured logging (e.g., logrus, zap)
4. **Database**: Add database connection and models as needed
5. **Authentication**: Implement JWT or session-based authentication
6. **Rate Limiting**: Add rate limiting middleware
7. **HTTPS**: Configure SSL/TLS certificates

## Testing

Test the API endpoints:

```bash
# Health check
curl http://localhost:8080/health

# API info
curl http://localhost:8080/api
```
