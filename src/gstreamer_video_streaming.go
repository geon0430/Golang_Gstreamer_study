package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/giorgisio/goav/gstreamer"
)

func main() {
	gstreamer.Init(nil)

	pipelineStr := "filesrc location=/path/to/your/video.mp4 ! decodebin ! videoconvert ! x264enc bitrate=500 ! rtph264pay config-interval=1 pt=96 ! udpsink host=127.0.0.1 port=5000"
	pipeline, err := gstreamer.ParseLaunch(pipelineStr)
	if err != nil {
		fmt.Println("Failed to create pipeline:", err)
		return
	}

	pipeline.SetState(gstreamer.StatePlaying)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pipeline.SetState(gstreamer.StateNull)
		os.Exit(0)
	}()

	gstreamer.MainLoopRun()
}
