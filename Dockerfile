# --- Stage 1: Build ---
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary (better for containers)
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# --- Stage 2: Final Image ---
FROM alpine:latest

# Install CA certificates (needed for HTTPS calls)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary and migrations directory from the builder stage
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations

# Expose the port your app runs on
EXPOSE 8080
EXPOSE 50051

# Command to run the application
CMD ["./server"]