#!/bin/bash

apt-get update && \
	apt install -y build-essential ninja-build flex bison && \
	python3 -m pip install meson

cd / && \
mkdir -p /programs && cd /programs && \
	git clone https://github.com/GStreamer/gstreamer.git && \
	cd gstreamer && \
	meson setup builddir && \
	meson compile -C builddir && \
	ninja -C builddir && \
	ninja -C builddir install 

ldconfig

apt install -y gstreamer1.0-rtsp libcanberra-gtk-module libcanberra-gtk3-module