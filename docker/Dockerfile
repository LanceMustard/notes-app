# Stage 1: Build the Go app
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and main.go from src
COPY ../src/go.mod .
COPY ../src/main.go .

# Download dependencies and build binary
RUN go mod tidy
RUN go build -o server

# Stage 2: Use a minimal image to run the app
FROM gcr.io/distroless/static

# Copy the binary from the builder stage
COPY --from=builder /app/server /server

# Expose the port
EXPOSE 8080

# Command to run
ENTRYPOINT ["/server"]

