# Stage 1: Build the application
FROM golang:1.21.6 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/

# Stage 2: Create the final lightweight image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
