# Email Service

Email service is responsible for managing sequence. A Sequence consists of multiple steps.

### Things to know
For making changes and easily navigating kitchen service's codebase, here are a few things that you must have a knowledge of -
1. [Golang](https://tour.golang.org/welcome/1)
2. [REST/HTTP](https://aws.amazon.com/what-is/restful-api/)
3. [PostgreSQL](https://www.postgresql.org/)

#### Good to know
1. [Docker](https://docs.docker.com/get-started/)


### Setting up the service

1. Fork this repository `googler2129/Email-Sequence` to `<YOUR_GITHUB_USER_NAME>/Email-Sequence`.
2. Clone your fork here by `git clone git@github.com:<YOUR_GITHUB_USER_NAME>/email-sequence.git`
3. Run `docker compose up` for running postgreSQL server
4. Run DB migrations using `CONFIG_SOURCE=local go run main.go --mode=migration`
5. Run `CONFIG_SOURCE=local go run main.go`

### APIs

1. `POST: api/v1/sequences` - This API is responsible for creating an email sequence along with steps.
```go
Request:
   {
    "name": "Sequence 1",
    "is_open_tracking_enabled": false,
    "is_click_tracking_enabled": true,
    "sequence_steps": [
        {
            "subject": "Leave",
            "content": "Kindly allow leave for 10 days due to medical emergency",
            "waiting_days": 5,
            "serial_order": 1
        }
    ],
    "user_id": "10",
    "user_name": "Daniel"
   }
   
Response:   
     "Success"
   
```
2. `PUT: api/v1/sequences/:sequence_id` - This API is responsible for updating an email sequence.
```go
Request: 
   {
    "is_open_tracking_enabled": false,
    "is_click_tracking_enabled": true,
    "user_id": "5",
    "user_name": "Depender"
   }
   
Response:   
     "Success"
   
```
3. `PUT: api/v1/sequences/:sequence_id/steps/:step_id` - This API is responsible for updating a step.
```go
Request:
   {
    "subject": "Feature announcement",
    "content": "Thrilled to share email feature is released now",
    "user_id": "5",
    "user_name": "Depender"
   }
   
Response:   
     "Success"
```   
4. `DELETE: PUT: api/v1/sequences/:sequence_id/steps/:step_id` - This API is responsible for deleting a step.
```go
Response:   
     "Success"
```
