# Step 1: Build the Go app
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies (cached if unchanged)
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the app
RUN go build -o main .

# Step 2: Minimal final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
