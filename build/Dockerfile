# Use an official Go runtime as a parent image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main ./internal/cmd/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
