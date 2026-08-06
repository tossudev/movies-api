# API Documentation

## Base URL

```
http://localhost:8080/api
```

## Authentication

This API does not require authentication.

## Content Type

All requests and responses use JSON.

```
Content-Type: application/json
Accept: application/json
```

---

# Movies

## Create Movie

**POST** `/movies`

### Request

```json
{
  "title": "Inception",
  "releaseYear": 2010,
  "duration": 148,
  "genreIds": [1, 3],
  "actorIds": [2, 5, 8]
}
```

### Response

**201 Created**

```json
{
  "id": 21,
  "title": "Inception",
  "releaseYear": 2010,
  "duration": 148,
  "genres": [...],
  "actors": [...]
}
```

---

## Get All Movies

**GET** `/movies`

### Query Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| page | int | Page number |
| size | int | Items per page |

### Example

```
GET /api/movies?page=0&size=10
```

---

## Get Movie By ID

**GET** `/movies/{id}`

### Example

```
GET /api/movies/15
```

---

## Update Movie

**PATCH** `/movies/{id}`

### Request

```json
{
  "duration": 152
}
```

---

## Delete Movie

**DELETE** `/movies/{id}`

### Force Delete

```
DELETE /movies/15?force=true
```

---

## Search Movies

**GET** `/movies/search?title=matrix`

---

## Filter Movies

### By Genre

```
GET /movies?genre=3
```

### By Release Year

```
GET /movies?year=1999
```

### By Actor

```
GET /movies?actor=5
```

---

## Get Actors in a Movie

```
GET /movies/{id}/actors
```

---

# Genres

## Endpoints

| Method | Endpoint |
|---------|----------|
| POST | `/genres` |
| GET | `/genres` |
| GET | `/genres/{id}` |
| PATCH | `/genres/{id}` |
| DELETE | `/genres/{id}` |

### Get Movies in Genre

```
GET /genres/{id}/movies
```

---

# Actors

## Endpoints

| Method | Endpoint |
|---------|----------|
| POST | `/actors` |
| GET | `/actors` |
| GET | `/actors/{id}` |
| PATCH | `/actors/{id}` |
| DELETE | `/actors/{id}` |

### Filter by Name

```
GET /actors?name=tom
```

### Get Movies by Actor

```
GET /movies?actor=5
```

---

# Error Responses

## Validation Error

**400 Bad Request**

```json
{
  "error": "title is required"
}
```

---

## Resource Not Found

**404 Not Found**

```json
{
  "error": "movie not found"
}
```

---

## Relationship Constraint

**400 Bad Request**

```json
{
  "error": "Cannot delete genre 'Action' because it has associated movies."
}
```

---

# Pagination

Example response:

```json
{
  "page": 0,
  "size": 10,
  "total": 20,
  "items": [
    ...
  ]
}
```

---

# Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |
