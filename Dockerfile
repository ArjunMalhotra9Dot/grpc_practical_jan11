#! Stage 1: Build Stage
FROM golang:1.23-alpine AS build

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o email_service main.go

#! Stage 2: Final Stage
FROM alpine:3.18

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/email_service .

# Install CA certificates and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Expose the gRPC port
EXPOSE 50051

# Start the email service
CMD ["./email_service"]
