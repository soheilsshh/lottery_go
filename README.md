# Lottery/Draw System API

A production-ready lottery/draw system API built with Go, Gin framework, and Clean Architecture pattern.

## Architecture

This project follows **Clean Architecture** (also known as Hexagonal Architecture) principles:

- **Entities**: Core business objects (User, Draw)
- **Use Cases**: Business logic and application rules
- **Repositories**: Data access layer (abstracted for easy database switching)
- **Controllers**: HTTP request handlers
- **Routes**: API endpoint definitions

## Features

- User registration and authentication (JWT ready)
- Entering lottery draws
- Performing lottery draws (random winner selection)
- Listing past winners
- Database abstraction layer (SQLite for dev, PostgreSQL for prod)

## Project Structure

```
lottery-system/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── config/                 # Configuration management
│   └── config.go
├── internal/               # Internal application code
│   ├── entities/          # Domain entities
│   ├── usecases/          # Business logic
│   ├── repositories/      # Data access layer
│   ├── controllers/       # HTTP handlers
│   └── routes/            # API routes
├── pkg/                    # Reusable packages
│   └── database/          # Database connection and setup
└── migrations/             # Database migration files
```

## Prerequisites

- Go 1.24.5 or higher
- SQLite (for development)
- PostgreSQL (for production, optional)

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd lottery
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables (create `.env` file):
```env
DB_TYPE=sqlite
DB_PATH=./lottery.db
SERVER_PORT=8080
JWT_SECRET=your-secret-key-here
```

4. Run database migrations (when implemented):
```bash
# TODO: Add migration command
```

5. Start the server:
```bash
go run main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints (To be implemented)

### User Endpoints
- `POST /api/v1/users/register` - Register a new user
- `POST /api/v1/users/login` - User login (JWT)
- `GET /api/v1/users/me` - Get current user info

### Draw Endpoints
- `POST /api/v1/draws/enter` - Enter a lottery draw
- `POST /api/v1/draws/perform` - Perform lottery (select winner)
- `GET /api/v1/draws/winners` - List past winners
- `GET /api/v1/draws/history` - Get draw history

## Database

- **Development**: SQLite (file-based, no setup required)
- **Production**: PostgreSQL (easily switchable via configuration)

The repository pattern allows switching databases without changing business logic.

## Authentication

JWT authentication is prepared but not yet implemented. The middleware structure is in place for future implementation.

## Development

This is a skeleton project. Implementation details are marked with `TODO` comments throughout the codebase.

## License

[Add your license here]

