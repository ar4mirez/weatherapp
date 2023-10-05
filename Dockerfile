# Use the official Golang image as a parent image
FROM golang:1.21.1-alpine

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . .

RUN go mod download

# Build the Golang application
RUN go build -o weatherapp

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run your application
CMD ["./weatherapp"]
