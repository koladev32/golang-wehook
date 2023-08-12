# Use an official Go image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build .

# Specify the command to run when the container starts
CMD ["./app"]