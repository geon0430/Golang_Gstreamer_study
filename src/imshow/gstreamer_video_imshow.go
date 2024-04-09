package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("bash", "-c",
		"gst-launch-1.0 videotestsrc ! "+
			"videoconvert ! nvh264enc ! "+
			"rtspclientsink location='rtsp://localhost:8554/mystream'")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error starting command: %v", err)
		return
	}

	cmd.Wait()
}