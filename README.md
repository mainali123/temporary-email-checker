# Temp Mail Service

This is a Go-based web application that validates email addresses to check if they belong to temporary email domains. It also periodically updates the list of temporary email domains from an external source.

## Features

- Validate email addresses to check if they belong to temporary email domains.
- Automatically update the list of temporary email domains every 7 days.
- RESTful API with endpoints for email validation.
- Logging for application events and errors.

## Prerequisites

- Go 1.24 or later installed on your system.
- Internet access to fetch the list of temporary email domains.

## Installation

1. Clone the repository:
   ```shell
   git clone https://github.com/mainali123/temporary-email-checker
   cd temporary-email-checker

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

## Dependencies

The project uses the following Go modules:
- `github.com/go-co-op/gocron/v2`
- `github.com/google/uuid`
- `github.com/gorilla/mux`
- `github.com/jonboulle/clockwork`
- `github.com/robfig/cron/v3`

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```