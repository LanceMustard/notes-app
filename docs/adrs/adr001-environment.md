# ADR 001: Environment Separation – Dev and Prod
Status: Accepted

### Context:
The application requires two distinct environments: Development (Dev) for testing and experimentation, and Production (Prod) for stable, user-facing deployment.

### Decision:
* Each environment will use isolated AWS resources, ideally in separate accounts or at least with clear resource separation.
* Dev and Prod will have independent deployment pipelines and infrastructure to prevent cross-contamination and enable safe experimentation.

### Consequences:
* Reduces risk of accidental data leaks or outages affecting production.
* Enables rapid iteration in Dev without impacting Prod.
* Slightly increases AWS resource management complexity.
