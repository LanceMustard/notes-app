# ADR 005: AWS Hosting Strategy
Status: Accepted

### Context:
The application must be hosted in the cloud, with support for containerized workloads and easy integration with CI/CD.

### Decision:
* Use AWS ECS (Fargate), EC2, or App Runner as deployment targets.
* Separate clusters, instances, or services will be used for Dev and Prod.

### Consequences:
* Leverages managed AWS services for scalability and reliability.
* Simplifies integration with Docker and GitHub Actions.
* May incur additional AWS costs for environment separation.
