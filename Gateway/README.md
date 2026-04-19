# Gateway

A Fastify-based API gateway that proxies requests to an incident management service.

## Run

```bash
npm start
```

The server starts on port 3000.

## API Endpoints

### Authentication
- `POST /api/v1/login` - Authenticate user and get JWT token
  - Body: `{ "username": "admin", "password": "password" }`
  - Returns: `{ "token-generated": "jwt_token" }`

### Incidents (requires authentication)
- `GET /api/v1/incidents` - Get all incidents
- `POST /api/v1/incidents` - Create a new incident
- `GET /api/v1/incidents/:id` - Get incident by ID

### Root
- `GET /` - Health check endpoint

## Proxy Configuration

The gateway proxies incident requests to `http://localhost:8080/api/v1/incidents`.

## Docker

Build and run with Docker:

```bash
docker build -t gateway .
docker run -p 3000:3000 gateway
```
