package util

import (
	"testing"
	// "github.com/sirupsen/logrus"
	"time"
	"context"
	"os"
	"fmt"
)

func TestLog(t *testing.T){

	Context, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(500 * time.Millisecond)
	}()

	logPath := "/go_vms/log"
	MODEL := "MODEL"
	NAME := "TEST"
	logLevel := "DEBUG"

	logger := SetupLogging(Context,logPath, MODEL, NAME, logLevel)
	
	logFilePath := "/go_vms/log/TEST/TEST_DEBUG_go.log"
	go func() {
		for {
			select {
			case <-Context.Done():
				return
			case <-time.After(1 * time.Second):
				fileInfo, err := os.Stat(logFilePath)
				if err != nil {
					t.Fatalf("Failed to get file info: %v", err)
				}


				fileSizeMB := float64(fileInfo.Size()) / 1024 / 1024
				fmt.Printf("File size: %.2f MB\n", fileSizeMB)
			}
		}
	}()


	frameDuration := time.Duration( 1 * time.Nanosecond)
	ticker := time.NewTicker(frameDuration)
	defer ticker.Stop()

	for {
		select {
		case <- Context.Done():
			return
		case <- ticker.C:
			logger.Debug("This is a DEBUG level message")
			logger.Info("This is an INFO level message")
			logger.Warn("This is a WARNING level message")
			logger.Error("This is an ERROR level message")

		}
	}
}
