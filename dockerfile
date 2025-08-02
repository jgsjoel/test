# Step 1: Build the Go app
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy only the files you have
COPY go.mod ./
# If you don't have go.sum yet, this is fine — `go mod tidy` will create it
RUN go mod tidy

COPY main.go ./
RUN go build -o main .

# Step 2: Minimal final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]
