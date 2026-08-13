# Movies API

Example REST API for managing a movie database. Built with Go & SQLite.

## Features

* CRUD operations
* Search and filtering
* Pagination
* SQLite persistence
* Input validation

## Project Structure

See [docs/project-structure.md](docs/project-structure.md) for a detailed explanation of the project layout.

## Getting Started

### Prerequisites

* Go 1.26+

### Installation

Clone the repository:

```bash
git clone https://gitea.kood.tech/jerejuhanimeskanen/movies-api.git
cd movies-api
```

Install the project dependencies:

```bash
go mod download
```

### Running the Application

Start the API server:

```bash
go run ./cmd/api/
```

The server will be available at:

```text
http://localhost:8080
```

The application uses SQLite through the `github.com/mattn/go-sqlite3` driver. The database file is created automatically if it does not already exist.

## Documentation

Additional documentation is available in the `docs/` directory:

* `architecture.md` – Overview of the application architecture and request lifecycle.
* `api.md` – API endpoints, request and response formats, and status codes.
* `project-structure.md` – Explanation of the project layout and package responsibilities.

## Testing

A Postman collection is included in the `postman/` directory for testing all available endpoints.

## License

This project is licensed under the MIT License.
