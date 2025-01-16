# Base image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy source code
COPY . .

# Download dependencies and build the application
RUN go mod download
RUN go build -o email_service main.go

# Expose the gRPC port
EXPOSE 50051

# Start the email service
CMD ["./email_service"]
