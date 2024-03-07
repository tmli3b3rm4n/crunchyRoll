# Start with the Go official image as the base image
FROM golang:1.21.3

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application with CGO disabled and for the Linux OS
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Expose port 8080 to allow external access to the application
EXPOSE 8080

# Run the application when the container starts
CMD ["/docker-gs-ping"]