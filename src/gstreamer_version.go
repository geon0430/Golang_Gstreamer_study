package main

import (
	"fmt"
	"github.com/notedit/gst"
)

func main() {
	gst.Init(nil)

	version := gst.GetVersion()

	fmt.Printf("GStreamer version: %s\n", version)
}
