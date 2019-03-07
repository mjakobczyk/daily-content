# Retrieve the `golang:alpine` image to provide the 
# necessary Golang tooling for building Go binaries.
# 
# Here I retrieve the `alpine`-based just for the 
# convenience of using a tiny image.
FROM golang:alpine AS builder

# Add the project content.
ADD . /go/src/github.com/daily-content/

# Add the vendor dependencies.
ADD ./vendor /go/src/github.com/daily-content/vendor

# Set the current working directory inside the container.
WORKDIR /go/src/github.com/daily-content/

# Install make tool/
RUN apk add make bash

# Build project.
# RUN make build
RUN make build && \
    mv ./daily-content /usr/bin/daily-content

# Start a new stage from scratch
FROM alpine:latest

# Enable no cache to keep container as small as possible.
# Install Certificates in Alpine Image to establish Secured Communication.
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
# to the root directory.
COPY --from=builder /usr/bin/daily-content .

# Expose port which is going to be used.
EXPOSE 8080

# Run project from the command line.
ENTRYPOINT ./daily-content
