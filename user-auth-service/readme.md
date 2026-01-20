# Explains File Structure

user-auth-service/
│
├── cmd/
│   └── server/
│       └── main.go - to load configs and start the server
│
├── internal/
│   ├── config/
│   │   ├── config.go - load env variables and prepare configs
│   │
│   ├── database/
│   │   └── postgres.go - connect with db and automatic migrations
│   │
│   ├── modules/
│   │   └── auth/
│   │       ├── auth_handler.go - handle http work like status codes and json errors
│   │       ├── auth_service.go - connect all code - utils, repository, and model
│   │       ├── auth_repository.go - actual communication with db 
│   │       ├── auth_routes.go - to define routes of apis 
│   │       ├── auth_dto.go - to define body of requests for apis
│   │       └── auth_model.go - to define user model
│   │
│   ├── middleware/
│   │   └── auth_middleware.go - auth middleware for userId
│   │
│   ├── utils/
│   │   ├── jwt.go - to validate and generate jwt tokens
│   │   ├── password.go - to compare and genreate hashed password
│   │   └── token.go  - to create a hashedtoken
│   │
│   │
│   └── server/
│       └── server.go - create server and run function with all configs and db connections
│
│
├── .env
├── go.mod
└── README.md
