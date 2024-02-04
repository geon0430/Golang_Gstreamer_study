FROM nvcr.io/nvidia/pytorch:23.01-py3

ENV TZ=Asia/Seoul
ENV DEBIAN_FRONTEND=noninteractive

ENV NVIDIA_VISIBLE_DEVICES all
ENV NVIDIA_DRIVER_CAPABILITIES all

ENV XDG_RUNTIME_DIR "/tmp"

RUN python3 -m pip install --upgrade pip

WORKDIR /
RUN mkdir -p /Golang_Gstreamer
COPY . /Golang_Gstreamer

RUN bash /Golang_Gstreamer/setting_scripts/install_dependencies.sh

RUN bash /Golang_Gstreamer/setting_scripts/install_golang.sh

ARG PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

ENV GOTOOLCHAIN=local
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:$PATH
ENV GOROOT=/usr/local/go
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"

RUN bash /Golang_Gstreamer/setting_scripts/install_ffmpeg.sh

RUN bash /Golang_Gstreamer/setting_scripts/install_OpenCV.sh

RUN apt-get update && apt-get install -y sudo
