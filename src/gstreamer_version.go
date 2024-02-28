package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("gst-launch-1.0", "--version")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error starting command: %v", err)
		return
	}
}