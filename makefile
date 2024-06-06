IMAGE_NAME := my-image ## Docker image name

CONTAINER_NAME := my-container ## Docker container name

PORT := 80## Exposed port on host

build-and-run: build-image run-image ## Build and run image

build-image:## Build image
    @docker build -t ${IMAGE_NAME}:latest .

## Run image
run-p: 
    go run server.go