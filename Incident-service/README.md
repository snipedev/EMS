# Incident Service

A RESTful API for managing incidents — create, retrieve, update status, and close incidents.

## Tech Stack

- **Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL

## Base URL

```
/api/v1
```

## Endpoints

### Get All Incidents

```
GET /api/v1/incidents
```

Returns a list of all incidents.

**Response `200 OK`**
```json
[
  {
    "id": "1",
    "title": "Server down",
    "status": "open",
    "created_at": "2024-01-01T10:00:00Z"
  }
]
```

---

### Get Incident by ID

```
GET /api/v1/incidents/:id
```

Returns a single incident by its ID.

**Path Parameters**

| Parameter | Type   | Description     |
|-----------|--------|-----------------|
| `id`      | string | Incident ID     |

**Response `200 OK`**
```json
{
  "id": "1",
  "title": "Server down",
  "status": "open",
  "created_at": "2024-01-01T10:00:00Z"
}
```

**Response `404 Not Found`**
```json
{
  "error": "incident not found"
}
```

---

### Create Incident

```
POST /api/v1/incidents
```

Creates a new incident.

**Request Body**
```json
{
  "title": "Server down",
  "description": "Production server is unresponsive"
}
```

**Response `201 Created`**
```json
{
  "id": "2",
  "title": "Server down",
  "status": "open",
  "created_at": "2024-01-01T10:00:00Z"
}
```

---

### Update Incident Status

```
PATCH /api/v1/incidents/:id
```

Updates the status of an existing incident.

**Path Parameters**

| Parameter | Type   | Description     |
|-----------|--------|-----------------|
| `id`      | string | Incident ID     |

**Request Body**
```json
{
  "status": "in_progress"
}
```

**Response `200 OK`**
```json
{
  "id": "1",
  "title": "Server down",
  "status": "in_progress",
  "updated_at": "2024-01-01T11:00:00Z"
}
```

---

### Close Incident

```
DELETE /api/v1/incidents/:id
```

Closes an incident by its ID.

**Path Parameters**

| Parameter | Type   | Description     |
|-----------|--------|-----------------|
| `id`      | string | Incident ID     |

**Response `200 OK`**
```json
{
  "message": "incident closed successfully"
}
```

**Response `404 Not Found`**
```json
{
  "error": "incident not found"
}
```

---

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- A `.env` file (see below)

### Environment Variables

Create a `.env` file in the root directory:

```env
DATABASE_URL=postgres://user:password@localhost:5432/incidents_db
PORT=8080
```

### Installation

```bash
# Clone the repository
git clone https://github.com/your-org/incident-service.git
cd incident-service

# Install dependencies
go mod tidy

# Run the service
go run main.go
```

The service will start on `http://localhost:8080`.

---

## Project Structure

```
incident-service/
├── main.go
├── .env
├── go.mod
├── go.sum
├── handlers/
│   └── incident.go
├── models/
│   └── incident.go
└── db/
    └── db.go
```
