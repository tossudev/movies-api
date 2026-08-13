# Project Structure

## Overview

The project is organized using a layered architecture that separates routing, HTTP handling, business logic, and database access into distinct packages. This approach improves readability, maintainability, and scalability while following common Go project conventions.

```text
movies-api/
├── cmd/
│   └── api/
│       └── main.go                  # Application entry point
│
├── internal/
│   ├── config/
│   │   └── config.go                # Application configuration
│   │
│   ├── db/
│   │   └── db.go                    # SQLite connection and migration runner
│   │
│   ├── models/
│   │   ├── movie.go                 # Movie entity
│   │   ├── genre.go                 # Genre entity
│   │   └── actor.go                 # Actor entity
│   │
│   ├── dto/
│   │   ├── error.go                 # Error response DTO
│   │   ├── movie.go                 # Request and response DTOs
│   │   ├── genre.go
│   │   └── actor.go
│   │
│   ├── repository/
│   │   ├── movie.go                 # Database operations
│   │   ├── genre.go
│   │   ├── actor.go
│   │   └── errors.go                # Repository-specific errors
│   │
│   ├── service/
│   │   ├── movie.go                 # Business logic
│   │   ├── genre.go
│   │   └── actor.go
│   │
│   ├── handlers/
│   │   ├── movie.go                 # HTTP handlers
│   │   ├── genre.go
│   │   └── actor.go
│   │
│   ├── middleware/
│   │   ├── logging.go               # Request logging middleware
│   │   └── recover.go               # Panic recovery middleware
│   │
│   ├── response/
│   │   └── response.go              # JSON and error response helpers
│   │
│   ├── routes/
│   │   └── routes.go                # Route registration
│   │
│   └── validator/
│       └── validator.go             # Input validation
│
├── migrations/
│   └── 0001_init.sql                # Database schema
│
├── testdata/
│   └── seed.sql                     # Sample data
│
├── postman/
│   └── movies-api.postman_collection.json
│
├── docs/
│   ├── architecture.md
│   ├── api.md
│   └── project-structure.md
│
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

---

# Directory Overview

## cmd/

Contains the application's executable entry point.

| File | Purpose |
|------|---------|
| `api/main.go` | Initializes configuration, database connection, middleware, routes, and starts the HTTP server. |

---

## internal/

Contains the application's implementation. Packages inside `internal` cannot be imported by external projects, helping encapsulate the application's business logic.

### config/

Application configuration.

Responsibilities:

- Load configuration values
- Configure database path
- Configure server port
- Store application-wide settings

---

### db/

Database initialization and connection management.

Responsibilities:

- Open the SQLite connection
- Execute migrations
- Manage the database lifecycle

---

### models/

Domain models representing the core entities.

| File | Entity |
|------|--------|
| `movie.go` | Movie |
| `genre.go` | Genre |
| `actor.go` | Actor |

These structures closely mirror the database schema and are used internally throughout the application.

---

### dto/

Data Transfer Objects (DTOs).

DTOs define the JSON payloads exchanged with clients and keep API contracts separate from internal database models.

Examples include:

- CreateMovieRequest
- UpdateMovieRequest
- MovieResponse
- CreateActorRequest

Using DTOs allows the API to evolve without affecting the underlying data models.

---

### repository/

The data access layer.

Responsibilities include:

- CRUD operations
- SQL queries
- Join queries
- Transactions
- Repository-specific errors

This layer communicates directly with SQLite using Go's `database/sql` package.

---

### service/

The business logic layer.

Responsibilities include:

- Coordinating repositories
- Validation beyond simple field checks
- Search functionality
- Pagination
- Managing many-to-many relationships
- Force deletion logic
- Business rules

Services remain independent of HTTP and SQL implementation details.

---

### handlers/

HTTP request handlers.

Responsibilities include:

- Parsing requests
- Reading URL and query parameters
- Calling services
- Returning JSON responses
- Setting HTTP status codes

Handlers remain lightweight and contain no business logic.

---

### middleware/

Shared HTTP middleware executed before requests reach handlers.

Current middleware:

| Middleware | Purpose |
|------------|---------|
| `logging.go` | Logs incoming requests, response status, and execution time. |
| `recover.go` | Recovers from panics and returns a standardized `500 Internal Server Error` response. |

Middleware provides reusable functionality that applies to every endpoint.

---

### response/

Shared API response utilities for writing consistent responses.

Responsibilities include:

- JSON response helpers
- Error formatting
- Standardized API responses

This avoids duplicated response code across handlers.

---

### routes/

Registers all API endpoints and connects them to middleware and handlers.

Example:

```text
GET    /api/movies
POST   /api/movies
PATCH  /api/movies/{id}
DELETE /api/movies/{id}
```

---

### validator/

Contains reusable validation logic for incoming requests.

Typical validation includes:

- Required fields
- Positive numeric values
- Valid release years
- ISO 8601 birth date format
- Relationship validation

---

# Supporting Directories

## migrations/

Contains SQL migration files used to initialize the database schema.

Current migration:

```text
0001_init.sql
```

---

## testdata/

Contains SQL seed data used for testing and demonstrations.

The seed data includes:

- Genres
- Movies
- Actors
- Relationship data

---

## postman/

Contains the Postman collection used to test all API endpoints.

---

## docs/

Contains project documentation.

| Document | Purpose |
|----------|---------|
| `architecture.md` | Application architecture and request lifecycle |
| `api.md` | REST API reference |
| `project-structure.md` | Project organization and package responsibilities |

---

# Dependency Flow

Each layer depends only on the layer directly beneath it.

```text
Routes
   │
   ▼
Middleware
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ▼
Repositories
   │
   ▼
SQLite Database
```

This dependency flow keeps responsibilities clearly separated and minimizes coupling between packages.

---

# Design Principles

## Separation of Concerns

Each package has a single responsibility:

- **Routes** register endpoints.
- **Middleware** handles cross-cutting HTTP concerns.
- **Handlers** process HTTP requests and responses.
- **Services** implement business logic.
- **Repositories** perform database operations.
- **Database** manages persistence.

---

## Maintainability

The layered architecture makes the project easier to extend and maintain.

Adding a new feature typically involves:

1. Creating or updating a handler.
2. Implementing service logic.
3. Adding repository methods.
4. Registering the endpoint.

Existing components require minimal modification.

---

## Scalability

Although designed for a small REST API, the project structure supports future enhancements such as:

- Authentication and authorization
- Additional middleware
- Dependency injection
- Unit testing with mocked repositories
- OpenAPI (Swagger) documentation
- Docker deployment
- Configuration via environment variables

The modular structure allows these features to be introduced without significant architectural changes.
