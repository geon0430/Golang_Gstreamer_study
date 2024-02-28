package main

/*
#cgo pkg-config: gstreamer-1.0
#include <gst/gst.h>
#include <stdlib.h>
*/
import "C"
import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func main() {
	C.gst_init(nil, nil)
	
	pipelineStr := "filesrc location=../../video/sample_1080p_h264.mp4 ! qtdemux ! h264parse ! avdec_h264 ! videoconvert ! autovideosink"
	pipelineCStr := C.CString(pipelineStr)
	defer C.free(unsafe.Pointer(pipelineCStr))

	pipeline := C.gst_parse_launch(pipelineCStr, nil)
	if pipeline == nil {
		log.Fatal("Failed to create pipeline")
	}

	ret := C.gst_element_set_state(pipeline, C.GST_STATE_PLAYING)
	if ret == C.GST_STATE_CHANGE_FAILURE {
		log.Fatal("Failed to set pipeline state to playing")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for range c {
			C.gst_element_set_state(pipeline, C.GST_STATE_NULL)
			C.gst_object_unref(C.gpointer(pipeline))
			os.Exit(0)
		}
	}()

	C.gst_element_get_bus(pipeline)
	C.g_main_loop_new(nil, C.gboolean(0))
	loop := C.g_main_loop_new(nil, C.gboolean(0))
	C.g_main_loop_run(loop)
}
