# Start with a lightweight Go image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY cmd/db-populate-script ./cmd/db-populate-script

# Set the working directory to your script
WORKDIR /app/cmd/db-populate-script

# Build the Go application
RUN go build -o /app/main .

# Set the entry point to the compiled binary
ENTRYPOINT ["/app/main"]