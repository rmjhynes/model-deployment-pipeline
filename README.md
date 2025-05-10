# Model Deployment Pipeline
A repo to model the desired steps in a CI pipeline to build and release a containerised application image.
The CI tool of choice is GitHub Actions as it integrates well with this GitHub repo and is the tool I would most likely use at work.

# Language
The application code is written in Go. Why Go? I never write Go so thought I would switch it up from Python.
The actual application is an extremely basic calculator because I needed something to build as an image to release.

# Pipeline Stages
## Source
- When an engineer commits a change, pre-commit hooks run to perform basic file formatting and go-specific checks.
- When a PR is created, `golangci-lint` runs as a pre-merge check to perform Go linting before the change can be merged into main.

## Build
Once merged, the GitHub Actions CI pipeline is triggered.
- Run unit tests on application code.
- Build Docker image to containerise the application code.

## Secure
- Scan Docker image for _all_ severity vulnerabilties with [Trivy](trivy.dev/) and output in table format.
- Generate SBOM of all the project dependencies and send to GitHub Dependency Graph for visibility within the GitHub repo.

## Test
- Run integration tests with [Testcontainers](https://testcontainers.com/) to check the application can interact with other containerised instances.
