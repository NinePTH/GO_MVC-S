# Use Go 1.23 as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source files
COPY src/ ./src/

# Set the working directory to match your application's structure
WORKDIR /app/src

# Build the application
RUN go build -o ../main

# Expose the port your Echo app listens on
EXPOSE 1323

# Command to run the application
CMD ["../main"]
