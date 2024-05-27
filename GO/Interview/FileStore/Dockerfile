# Use the official Golang image as the base image
FROM golang:1.22.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o filestore 

# Use a minimal image for the final build
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/filestore .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./filestore"]
