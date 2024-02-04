#!/bin/bash/

port_num="2"
CONTAINER_NAME="geon_go_vms_ffmpeg"
IMAGE_NAME="hub.inbic.duckdns.org/dev-1-team/go_vms"
TAG="goCuda_FFmpeg-0.1"

port_num="2"
go_vms_path=$(pwd)

docker run \
    --runtime nvidia \
    --gpus all \
    -it \
    -p ${port_num}1574:1574 \
    -p ${port_num}1935:1935 \
    -p ${port_num}8888:8888 \
    -p ${port_num}8889:8889 \
    -p ${port_num}8444:8444 \
    -p ${port_num}8445:8445 \
    -p ${port_num}8554:8554 \
    -p ${port_num}8449:8449 \
    -p ${port_num}8540:8450 \
    -p ${port_num}9000:9000 \
    -p ${port_num}9001:9001 \
    --name ${CONTAINER_NAME} \
    --privileged \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    -v /home/inbic/geon:/go_vms \
    -v ${go_vms_path}//go_vms/ \
    -e DISPLAY=$DISPLAY \
    --shm-size 5g \
    --restart=always \
    -w /go_vms \
    ${IMAGE_NAME}:${TAG}
