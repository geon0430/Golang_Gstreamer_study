FROM nvcr.io/nvidia/pytorch:23.01-py3

ENV TZ=Asia/Seoul
ENV DEBIAN_FRONTEND=noninteractive

ENV NVIDIA_VISIBLE_DEVICES all
ENV NVIDIA_DRIVER_CAPABILITIES all

ENV XDG_RUNTIME_DIR "/tmp"

RUN python3 -m pip install --upgrade pip

WORKDIR /
RUN mkdir -p /go_vms
COPY . /go_vms

RUN bash /go_vms/setting-scripts/install_dependencies.sh

RUN bash /go_vms/setting-scripts/install_golang.sh

ARG PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

ENV GOTOOLCHAIN=local
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:$PATH
ENV GOROOT=/usr/local/go
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"

RUN bash /go_vms/setting-scripts/install_ffmpeg.sh

RUN bash /go_vms/setting-scripts/install_OpenCV.sh

# RUN apt-get update && apt-get install -y sudo
# RUN mkdir -p "$GOPATH/src/gocv.io/x/"
# RUN cd "$GOPATH/src/gocv.io/x/" && \
#     git clone https://github.com/hybridgroup/gocv.git && \
#     cd gocv && \
#     go run ./cmd/version/main.go
#         cd gocv && \
#         make install_cuda BUILD_SHARED_LIBS=OFF

#RUN bash /go_vms/setting-scripts/install_gocv.sh
#RUN bash /go_vms/setting-scripts/install_gst.sh

#RUN /bin/bash -c "ldconfig"
