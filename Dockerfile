# -------- Stage 1: Build the Go binary --------
# Use the official Go image based on Alpine Linux for a lightweight build environment
FROM golang:alpine3.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files to leverage Docker layer caching
# (these files rarely change, so this step is usually cached)
COPY go.* ./
RUN go mod download

# Copy the rest of the application source code into the container
COPY . .

# Build the Go application and output the binary as "calculator-app"
# Using a unique binary name avoids name conflicts with any directory names
RUN go build -v -o calculator-app .

# -------- Stage 2: Create a minimal runtime image --------
# Use a clean Alpine Linux image for running the app with minimal size
FROM alpine:latest

# Create a non-root user for better security
RUN adduser -D appuser

# Copy the compiled binary from the builder stage to a system-wide binary path
COPY --from=builder /app/calculator-app /usr/local/bin/calculator-app

# Make sure the binary has execute permissions
RUN chmod +x /usr/local/bin/calculator-app

# Switch to the non-root user for running the application
USER appuser

# Define the default executable when the container starts
ENTRYPOINT ["/usr/local/bin/calculator-app"]
