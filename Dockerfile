# Stage 1: Build binary
FROM golang:1.24.4-alpine AS builder

# Cài thêm git nếu cần go mod có package từ github
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files and download dependencies first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

RUN go mod vendor
# Đặt lại thư mục làm việc để build chính xác
WORKDIR /app/cmd/api

# Build the Go binary
RUN go build -mod=vendor -o api .

# Stage 2: Minimal image
FROM alpine:latest

WORKDIR /app

# Copy binary từ builder stage
COPY --from=builder /app/cmd/api/api .

# Copy .env nếu cần khi runtime
COPY .env .env

# Expose port của ứng dụng nếu cần (chỉnh theo thực tế)
EXPOSE 8080

# Run the binary
CMD ["./api"]
