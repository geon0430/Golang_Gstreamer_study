#!/bin/bash

GOLANG_VERSION="1.21.3"

# Golang
mkdir -p /go
cd /go && \
    wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz

rm go${GOLANG_VERSION}.linux-amd64.tar.gz

source ~/.bashrc
