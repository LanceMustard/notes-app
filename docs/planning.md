### 1. **Overview**
Developing a web based, cloud hosted Notes App.

- **Backend:** Golang, SQLite
- **Frontend:** HTMX, Tailwind CSS
- **Containerization:** Docker
- **Cloud:** AWS
- **Source Control:** GitHub
- **Environments:** Dev and Prod

### 2. **Environment Structure**

- **Development (Dev):** For testing new features, debugging, and integration.
- **Production (Prod):** Stable, user-facing environment.

**Best Practice:** Use separate AWS resources for each environment to avoid accidental cross-contamination and to enable safe experimentation.

### 3. **Containerization with Docker**

- **Multi-Stage Docker Builds:** Use multi-stage builds to keep images small and secure. Build Go binary in one stage, then copy it into a minimal base image (e.g., `scratch` or `alpine`).
- **SQLite Persistence:** Store SQLite database in a Docker volume to ensure data persists across container restarts and updates.
- **Environment Variables:** Use environment variables for configuration (e.g., database path, debug mode). Use `.env.dev` and `.env.prod` files for each environment.

### 4. **Docker Compose for Multi-Environment Management**

- **Base Compose File:** `docker-compose.yml` for shared configuration.
- **Overrides:** 
  - `docker-compose.dev.yml` for development-specific settings (e.g., hot reload, debug).
  - `docker-compose.prod.yml` for production (e.g., optimized builds, no debug).

**Scripts:**
```sh
run-dev.sh
run-prod.sh
```

### 5. **CI/CD Pipeline with GitHub Actions**

- **Build & Test:** On every push, build and test Docker image.
- **Push to Registry:** Push images to a container registry (e.g., AWS ECR).
- **Deploy:** Use GitHub Actions to deploy to AWS (ECS, EC2, or App Runner) for both environments. Use separate workflows or branches for Dev and Prod.

**Typical Workflow:**
1. Developer pushes code to GitHub.
2. GitHub Actions builds Docker image, runs tests.
3. Image is pushed to AWS ECR.
4. Deployment job updates the running service in the target environment.

### 6. **AWS Hosting Options**

- **Amazon ECS (Fargate):** Managed container orchestration, integrates well with Docker Compose and GitHub Actions.
- **EC2:** For more control, deploy containers directly to EC2 instances.
- **App Runner:** Simplifies deployment of containerized web apps, integrates with GitHub and ECR.

**Separation:** Use different ECS clusters, EC2 instances, or App Runner services for Dev and Prod.

### 7. **Frontend Build (HTMX & Tailwind)**

- **Tailwind CSS:** Use a build step in your Dockerfile to generate the final CSS. This can be done with Node.js and Tailwind CLI.
- **HTMX:** Include via CDN for simplicity, or bundle with your static assets for production.

### 8. **Database Management**

- **SQLite in Docker:** Mount the database file to a Docker volume for persistence.
- **Migrations:** Run database migrations as part of your container startup or as a separate job in your CI/CD pipeline.

### 9. **Summary Table**

| Component         | Dev Environment                | Prod Environment                |
|-------------------|-------------------------------|---------------------------------|
| AWS Resources     | Isolated (ideally separate)    | Isolated (ideally separate)     |
| Docker Compose    | `docker-compose.dev.yml`       | `docker-compose.prod.yml`       |
| Database          | SQLite (Docker volume)         | SQLite (Docker volume)          |
| CI/CD             | GitHub Actions (auto-deploy)   | GitHub Actions (manual/approval)|
| Deployment Target | ECS/EC2/App Runner             | ECS/EC2/App Runner              |

### 10. **References for Further Reading**

- Multi-environment Docker deployment best practices
- AWS multi-environment setup and CI/CD
- Dockerizing Go and SQLite
- GitHub Actions for AWS deployments
- Frontend build with HTMX and Tailwind
