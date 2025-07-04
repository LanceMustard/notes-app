# ADR 007: SQLite Database Management
Status: Accepted

### Context:
The application uses SQLite for data persistence.

### Decision:
* SQLite database file will be mounted to a Docker volume for persistence.
* Database migrations will be run as part of container startup or as a separate CI/CD job.

### Consequences:
* Ensures data persists across container restarts.
* Simplifies local development and testing.
* May require additional migration tooling for schema changes.
