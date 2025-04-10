# Use the official Go image as the base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o temp_mail ./main.go ./routes.go

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./temp_mail"]