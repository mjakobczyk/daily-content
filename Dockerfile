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

# Build project.
RUN --set ex && make build

# Start a new stage from scratch
FROM alpine:latest

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /go/daily-content content

# Expose port which is going to be used.
EXPOSE 8080

# Run project from the command line.
CMD ["./daily-content"]