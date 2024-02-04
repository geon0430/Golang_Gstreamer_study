#!/bin/bash/

IMAGE_NAME="golang_gstreamer"

TAG="test-0.1"

docker build --no-cache -t ${IMAGE_NAME}:${TAG} .
