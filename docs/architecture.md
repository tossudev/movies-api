# Architecture

## Overview

The Movies API follows a layered architecture to separate concerns and improve maintainability. Each layer has a single responsibility and communicates only with the layer directly below it.

```text
                HTTP Request
                      │
                      ▼
                Route Registration
                      │
                      ▼
                  Middleware
                      │
                      ▼
                 HTTP Handlers
                      │
                      ▼
                 Service Layer
                      │
                      ▼
               Repository Layer
                      │
                      ▼
            SQLite Database (database/sql)
```

This design keeps HTTP concerns, business logic, and data access independent, making the application easier to understand, test, and extend.

---

# Architectural Layers

## Route Layer

**Location**

```text
internal/routes/
```

### Responsibilities

- Register all API endpoints
- Map HTTP methods and URLs to handlers
- Configure middleware
- Configure the application's routing

### Example

```text
GET    /api/movies
POST   /api/movies
PATCH  /api/movies/{id}
DELETE /api/movies/{id}
```

The routing layer contains no business logic and simply connects middleware and handlers.

---

## Middleware Layer

**Location**

```text
internal/middleware/
```

### Responsibilities

Middleware intercepts every incoming HTTP request before it reaches the handlers. It is responsible for cross-cutting concerns that apply to multiple endpoints without duplicating code.

Current middleware includes:

- **Request Logging** – Logs each incoming request, including HTTP method, path, status code, and execution time.
- **Panic Recovery** – Recovers from unexpected panics and returns a standardized `500 Internal Server Error` response instead of terminating the application.

### Request Flow

```text
HTTP Request
      │
      ▼
Logging Middleware
      │
      ▼
Recovery Middleware
      │
      ▼
Handler
```

Middleware should remain lightweight and should not contain business logic, validation, or database operations.

---

## Handler Layer

**Location**

```text
internal/handlers/
```

### Responsibilities

Handlers are responsible for processing HTTP requests and responses.

Responsibilities include:

- Parsing JSON request bodies
- Reading path parameters
- Reading query parameters
- Calling the appropriate service
- Returning JSON responses
- Setting HTTP status codes

### Example Flow

```text
HTTP Request
      │
      ▼
MovieHandler.CreateMovie()
      │
      ▼
MovieService.CreateMovie()
```

Handlers should remain lightweight and should not contain SQL queries or business rules.

---

## Service Layer

**Location**

```text
internal/service/
```

### Responsibilities

The service layer contains all business logic.

Examples include:

- Creating and updating entities
- Managing movie-genre relationships
- Managing movie-actor relationships
- Filtering movies
- Searching movies
- Pagination
- Force deletion logic
- Validation beyond simple field checks

The service layer coordinates one or more repositories to complete an operation.

### Example

```text
Validate Request
        │
        ▼
Create Movie
        │
        ▼
Assign Genres
        │
        ▼
Assign Actors
        │
        ▼
Return Result
```

---

## Repository Layer

**Location**

```text
internal/repository/
```

### Responsibilities

The repository layer interacts directly with the SQLite database.

Responsibilities include:

- CRUD operations
- Raw SQL queries
- Join queries
- Transactions
- Returning repository-specific errors

Repositories should not contain HTTP logic or business rules.

### Example

```sql
SELECT *
FROM movies
WHERE release_year = ?;
```

The application uses Go's standard `database/sql` package together with the `go-sqlite3` driver.

---

## Database Layer

**Location**

```text
internal/db/
```

### Responsibilities

- Open the SQLite connection
- Initialize the database
- Execute SQL migrations
- Manage the database lifecycle

Database schema is maintained using migration files located in:

```text
migrations/
```

---

# Request Lifecycle

The following sequence illustrates how a typical request flows through the application.

```text
Client
   │
   ▼
HTTP Request
   │
   ▼
Routes
   │
   ▼
Logging Middleware
   │
   ▼
Recovery Middleware
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
SQLite Database
   │
   ▼
Repository
   │
   ▼
Service
   │
   ▼
Handler
   │
   ▼
HTTP Response
```

Each layer performs only its designated responsibility before passing control to the next layer.

---

# Project Components

## Models

**Location**

```text
internal/models/
```

Models represent the application's core domain entities.

- Movie
- Genre
- Actor

These structures closely mirror the database schema.

---

## Data Transfer Objects (DTOs)

**Location**

```text
internal/dto/
```

DTOs define the request and response payloads exposed through the REST API.

Separating DTOs from models:

- Prevents exposing database models directly
- Allows API contracts to evolve independently
- Simplifies validation for create and update requests

Examples include:

```text
CreateMovieRequest
UpdateMovieRequest
MovieResponse
```

---

## Validation

**Location**

```text
internal/validator/
```

Validation ensures incoming data is correct before business logic is executed.

Typical validation includes:

- Required fields
- Positive duration values
- Valid release year
- ISO 8601 birth date format
- Valid relationship identifiers

---

## HTTP Utilities

**Location**

```text
internal/http/
```

Shared HTTP utilities provide:

- Standard JSON responses
- Error formatting
- Response helpers

Keeping these utilities centralized ensures all endpoints produce consistent API responses.

---

# Error Handling

Errors propagate upward through the application layers.

```text
Repository Error
        │
        ▼
Service
        │
Business Decision
        │
        ▼
Handler
        │
        ▼
HTTP Response
```

Example response:

```json
{
    "error": "movie not found"
}
```

Common status codes include:

| Status | Description |
|--------|-------------|
| 200 | OK |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

---

# Database Relationships

The application models two many-to-many relationships.

```text
           Movie
          /     \
         /       \
movie_genres   movie_actors
       |             |
       ▼             ▼
    Genre         Actor
```

Join tables:

- `movie_genres`
- `movie_actors`

These relationships allow:

- Movies to belong to multiple genres
- Genres to contain multiple movies
- Movies to feature multiple actors
- Actors to appear in multiple movies

---

# Design Principles

## Separation of Concerns

Each architectural layer has a single, clearly defined responsibility.

| Layer | Responsibility |
|--------|----------------|
| Routes | Register endpoints and middleware |
| Middleware | Cross-cutting HTTP concerns |
| Handlers | Process HTTP requests and responses |
| Services | Business logic |
| Repositories | Database access |
| Database | Data persistence |

---

## Single Responsibility Principle

Each package focuses on one area of responsibility.

For example:

- Routes never contain business logic.
- Middleware never performs database operations.
- Handlers never execute SQL.
- Services never parse HTTP requests.
- Repositories never generate HTTP responses.

---

## Maintainability

Separating responsibilities makes the project easier to extend.

Adding a new feature typically requires:

1. Registering a new route.
2. Creating or updating a handler.
3. Implementing business logic in the service.
4. Adding repository methods if database access is required.

Existing layers remain largely unaffected.

---

## Scalability

Although designed as a small educational project, this architecture scales well as the application grows.

Potential future enhancements include:

- Authentication and authorization middleware
- CORS middleware
- Rate limiting
- Dependency injection
- Structured logging
- Unit testing with mocked repositories
- OpenAPI (Swagger) documentation
- Containerization with Docker

The layered architecture allows these enhancements to be introduced with minimal impact on the existing codebase.
