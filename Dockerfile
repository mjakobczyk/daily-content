# Retrieve the `golang:alpine` image to provide the
# necessary Golang tooling for building Go binaries.
# 
# Here I retrieve the `alpine`-based just for the 
# convenience of using a tiny image.
FROM golang:alpine AS builder

# Set working directory.
WORKDIR /

# Install required applications.
RUN apk add make bash

# Copy entire content of application.
COPY . .

# Build project.
RUN make build

# Start a new stage from scratch.
FROM alpine:latest

# Enable no cache to keep container as small as possible.
# Install Certificates in Alpine Image to establish Secured Communication.
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
# to the root directory.
COPY --from=builder daily-content .

# Expose port which is going to be used.
EXPOSE 8080

# Run project from the command line.
CMD ["./daily-content"]
