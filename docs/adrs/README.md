# Architecture decision records
| ADR# | Decision Area | Key Decision |
|------|---------------|--------------|
| 001	| Environment Separation	| Isolated Dev and Prod environments on AWS |
| 002	| Containerization	| Docker with multi-stage builds and volumes |
| 003	| Env Management	| Docker Compose with override files |
| 004	| CI/CD	| GitHub Actions for build, test, deploy |
| 005	| AWS Hosting	| ECS/EC2/App Runner, separate for each environment |
| 006	| Frontend Build	| Tailwind build in Docker, HTMX via CDN or bundle |
| 007	| Database Management	| SQLite in Docker volume, migrations in CI/CD |
