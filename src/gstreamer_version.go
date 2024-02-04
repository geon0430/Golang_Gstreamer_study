package main

import (
	"fmt"
	"github.com/notedit/gst"
)

func main() {
	// GStreamer의 초기화
	gst.Init(nil)

	// GStreamer 버전 정보 가져오기
	version := gst.GetVersion()

	fmt.Printf("GStreamer version: %s\n", version)
}
