#!/bin/bash

apt-get update && apt-get install -y sudo

mkdir -p $GOPATH/src/gocv.io/x/

cd $GOPATH/src/gocv.io/x/ && git clone https://github.com/hybridgroup/gocv.git && \
        cd gocv && \
        make install_cuda BUILD_SHARED_LIBS=OFF
