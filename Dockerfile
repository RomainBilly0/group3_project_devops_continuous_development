# STAGE 1: Build the binary
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files first (improves caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary (ideal for containers)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./webapi

# STAGE 2: Create the final lightweight image
FROM alpine:latest

WORKDIR /root/

# Copy only the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your API runs on (change 8080 if needed)
EXPOSE 8080

# Run the binary
CMD ["./main"]