# Use an official Golang runtime as a parent image
FROM golang:1.16

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install any needed dependencies specified in go.mod and go.sum
RUN go mod download

# Build the Go application
RUN go build -o url-shortner

# Expose port 8080 for the application
EXPOSE 8080

# Run the application
CMD ["./url-shortner"]
