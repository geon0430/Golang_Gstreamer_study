package main

import (
	"log"
	"github.com/your/gstreamer-go"
)

func main() {
	err := gstreamer.Init()
	if err != nil {
		log.Fatal("Failed to initialize GStreamer:", err)
	}

	pipelineStr := "rtspsrc location=rtp://YOUR_RTPS_ADDRESS ! decodebin ! autovideosink"
	pipeline, err := gstreamer.ParseLaunch(pipelineStr)
	if err != nil {
		log.Fatal("Failed to create pipeline:", err)
	}

	err = pipeline.SetState(gstreamer.StatePlaying)
	if err != nil {
		log.Fatal("Failed to set pipeline to playing state:", err)
	}

	gstreamer.MainLoopRun()
}
