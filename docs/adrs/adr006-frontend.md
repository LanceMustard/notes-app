# ADR 006: Frontend Build and Asset Management
Status: Accepted

### Context:
Frontend uses HTMX and Tailwind CSS. Efficient asset management and build processes are required.

### Decision:
* Tailwind CSS will be built during the Docker build process using Node.js and Tailwind CLI.
* HTMX will be included via CDN for simplicity or bundled with static assets for production.

### Consequences:
* Ensures optimized frontend assets in production.
* Keeps Docker images self-contained.
