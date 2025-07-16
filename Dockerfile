# Start with the official Go image
FROM golang:1.22-alpine


# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go app
RUN go build -o server .

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./server"]
