FROM golang:1.22

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the entire source code to the working directory
COPY . .

# Download and verify dependencies
RUN go mod download && go mod verify

# Build the Go application
RUN go build -v -o ./bin/main ./cmd/main.go

# Expose the port outside the container
EXPOSE 8080

# Run the application
CMD ["./bin/main"]