package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	gst "github.com/go-gst/go-gst"
)

func main() {
	if err := gst.Init(); err != nil {
		log.Fatalf("failed to initialize gst: %v", err)
	}

	pipeline, err := gst.ParseLaunch("udpsrc port=9000 caps=\"application/x-rtp, media=(string)video, clock-rate=(int)90000, encoding-name=(string)H264\" ! rtph264depay ! avdec_h264 ! videoconvert ! autovideosink")
	if err != nil {
		log.Fatalf("failed to create pipeline: %v", err)
	}

	if err := pipeline.SetState(gst.StatePlaying); err != nil {
		log.Fatalf("failed to set pipeline state to playing: %v", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		pipeline.SetState(gst.StateNull)
		os.Exit(0)
	}()

	gst.MainLoopRun()
}
