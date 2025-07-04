# ADR 004: CI/CD Pipeline with GitHub Actions
Status: Accepted

### Context:
Automated, reliable, and repeatable deployments are required for both environments.

### Decision:
* Use GitHub Actions for CI/CD.
* On every push, build and test Docker images.
* Push images to AWS Elastic Container Registry (ECR).
* Deploy to AWS (ECS, EC2, or App Runner) using separate workflows or branches for Dev and Prod.

### Consequences:
* Enables automated testing and deployment.
* Supports approval gates for production deployments.
* Requires maintenance of GitHub Actions workflows.
