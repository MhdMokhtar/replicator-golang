# Use the official Golang image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 9000

# Command to run the application
CMD ["./main"]