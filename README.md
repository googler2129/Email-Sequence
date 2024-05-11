# Email Service

Email service is responsible for managing sequence. A Sequence consists of multiple steps.

### Things to know
For making changes and easily navigating kitchen service's codebase, here are a few things that you must have a knowledge of -
1. [Golang](https://tour.golang.org/welcome/1)
2. REST/HTTP
3. PostgreSQL

#### Good to know
1. [Docker](https://docs.docker.com/get-started/)


### Setting up the service

1. Fork this repository `googler2129/Email-Sequence` to `<YOUR_GITHUB_USER_NAME>/Email-Sequence`.
2. Clone your fork here by `git clone git@github.com:<YOUR_GITHUB_USER_NAME>/email-sequence.git`
3. Run DB migrations using `CONFIG_SOURCE=local go run main.go --mode=migration`

### APIs

1. POST: api/v1/sequences -
2. PUT: api/v1/sequences/:sequence_id
3. PUT: api/v1/sequences/:sequence_id/steps/:step_id
4. DELETE: PUT: api/v1/sequences/:sequence_id/steps/:step_id
