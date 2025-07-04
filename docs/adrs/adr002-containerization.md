# ADR 002: Containerization with Docker
Status: Accepted

### Context:
The application stack includes Golang (backend), SQLite (database), HTMX and Tailwind (frontend). Consistent, reproducible deployments are required across environments.

### Decision:
* All application components will be containerized using Docker.
* Multi-stage Docker builds will be used to optimize image size and security.
* SQLite database will be persisted using Docker volumes.
* Environment variables will be used for configuration, with .env.dev and .env.prod files for each environment.

### Consequences:
* Ensures consistent builds and deployments.
* Simplifies local development and cloud deployment.
* Requires Docker knowledge for all contributors.
