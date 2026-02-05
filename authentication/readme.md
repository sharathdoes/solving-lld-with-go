# User Auth Service (Gin + GORM)

## File Structure (one-line each)

cmd/main.go – App entrypoint, loads config and starts server

internal/config – Loads env config (DB URL, JWT secret, TTLs)

internal/server – Gin engine setup, middleware, route registration

internal/database – Postgres connection + AutoMigrate

internal/modules/auth

* auth_handler.go – HTTP handlers (signup, login, refresh)
* auth_service.go – Business logic (hashing, tokens, rotation)
* auth_repository.go – DB queries via GORM
* auth_models.go – User and RefreshToken DB models
* auth_routes.go – /auth route wiring
* auth_dto.go – Request structs + validation

internal/middleware – JWT auth middleware (access token validation)

internal/utils

* jwt.go – Generate & validate JWTs
* password.go – Hash & compare passwords

---

## User Flow (8 steps)

1. User signs up → password is hashed → user stored in DB
2. User logs in with email + password
3. Password is verified against hash
4. Server issues short-lived access token (JWT with user_id)
5. Server issues long-lived refresh token (random string)
6. Refresh token hash is stored in DB (not the raw token)
7. Client uses access token for protected APIs
8. When access token expires, client calls /auth/refresh to get new tokens
