# Stage 1
FROM golang:alpine AS builder

# Setup base application directory
ENV BASE_APP_DIR /go/src/github.com/mjakobczyk/daily-content
WORKDIR ${BASE_APP_DIR}

# Enable go modules
ENV GO111MODULE=on

# Install necessary applications
RUN apk add make bash

# Copy files
COPY . .

# Build project and copy executable to dedicated directory
RUN make build && mkdir /app && mv ./main /app/main

# Stage 2
FROM alpine:latest

# Setup working directory
RUN mkdir /app && chmod 777 /app
WORKDIR /app

# Enable no cache to keep container as small as possible
# Install Certificates in Alpine Image to establish Secured Communication
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app /app

# Expose port which is going to be used
EXPOSE 8080

# Run project from the command line
CMD ["/app/main"]
