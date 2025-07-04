# ADR 003: Environment Management with Docker Compose
Status: Accepted

### Context:
Managing different configurations for Dev and Prod is necessary.

### Decision:
* Use a base docker-compose.yml for shared configuration.
* Use override files: docker-compose.dev.yml for development and docker-compose.prod.yml for production.
* Developers will use Docker Compose commands with appropriate override files to start the correct environment.

### Consequences:
* Simplifies switching between environments.
* Reduces configuration drift.
* Adds minor complexity to Docker Compose management.
