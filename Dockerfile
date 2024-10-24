# Dockerfile References: https://docs.docker.com/engine/reference/builder/
#
# export DOCKER_BUILDKIT=0
# export COMPOSE_DOCKER_CLI_BUILD=0
#
# Build the docker image
# `docker build -t go-kubernetes .`
#
# Tag the image
# `docker tag go-kubernetes <username>/go-name-store:1.0.0`
#
# Login to docker with your docker Id
# `docker login`
#
# Login with your Docker ID to push and pull images from Docker Hub. 
# If you do not have a Docker ID, head over to https://hub.docker.com to create one.
# Username: 
# Password:
# Login Succeeded
#
# Push the image to docker hub
# `docker push <username>/go-name-store:1.0.0`

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Hasti"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
