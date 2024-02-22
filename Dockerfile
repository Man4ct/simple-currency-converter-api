# Use the official Golang image as a base image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -v -o /usr/local/bin/app ./main.go

# Start a new stage
FROM mongo:latest

# Expose port 8080 for the Go application
EXPOSE 8080

# Set the working directory inside the container
WORKDIR /usr/local/bin

# Copy the built Go binary from the previous stage
COPY --from=builder /usr/local/bin/app .

# Start the Go application
CMD ["./app"]
