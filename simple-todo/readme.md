# Simple Project-Tasks-Todo API

A RESTful API for managing todos with authentication, built with Go, Gin framework, and GORM ORM.

## Project Overview

This is a task management application that allows users to:
- Create and manage user accounts (Sign up, Login)
- Manage projects
- Manage tasks within projects
- Handle JWT-based authentication with refresh tokens

## Technology Stack

- **Language**: Go
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL (via GORM ORM)
- **Authentication**: JWT (JSON Web Tokens)
- **Module Structure**: Modular architecture with clean separation of concerns

## Project Architecture

```
simple-todo/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   │   └── config.go
│   ├── database/               # Database connection & migrations
│   │   ├── database.go
│   │   └── migrate.go
│   ├── models/                 # Data models
│   │   ├── auth_model.go       # User & RefreshTokenTTL
│   │   ├── project_model.go
│   │   └── task_model.go
│   ├── modules/                # Business logic (handler → service → repository)
│   │   ├── auth/
│   │   │   ├── auth_handler.go      # HTTP handlers
│   │   │   ├── auth_service.go      # Business logic
│   │   │   ├── auth_repository.go   # Database operations
│   │   │   ├── auth_routes.go       # Route definitions
│   │   │   └── auth_dto.go          # Request/Response DTOs
│   │   ├── projects/           # (Same structure)
│   │   ├── tasks/              # (Same structure)
│   │   └── middleware/
│   │       └── auth_middleware.go   # JWT validation middleware
│   ├── server/
│   │   └── server.go           # Gin server setup & route registration
│   └── utils/
│       ├── jwt.go              # JWT token generation
│       ├── password.go         # Password hashing
│       └── token.go            # Token utilities
└── go.mod
```

## Getting Started

### Prerequisites
- Go 1.20+
- PostgreSQL
- Environment variables configured



## Module Structure

Each module (auth, projects, tasks) follows this pattern:

```
module/
├── {module}_handler.go     # HTTP handlers, request binding, response formatting
├── {module}_service.go     # Business logic, validation, orchestration
├── {module}_repository.go  # GORM queries, database operations
├── {module}_routes.go      # Route definitions with middleware
├── {module}_dto.go         # Request/Response DTOs
```

### Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Configure environment**
   ```
   DB_URL=postgresql://user:password@localhost:5432/simple_todo
   PORT=8080
   JWT_SECRET=your-secret-key
   ```

3. **Start server**
   ```bash
   go run cmd/main.go
   ```


