# Stage 1: Build
FROM golang:latest AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Set executable permissions
RUN chmod +x main

EXPOSE 8080

# Run the application
CMD ["./main"]
