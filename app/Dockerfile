# Use the official Golang 1.22 image to build the application
FROM golang:1.22 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Verify the build output
RUN echo "Build stage: /app directory contents:" && ls -l /app

# Use debian as the base image
FROM debian:latest

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main /root/main

# Verify the copied binary
RUN echo "Final stage: /root directory contents:" && ls -l /root

# Ensure the binary is executable
RUN chmod +x /root/main

# Verify permissions
RUN echo "Permissions check:" && ls -l /root/main

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/root/main"]
