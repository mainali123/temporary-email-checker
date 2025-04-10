# Temp Mail Service

This is a Go-based web application that validates email addresses to check if they belong to temporary email domains. It also periodically updates the list of temporary email domains from an external source.

## Features

- Validate email addresses to check if they belong to temporary email domains.
- Automatically update the list of temporary email domains every 7 days.
- RESTful API with endpoints for email validation.
- Logging for application events and errors.
- Docker support for containerized deployment.

## Prerequisites

- Go 1.24 or later installed on your system (if not using Docker).
- Docker and Docker Compose installed on your system.
- Internet access to fetch the list of temporary email domains.

## Installation

### Using Docker

1. Clone the repository:
   ```shell
   git clone https://github.com/mainali123/temporary-email-checker
   cd temporary-email-checker
   ```

2. Build and run the application using Docker Compose:
   ```shell
   docker-compose up --build
   ```

3. Access the application at `http://localhost:8080`.

### Without Docker

1. Clone the repository:
   ```shell
   git clone https://github.com/mainali123/temporary-email-checker
   cd temporary-email-checker
   ```

2. Build the application:
   ```shell
   go build -o builds/temp_mail ./main.go ./routes.go
   ```

3. Create the required directories:
   ```shell
   cd builds
   mkdir email
   ```

4. Run the application:
   ```shell
   ./temp_mail
   ```

## API Endpoints

### GET `/`

Returns a welcome message and the application version.

**Response:**

```json
{
  "message": "Welcome to the Go web server!",
  "version": "1.0.0"
}
```

### POST `/validate-email`

Validates if the provided email address belongs to a temporary email domain.

**Request Body:**

```json
{
  "email": "example@domain.com"
}
```

**Response:**

```json
{
  "email": "example@domain.com",
  "isTemporary": true
}
```

## Configuration

- The list of temporary email domains is stored in `email/emails.txt`.
- The application fetches the latest list from `https://raw.githubusercontent.com/7c/fakefilter/main/txt/data.txt`.

## Logging

- Application logs are stored in `main.log`.
- Email-related logs are stored in `email/email.log`.

## Scheduler

- The application uses a scheduler to update the list of temporary email domains every 7 days.

## Docker

- The `Dockerfile` defines the container for the application.
- The `docker-compose.yml` file sets up the service and maps the necessary ports and volumes.

### Commands

- Build and run the application:
  ```shell
  docker-compose up --build
  ```

- Stop the application:
  ```shell
  docker-compose down
  ```

- View logs:
  ```shell
  docker logs <container_name>
  ```

## Dependencies

The project uses the following Go modules:

- `github.com/go-co-op/gocron/v2`
- `github.com/google/uuid`
- `github.com/gorilla/mux`
- `github.com/jonboulle/clockwork`
- `github.com/robfig/cron/v3`

## License

MIT License

Copyright (c) 2025 mainali123

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to use the Software, subject to the following conditions:

1. Credit must be given to the original author (mainali123) in any derivative works, publications, or distributions of this Software.
2. The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF, OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.```