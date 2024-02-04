#!/bin/bash/

IMAGE_NAME="golang_Gstreamer"

TAG="test-0.1"

docker build --no-cache -t ${IMAGE_NAME}:${TAG} .
