package main

import (
    "fmt"
    "gocv.io/x/gocv"
)

func main() {
    url := "rtsp://admin:qazwsx123!@192.168.10.70/0/1080p/media.smp"

    cap, err := gocv.OpenVideoCapture(url)
    if err != nil {
        fmt.Printf("Error opening video capture: %v\n", err)
        return
    }
    defer cap.Close()

    window := gocv.NewWindow("RTSP Stream")
    defer window.Close()

    img := gocv.NewMat()
    defer img.Close()

    for {
        if ok := cap.Read(&img); !ok {
            fmt.Println("Error reading frame from stream")
            break
        }

        if img.Empty() {
            fmt.Println("Empty frame received")
            continue
        }

        window.IMShow(img)
        if window.WaitKey(1) >= 0 {
            break
        }
    }
}
