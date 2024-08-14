# Stage 1: Build stage
FROM golang:1.22.1 AS builder

WORKDIR /app

# Copy and download dependencies
# COPY go.mod .
# COPY go.sum .

# Copy the rest of the application
COPY . .
RUN go mod download

# Optionally copy the .env file if it's needed
COPY .env .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .
COPY --from=builder /app/app.log ./logs/

# Copy the configuration files
# COPY --from=builder /app/internal/casbin/rbac_model.conf ./internal/casbin/
# COPY --from=builder /app/internal/casbin/policy.csv ./internal/casbin/

# Optionally copy the .env file if it's needed
COPY --from=builder /app/.env .
COPY --from=builder /app/casbin/model.conf ./casbin/

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]